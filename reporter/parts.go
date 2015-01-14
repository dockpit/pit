package reporter

import (
	"fmt"
)

//Installation reporting
type Install struct{}

func (i Install) ID() string { return "install" }

func (i Install) StartingInstall(args ...interface{}) (int, string) {
	return 1, fmt.Sprintf("Starting installation")
}

func (i Install) InstallingInto(args ...interface{}) (int, string) {
	return 1, fmt.Sprintf("Installing dependencies into '%s'", args[0])
}

//Build reporting
type Build struct{}

func (b Build) ID() string { return "build" }

func (b Build) StartingBuild(args ...interface{}) (int, string) {
	return 1, fmt.Sprintf("Starting build")
}

//Manifest Reporting
type Manifest struct{}

func (m Manifest) ID() string { return "manifest" }
func (m Manifest) ParsingExamples(args ...interface{}) (int, string) {
	return 1, fmt.Sprintf("Parsing examples from './%s'", args[0])
}

//Reporting
type Dep struct{}

func (d Dep) ID() string { return "dep" }

func (d Dep) InstallingDep(args ...interface{}) (int, string) {
	return 1, fmt.Sprintf("Installing dependency: '%s'", args[0])
}

func (d Dep) InstalledDep(args ...interface{}) (int, string) {
	return 1, fmt.Sprintf("Installed successfully")
}

//
type Config struct{}

func (c Config) LoadingConfig(args ...interface{}) (int, string) {
	return 1, fmt.Sprintf("Loading configuration from './%s'", args[0])
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
	return 1, fmt.Sprintf("Creating manager for states in './%s'", args[0])
}
