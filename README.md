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
- verify windows support, in particular:
	* dirtar directory seperator
	* os.Interrupt signal

## Roadmap
- investigate what to do with services that expose multiple ports for testing
- (optional) also prevent double case names in contract package