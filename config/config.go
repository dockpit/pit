package config

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/dockpit/go-dockerclient"
)

//
type RunConfig struct {
	ReadyExp     *regexp.Regexp
	Cmd          []string
	ReadyTimeout time.Duration
	Dir          string
}

func ParseRunConfig(data *RunData) (*RunConfig, error) {
	var err error
	conf := &RunConfig{}

	if data.ReadyPattern == "" {
		data.ReadyPattern = ".*"
	}

	if data.ReadyTimeout == "" {
		data.ReadyTimeout = "100ms"
	}

	conf.ReadyTimeout, err = time.ParseDuration(data.ReadyTimeout)
	if err != nil {
		return nil, err
	}

	conf.ReadyExp, err = regexp.Compile(data.ReadyPattern)
	if err != nil {
		return nil, err
	}

	conf.Cmd = data.Command
	conf.Dir = data.Dir

	return conf, nil
}

//
type StateProviderConfig struct {
	Name string

	portBindings map[docker.Port][]docker.PortBinding
	readyExp     *regexp.Regexp
	cmd          []string
	readyTimeout time.Duration
}

func (s *StateProviderConfig) PortBindings() map[docker.Port][]docker.PortBinding {
	return s.portBindings
}
func (s *StateProviderConfig) ReadyExp() *regexp.Regexp    { return s.readyExp }
func (s *StateProviderConfig) ReadyTimeout() time.Duration { return s.readyTimeout }
func (s *StateProviderConfig) Cmd() []string               { return s.cmd }

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
	runConfig  *RunConfig
}

func Parse(cd *ConfigData) (*Config, error) {

	//@todo remove code duplication below

	//parse deps into port configs
	depsconf := []*DependencyConfig{}
	for dep, confs := range cd.Dependencies {
		portb := map[docker.Port][]docker.PortBinding{}

		//parse public, private
		for _, conf := range *confs {
			parts := strings.SplitN(conf, ":", 2)
			if len(parts) != 2 {
				return nil, fmt.Errorf("Invalid port format: '%s'", conf)
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
	for pname, conf := range cd.StateProviders {
		portb := map[docker.Port][]docker.PortBinding{}

		//parse public, private ports
		for _, pconf := range conf.Ports {

			parts := strings.SplitN(pconf, ":", 2)
			if len(parts) != 2 {
				return nil, fmt.Errorf("Invalid port format: '%s'", pconf)
			}

			portb[docker.Port(parts[0]+"/tcp")] = []docker.PortBinding{docker.PortBinding{HostPort: parts[1]}}
		}

		//get regexp to determine when the state provider is ready
		exp, err := regexp.Compile(conf.ReadyPattern)
		if err != nil {
			return nil, err
		}

		//get timeout by parsing duration
		if conf.ReadyTimeout == "" {
			return nil, fmt.Errorf("No 'ready_timeout' for provider: %s", pname)
		}

		d, err := time.ParseDuration(conf.ReadyTimeout)
		if err != nil {
			return nil, err
		}

		spconf = append(spconf, &StateProviderConfig{
			Name: pname,

			portBindings: portb,
			readyExp:     exp,
			cmd:          conf.Cmd,
			readyTimeout: d,
		})
	}

	//parse data about running the tested service, default to empty struct
	if cd.Run == nil {
		cd.Run = &RunData{}
	}

	rconf, err := ParseRunConfig(cd.Run)
	if err != nil {
		return nil, err
	}

	return &Config{cd, depsconf, spconf, rconf}, nil
}

func (c *Config) RunConfig() *RunConfig {
	return c.runConfig
}

func (c *Config) StateProviderConfig(pname string) StateProviderC {
	for _, spc := range c.spConfigs {
		if spc.Name == pname {
			return spc
		}
	}

	return nil
}

func (c *Config) PortBindingsForState(pname string) map[docker.Port][]docker.PortBinding {
	res := map[docker.Port][]docker.PortBinding{}
	for _, spc := range c.spConfigs {
		if spc.Name == pname {
			return spc.portBindings
		}
	}

	return res
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
