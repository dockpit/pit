dockpit cli
===========
A command line tool for (micro)service developers that allows them to define container functionality by writing examples and testing against them.

- Develop a single service in isolation by mocking all dependencies and set up the addresses as environment variables with a single command.

- When dependencies are mocked you can run the test cases quickly and reliably to make sure the service is implemented correctly.
	

## Backlog	
- verify windows support, in particular:
	* dirtar directory seperator
	* os.Interrupt signal
	* Child processes interruption

## Roadmap
- investigate what to do with services that expose multiple ports for testing
- (optional) also prevent double case names in contract package