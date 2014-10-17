dockpit cli
===========

Ahoy Matey! Welcome to the _dockpit_, the dockpit is a cockpit for your docker based microservice architecture. From the comfort of your command line it will be possible to rapidly develop individuel services without the hassle that normally comes with microservice development, it is not [a free lunch](http://highscalability.com/blog/2014/4/8/microservices-not-a-free-lunch.html) after all.

The CLI focusses on three things and three things only:

1. `pit switch <name>` - Start or continue development on a (different) named service.
2. `pit mock` - Mock the dependencies for the current service based on the specification and create an isolated environment.
3. `pit test` - Using the isolated environment test the current service using the specification.
