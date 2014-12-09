dockpit cli
===========
A command line tool for (micro)service developers that allows them to define container functionality by writing examples and testing against them.

- Develop a single service in isolation by mocking all dependencies and set up the addresses as environment variables with a single command.

- When dependencies are mocked you can run the test cases quickly and reliably to make sure the service is implemented correctly.

## TODO
- format test output to be readable
- actually assert
	* http headers
	* request body
- centralized `.manifest` folder name and substructure 
- investigate what to do with services that expose multiple ports for testing
- windows support, in particular:
	* dirtar directory seperator

- (optional) make it easy to see files of failed test cases, maybe easily open it using pit? (e.g pit open $!)
- (optional) seperate request/response body content into seperate 'data files'
- (optional) also prevent double case names in contract package