# Dockpit
Painless isolation for your (micro)service development process. Comes with a minimal web UI to set up dependencies in predictive states and a CLI to quickly switch between them. It uses Docker to make each isolation predictable and portable.

## Getting Started
NOTE: You'll need to have [Docker](https://docker.com) installed and running, the official guides for setting up Docker can be found [here](https://docs.docker.com/installation/#installation). 

---
First, download any of the [prebuild binaries](https://github.com/dockpit/pit/releases/latest) or use the Go toolchain: 

```
go get github.com/dockpit/pit
```
Make sure you're in the project directory of the (micro)service you're looking to isolate and then run the `start` command:

```
cd ~/my_project
pit start
```
A text-based UI becomes available in your shell and a web interface is now being served on [http://localhost:3838](http://localhost:3838), follow on-screen instructions there to setup your first isolation.

## Todo
- [ ] Allow clearing the current isolation from text ui
- [ ] Add timeout and message when Docker is not running