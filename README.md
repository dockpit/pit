dockpit cli
===========

Ahoy Matey! Welcome to the _dockpit_, the dockpit is a cockpit for your docker based microservice architecture. From the comfort of your command line it will be possible to rapidly develop individuel services without the hassle that normally comes with microservice development, it is not [a free lunch](http://highscalability.com/blog/2014/4/8/microservices-not-a-free-lunch.html) after all.

Dockpit has one goal: __Making you 10x more productive while developing micro-services__ , and achieves this by providing two functionalities:

1. `pit mock` - Mock the dependencies for the current service based on the specification in the current directory and create an isolated environment using docker (based on the DOCKER_HOST env variable).
2. `pit test` - Using the isolated environment to black-box test the current service against the specification.

the specification
-----------------
At the basis is a JSON specification file called `dockpit.json` that should be present in the directory of the current service. The file describes a set of JSON over HTTP Rest calls as test cases, all different cases together should describe the service's behaviour.

Below is an imaginary 'notes' endpoint that supports the listing, creation and deletion of notes:
```
{
	"endpoints": [
		{
			"name": "delete_notes",
			"cases": [
				{
					//'given' describes the state of the service before sending the request, for example by loading data in the database 
					"given": {

					},
					//behaviour describing cases, 'when' specifies the request that the service expects
					"when": {
						"method": "DELETE",
						"path": "/notes"
					},
					//'then' describes the response the service is to return for the request
					"then": {
						"status_code": 200
					},
					//'while' lists other service that should be involved in the formulation of the
					"while": [
						{"service": "users"}
					]
				}
			]
		}
	]
}
```
JSON Schema&Specification 
-----------------
- http://orderly-json.org/ (http://orderly-json.org/)
- http://json-schema.org/ (https://github.com/xeipuuv/gojsonschema)
- https://github.com/apiaryio/gavel (https://github.com/apiaryio/gavel.js)


Adjacent products
-----------------

- fig:
- tug: https://news.ycombinator.com/item?id=8498804
- crane: https://github.com/michaelsauter/crane
