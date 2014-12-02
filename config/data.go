package config

//
type DependencyConfigData []string

//
type StateProviderConfigData []string

//
type ConfigData struct {
	Dependencies   map[string]*DependencyConfigData    `json:"deps"`
	StateProviders map[string]*StateProviderConfigData `json:"states"`
}
