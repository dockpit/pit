package command

import (
	"fmt"
	"net/url"
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
	host, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}

	//change to http connection
	host.Scheme = "https"

	c, err := docker.NewTLSClient(host.String(), filepath.Join(cert, "cert.pem"), filepath.Join(cert, "key.pem"), filepath.Join(cert, "ca.pem"))
	if err != nil {
		return nil, err
	}

	return &Docker{c, version}, nil
}

func (d *Docker) toContainerName(serviceName string) string {
	return fmt.Sprintf("%s.%s", NamePrefix, serviceName)
}

func (d *Docker) RemoveAll() error {
	lopts := docker.ListContainersOptions{}
	ropts := docker.RemoveContainerOptions{Force: true}

	//get all containers
	cs, err := d.client.ListContainers(lopts)
	if err != nil {
		return err
	}

	// remove all containers with the dockpit prefix
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

		//assume link
		link, err := spec.NewLink(sname)
		if err != nil {
			return err
		}

		//set container name
		copts.Name = d.toContainerName(link.Name())

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
