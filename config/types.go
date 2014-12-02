package config

type DependencyC interface{}

type StateProviderC interface{}

type C interface {
	DependencyConfigs() []DependencyC
	ProviderConfigs() []StateProviderC
}
