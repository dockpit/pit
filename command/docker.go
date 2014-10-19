package command

import (
	"crypto/md5"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/fsouza/go-dockerclient"

	"github.com/dockpit/pit/spec"
)

var NamePrefix = "dockpit"
var ImageName = "dockpit/pit"

type Docker struct {
	client  *docker.Client
	version string
}

func NewDocker(addr string, cert string, version string) (*Docker, error) {
	c, err := docker.NewClient(addr)
	if err != nil {
		return nil, err
	}

	//we use our own transform and client to support boot2docker tls requirements
	//@see https://github.com/boot2docker/boot2docker/issues/576
	//@see http://stackoverflow.com/questions/21562269/golang-how-to-specify-certificate-in-tls-config-for-http-client
	cas := x509.NewCertPool()
	pemData, err := ioutil.ReadFile(filepath.Join(cert, "ca.pem"))
	if err != nil {
		return nil, err
	}

	//add to pool and configrue tls
	cas.AppendCertsFromPEM(pemData)

	//load pair
	pair, err := tls.LoadX509KeyPair(filepath.Join(cert, "cert.pem"), filepath.Join(cert, "key.pem"))
	if err != nil {
		return nil, err
	}

	//create new tls config with the created ca and pair
	conf := &tls.Config{
		RootCAs:      cas,
		Certificates: []tls.Certificate{pair},
	}

	//create our own transport
	tr := &http.Transport{
		TLSClientConfig: conf,
	}

	//set docker client with new transport
	c.HTTPClient = &http.Client{Transport: tr}

	return &Docker{c, version}, nil
}

func (d *Docker) toContainerName(serviceName string) string {

	//go make an hash that is always a valid name
	h := md5.New()
	io.WriteString(h, serviceName)

	return fmt.Sprintf("%s.%s", NamePrefix, hex.EncodeToString(h.Sum(nil)))
}

func (d *Docker) RemoveAll() error {
	lopts := docker.ListContainersOptions{All: true}
	ropts := docker.RemoveContainerOptions{Force: true}

	//get all containers
	cs, err := d.client.ListContainers(lopts)
	if err != nil {
		return err
	}

	//remove all containers with the dockpit prefix
	for _, c := range cs {
		for _, name := range c.Names {
			if strings.HasPrefix(name, "/"+NamePrefix) {

				//remove the container if it matches
				ropts.ID = c.ID
				err := d.client.RemoveContainer(ropts)
				if err != nil {
					return err
				}

				//no longer care about other names
				break
			}
		}
	}

	return nil
}

func (d *Docker) Start(deps *spec.Dependencies) error {
	copts := docker.CreateContainerOptions{
		Config: &docker.Config{
			Image: fmt.Sprintf("%s:%s", ImageName, d.version),
			Cmd:   []string{"server"},
		},
	}

	sopts := &docker.HostConfig{PublishAllPorts: true}

	for sname, _ := range deps.Map {

		//set container name
		copts.Name = d.toContainerName(sname)

		//set second cmd argument to the spec to load
		copts.Config.Cmd = append(copts.Config.Cmd, sname)

		//create container
		c, err := d.client.CreateContainer(copts)
		if err != nil {
			return err
		}

		//start container
		err = d.client.StartContainer(c.ID, sopts)
		if err != nil {
			return err
		}

	}

	return nil
}
