package reporter

import (
	"fmt"
)

//
type Build struct{}

func (b Build) ID() string { return "build" }

func (b Build) StartingBuild(args ...interface{}) (int, string) {
	return 1, fmt.Sprintf("Starting build")
}

//
type Manifest struct{}

func (m Manifest) ID() string { return "manifest" }
func (m Manifest) ParsingExamples(args ...interface{}) (int, string) {
	return 1, fmt.Sprintf("Parsing examples")
}

//
func (m Manifest) RetrievingStates(args ...interface{}) (int, string) {
	return 1, fmt.Sprintf("Retrieving states")
}

//
type Config struct{}

func (c Config) LoadingConfig(args ...interface{}) (int, string) {
	return 1, fmt.Sprintf("Loading configuration")
}

func (c Config) ID() string { return "config" }

//
type State struct{}

func (s State) ID() string { return "state" }

func (s State) BuildingProvider(args ...interface{}) (int, string) {
	return 1, fmt.Sprintf("Building states for provider: '%s'", args[0])
}

func (s State) BuildingState(args ...interface{}) (int, string) {
	return 1, fmt.Sprintf("Building state: '%s'", args[0])
}

func (s State) BuiltState(args ...interface{}) (int, string) {
	return 1, fmt.Sprintf("Successfully built state image: '%s'", args[0])
}

func (s State) CreatingManager(args ...interface{}) (int, string) {
	return 1, fmt.Sprintf("Creating state manager")
}
