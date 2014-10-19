package command

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/fsouza/go-dockerclient"

	"github.com/dockpit/pit/spec"
)

type Docker struct {
	client *docker.Client
}

func NewDocker(addr string, cert string) (*Docker, error) {
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

	return &Docker{c}, nil
}

func (d *Docker) StopAll() error {

	//@todo implement

	return nil
}

func (d *Docker) Start(deps *spec.Dependencies) error {
	opts := docker.CreateContainerOptions{
		Config: &docker.Config{},
	}

	for sname, eps := range deps.Map {

		//customize options
		opts.Name = sname

		_ = eps

		//@todo implement

	}

	return nil
}
