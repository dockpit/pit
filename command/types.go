package command

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
)

var DBFlag = cli.StringFlag{Name: "db", Value: "dockpit.db", Usage: "The location of the Dockpit database that will be used"}
var DockerHostFlag = cli.StringFlag{Name: "docker-host", Value: os.Getenv("DOCKER_HOST"), Usage: "Must contain the address and port on which the Docker host can be reached, defaults to using the DOCKER_HOST environment variable"}
var DockerCertFlag = cli.StringFlag{Name: "docker-cert-path", Value: os.Getenv("DOCKER_CERT_PATH"), Usage: "(Optional) The directory that contains certificate information creating a secure connection to the Docker host, defaults to using the DOCKER_CERT_PATH environment variable. When left empty it will not attempt to secure the connection"}

type command struct {
	*log.Logger
}

func newCommand() *command {
	return &command{
		log.New(os.Stderr, "", log.LstdFlags),
	}
}

func (c *command) Action(fn func(c *cli.Context) error) func(ctx *cli.Context) {
	return func(ctx *cli.Context) {
		err := fn(ctx)
		if err != nil {
			c.Fatalf("[Error]: %s", err)
			return
		}
	}
}
