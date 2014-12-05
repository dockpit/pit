dockpit cli
===========

Ahoy Matey! Welcome to the _dockpit_, the dockpit is a cockpit for your docker based microservice architecture. From the comfort of your command line it will be possible to rapidly develop individuel services without the hassle that normally comes with microservice development, it is not [a free lunch](http://highscalability.com/blog/2014/4/8/microservices-not-a-free-lunch.html) after all.

Dockpit has one goal: __Making you 10x more productive while developing micro-services__ , and achieves this by providing two functionalities:

1. `pit mock` - Mock the dependencies for the current service based on the specification in the current directory and create an isolated environment using docker (based on the DOCKER_HOST env variable).
2. `pit test` - Using the isolated environment to black-box test the current service against the specification.

Annotating type
-----------------

```javascript

//lets hope the peg and hole are both round
var plank = new Plank(peg, hole)
```


```go
type Round interface{}

//peg and hole can implement the Round interface by just implementing it
//not by specifying it explicetly: no 'implements' keyword
plank := NewPlank(peg Round, hole Round)
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
