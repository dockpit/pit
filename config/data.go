package config

//
type DependencyConfigData []string

//
type StateProviderConfigData struct {
	Ports        []string `json:"ports"`
	ReadyPattern string   `json:"ready_pattern"`
	ReadyTimeout string   `json:"ready_timeout"`
	Cmd          []string `json:"command"`
}

//
type RunData struct {
	Command []string `json:"command"`
}

//
type ConfigData struct {
	Dependencies   map[string]*DependencyConfigData    `json:"deps"`
	StateProviders map[string]*StateProviderConfigData `json:"states"`
	Run            *RunData                            `json:"run"`
}
