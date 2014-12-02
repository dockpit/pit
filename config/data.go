package config

//
type DependencyConfigData []string

//
type StateProviderConfigData []string

//
type ConfigData struct {
	Dependencies   map[string]*DependencyConfigData    `json:"dependencies"`
	StateProviders map[string]*StateProviderConfigData `json:"state_providers"`
}
