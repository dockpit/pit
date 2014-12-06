package config

import (
	"fmt"
	"strings"

	"github.com/fsouza/go-dockerclient"
)

//
type StateProviderConfig struct {
	Name      string
	PortSpecs []string
}

//
type DependencyConfig struct {
	Name         string
	PortBindings map[docker.Port][]docker.PortBinding
}

//
type Config struct {
	data *ConfigData

	depConfigs []*DependencyConfig
	spConfigs  []*StateProviderConfig
}

func NewConfig(cd *ConfigData) (*Config, error) {

	//parse deps into port configs
	depsconf := []*DependencyConfig{}
	for dep, confs := range cd.Dependencies {
		portb := map[docker.Port][]docker.PortBinding{}

		//@todo parse private public
		for _, conf := range *confs {
			parts := strings.SplitN(conf, ":", 2)
			if len(parts) != 2 {
				return nil, fmt.Errorf("Invalid port format '%s'", conf)
			}

			portb[docker.Port(parts[0]+"/tcp")] = []docker.PortBinding{docker.PortBinding{HostPort: parts[1]}}
		}

		depsconf = append(depsconf, &DependencyConfig{
			Name:         dep,
			PortBindings: portb,
		})
	}

	//parse state provider config
	spconf := []*StateProviderConfig{}
	for pname, confs := range cd.StateProviders {
		spconf = append(spconf, &StateProviderConfig{
			Name:      pname,
			PortSpecs: *confs,
		})
	}

	return &Config{cd, depsconf, spconf}, nil
}

func (c *Config) PortBindingsForDep(dep string) map[docker.Port][]docker.PortBinding {
	res := map[docker.Port][]docker.PortBinding{}
	for _, depc := range c.depConfigs {
		if depc.Name == dep {
			return depc.PortBindings
		}
	}

	return res
}

func (c *Config) DependencyConfigs() []DependencyC {
	res := []DependencyC{}

	for _, depc := range c.depConfigs {
		res = append(res, depc)
	}

	return res
}
func (c *Config) ProviderConfigs() []StateProviderC {
	res := []StateProviderC{}

	for _, spc := range c.spConfigs {
		res = append(res, spc)
	}

	return res
}
