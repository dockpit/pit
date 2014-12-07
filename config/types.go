package config

import (
	"github.com/dockpit/go-dockerclient"
)

type DependencyC interface{}

type StateProviderC interface{}

type C interface {
	DependencyConfigs() []DependencyC
	ProviderConfigs() []StateProviderC

	PortBindingsForDep(dep string) map[docker.Port][]docker.PortBinding
	PortBindingsForState(pname string) map[docker.Port][]docker.PortBinding
}
