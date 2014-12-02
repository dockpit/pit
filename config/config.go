package config

//
type StateProviderConfig struct {
	Name      string
	PortSpecs []string
}

//
type DependencyConfig struct {
	Name      string
	PortSpecs []string
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
		depsconf = append(depsconf, &DependencyConfig{
			Name:      dep,
			PortSpecs: *confs,
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
