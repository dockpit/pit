dockpit cli
===========
A command line tool for (micro)service developers that allows them to define container functionality by writing examples and testing against them.

- Develop a single service in isolation by mocking all dependencies and set up the addresses as environment variables with a single command.

- When dependencies are mocked you can run the test cases quickly and reliably to make sure the service is implemented correctly.

## Warnings/Documentation
- Using `go run` in run command
	// @todo about using go run:
	// - http://stackoverflow.com/questions/24982845/process-kill-on-child-processes
	// - https://groups.google.com/forum/#!searchin/golang-nuts/interrupt$20signal/golang-nuts/nayHpf8dVxI/_QO30bLWtrcJ

## Backlog
- (major) overhaul cli output
	* test result output
	* inform when there are no examples to test against instead of just saying 'success'
	* error numbers/categories
- (major) windows support, in particular:
	* dirtar directory seperator
	* enable cross-compilation (../../docker/docker/pkg/system/stat_unsupported.go:7: undefined: syscall.Stat_t)
	* os.Interrupt signal

- (minor) panics cause test containers to be left open, which cause subsequent tests to not run
## Refactor & Cleanup
- Switch to testify/assert everywhere
- switch to a single docker client everywhere (remove dependency on docker repo pkg)


## Roadmap
- investigate what to do with services that expose multiple ports for testing
- (optional) also prevent double case names in contract package