package config

import (
	"github.com/dockpit/go-dockerclient"
)

type DependencyConfig struct {
	Name         string
	PortBindings map[docker.Port][]docker.PortBinding
}
