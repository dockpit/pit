package command

import (
	"github.com/fsouza/go-dockerclient"

	"github.com/dockpit/pit/spec"
)

type Docker struct {
	client *docker.Client
}

func NewDocker(addr string) (*Docker, error) {
	c, err := docker.NewClient(addr)
	if err != nil {
		return nil, err
	}

	return &Docker{c}, nil
}

func (d *Docker) StopAll() error {

	return nil
}

func (d *Docker) Start(deps *spec.Dependencies) error {

	//@todo implement

	return nil
}
