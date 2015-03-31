# Dockpit
Painless dependency handling for your (micro)service development process: 

- Rapidly **configure and populate** databases, message queues and service registries through a Web UI and switch between different setups from the comfort of your shell.
- **mock the HTTP(S) services** you depend on to develop your own (micro)service in isolation, reliably.
- **Share setups with you’re team members** by simply commiting the *Dockpit* configuration file to the project’s version control.
- Build for *Docker*, written in *Go* and compiled into a single *~10mb* binary that works on **OSX, Windows and Linux**.

## Getting Started
1. First, [download Dockpit](https://dockpit.io/download.html) as any of the prebuild binaries or use the Go toolchain: `go get -u github.com/dockpit/pit`

3. Make sure the `pit` binary is somewhere in your PATH and that you're in the directory of the (micro)service you're looking to isolate, then simply run the `start` command:

	```
cd ~/my_project
pit start
```
3. A text-based UI becomes available in your shell and a web interface is now being served on [http://localhost:3838](http://localhost:3838), follow on-screen instructions there to setup your first dependency.

	---
NOTE: If something went wrong make sure [Docker](https://docker.com) is installed and running, the official guides for setting up Docker can be found [here](https://docs.docker.com/installation/#installation). 
	
	---

## Use Cases
- *Quickly test alternative technologies* - RabbitMQ or ZeroMQ? Dockpit makes it simple to experiment with different dependencies so you can make more informed decisions. 
- *Different data for different scenarios* - Want to test the functionality of your (micro)service with or without a populated database? Dockpit makes it easy to setup both scenarios and quickly switch between them.  
- *Migration to a new database version of schema* - Dockpit makes it trivial to test (micro)service functionaly across different database or schema versions.
- *Testing failing dependencies* - Dockpit allows you to test your (micro)service functions as expected when a dependency fails.

## Dockpit vs Other Software
- **Fig, Docker Compose**: Fig and Docker Compose use a configuration file for defining a single set of dependencies. Dockpit instead provides an easy-to-use UI and focusses on creating multiple sets of dependencies with well defined states that can be switched out when needed.  
  
- **Kitamatic** comes with an UI that makes it easy install Docker and launch images from the hub. *Dockpit* expects Docker to be running already and its UI focusses on constructing Docker images specific to your project. As such, it complements Kitamatic and can be configured to use the Docker host Kitamatic installs.

- **Vagrant, Virtual Box, VMWare, etc**: Virtual machines are often used to provide reliable development environments for teams. *Dockpit* aims to be more lightweight by depending on Docker as its portable runtime. But Dockpit itself can ofcourse also be run inside such virtual machines to provide the best of both worlds. 

- **npm, maven, pip, etc**: Language stacks often come with there own tools for managing code dependencies. *Dockpit* instead focusses on managing _network_ dependencies (other processes) such as databases, message queues, HTTP APIs, etc.

## Why?
_Dockpit_ was created after user research on the development of (micro)services showed the __difficulty of managing the dependencies of a single service__: HTTP(s) API responses, databases, message ques, service registries etc. They all need to be setup and respond predicatable in order to develop and test features that depend on them. Currently, this is often done through fragile shell scripts or dedicated internal tooling, both requiring significant effort to maintain. _Dockpit_ was created as an alternative to make managing these dependencies quick, reliable and, dare I say, __fun__.
