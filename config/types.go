package config

import (
	"regexp"
	"time"
)

var ConfigFile = "dockpit.json"

type DependencyC interface{}

type StateProviderC interface {
	Name() string
	Ports() []*PortConfig
	ReadyExp() *regexp.Regexp
	ReadyTimeout() time.Duration
	DefaultState() string
	Cmd() []string
}

type C interface {
	Subject() string
	Data() *ConfigData
	DependencyConfigs() []DependencyC
	RunConfig() *RunConfig
	ProviderConfigs() []StateProviderC
	PortsForDependency(dep string) []*PortConfig
	PortsForStateProvider(pname string) []*PortConfig
	StateProviderConfig(pname string) StateProviderC
}
