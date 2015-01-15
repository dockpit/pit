package reporter

import (
	"fmt"
)

//Installation reporting
type Install struct{}

func (i Install) ID() string { return "install" }

func (i Install) StartingInstall(args ...interface{}) (int, string) {
	return 1, fmt.Sprintf("Starting dependency installation...")
}

func (i Install) InstallingInto(args ...interface{}) (int, string) {
	return 1, fmt.Sprintf("Installing dependencies into '%s'", args[0])
}

//Mock reporting
type Mock struct{}

func (m Mock) ID() string { return "mock" }
func (m Mock) StartingMocks(args ...interface{}) (int, string) {
	return 1, fmt.Sprintf("\nStarting dependency mocks...")
}

func (m Mock) StoppingMocks(args ...interface{}) (int, string) {
	return 1, fmt.Sprintf("\nStopping dependency mocks...")
}

func (m Mock) MockingFrom(args ...interface{}) (int, string) {
	return 1, fmt.Sprintf("Mocking dependencies from '%s'", args[0])
}

//Building reporting
type Build struct{}

func (b Build) ID() string { return "build" }

func (b Build) StartingBuild(args ...interface{}) (int, string) {
	return 1, fmt.Sprintf("Starting state builds...")
}

//Manifest Reporting
type Manifest struct{}

func (m Manifest) ID() string { return "manifest" }
func (m Manifest) ParsingExamples(args ...interface{}) (int, string) {
	return 1, fmt.Sprintf("Parsing examples from './%s'", args[0])
}

//Reports about testing
type Test struct{}

func (d Test) ID() string { return "test" }

func (d Test) TestingCase(args ...interface{}) (int, string) {
	return 1, fmt.Sprintf("%s %s '%s'", args...)
}

func (d Test) FailedCase(args ...interface{}) (int, string) {
	return 1, fmt.Sprintf("Failed - %s", args...)
}

func (d Test) SkippedCase(args ...interface{}) (int, string) {
	return 1, fmt.Sprintf("%s %s '%s' (skipped)", args...)
}

func (d Test) TestedCase(args ...interface{}) (int, string) {
	return 1, fmt.Sprintf("Pass")
}

func (d Test) TestingResource(args ...interface{}) (int, string) {
	return 1, fmt.Sprintf("%s", args[0])
}

func (d Test) StartingTests(args ...interface{}) (int, string) {
	return 1, fmt.Sprintf("\nStarting tests...")
}

func (d Test) SomeTestsSkipped(args ...interface{}) (int, string) {
	res := args[0].(*Result)
	return 1, fmt.Sprintf("%d tests skipped", res.Skipped)
}

func (d Test) SomeTestsFailed(args ...interface{}) (int, string) {
	res := args[0].(*Result)
	return 1, fmt.Sprintf("%d tests failed", res.Failed)
}

func (d Test) SomeTestsPassed(args ...interface{}) (int, string) {
	res := args[0].(*Result)
	return 1, fmt.Sprintf("%d tests passed", res.Succeeded)
}

func (d Test) AllTestsPassed(args ...interface{}) (int, string) {
	return 1, fmt.Sprintf("All tests passed!")
}

//Reporting about a dependency
type Dep struct{}

func (d Dep) ID() string { return "dep" }

func (d Dep) InstallingDep(args ...interface{}) (int, string) {
	return 1, fmt.Sprintf("Installing dependency: '%s'", args[0])
}

func (d Dep) InstalledDep(args ...interface{}) (int, string) {
	return 1, fmt.Sprintf("Dependency Installed successfully")
}

func (d Dep) MockingDep(args ...interface{}) (int, string) {
	return 1, fmt.Sprintf("Mocking dependency: '%s'", args[0])
}

func (d Dep) MockedDep(args ...interface{}) (int, string) {
	return 1, fmt.Sprintf("Dependency Mocked successfully, running at: '%s'", args[0])
}

func (d Dep) UnmockingDep(args ...interface{}) (int, string) {
	return 1, fmt.Sprintf("Stopping mock: '%s'", args[0])
}

func (d Dep) UnmockedDep(args ...interface{}) (int, string) {
	return 1, fmt.Sprintf("Mock stopped successfully")
}

//Reporting about a configuration
type Config struct{}

func (c Config) LoadingConfig(args ...interface{}) (int, string) {
	return 1, fmt.Sprintf("Loading configuration from './%s'", args[0])
}

func (c Config) ParsingSelector(args ...interface{}) (int, string) {
	return 1, fmt.Sprintf("Parsing test case selector: '%s'", args[0])
}

func (c Config) SettingDockerHostname(args ...interface{}) (int, string) {
	return 1, fmt.Sprintf("Setting .DockerHostname is set to: '%s'", args[0])
}

func (c Config) ID() string { return "config" }

//reporting about a state
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

//reporting rerrors
type Error struct{}

func (err Error) ID() string { return "error" }

func (err Error) ThrowError(args ...interface{}) (int, string) {
	return 1, fmt.Sprintf("Error: %s", args...)
}
