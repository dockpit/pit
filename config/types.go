package config

import (
	"regexp"
	"time"

	"github.com/dockpit/go-dockerclient"
)

var ConfigFile = "dockpit.json"

type DependencyC interface{}

type StateProviderC interface {
	PortBindings() map[docker.Port][]docker.PortBinding
	ReadyExp() *regexp.Regexp
	ReadyTimeout() time.Duration
	Cmd() []string
}

type C interface {
	DependencyConfigs() []DependencyC
	ProviderConfigs() []StateProviderC
	PortBindingsForDep(dep string) map[docker.Port][]docker.PortBinding
	StateProviderConfig(pname string) StateProviderC
	RunConfig() *RunConfig
}
