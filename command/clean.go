package command

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/hashicorp/errwrap"

	"github.com/dockpit/pit/client"
	"github.com/dockpit/pit/model"
)

type Clean struct {
	*command
}

func NewClean() *Clean {
	return &Clean{newCommand()}
}

func (c *Clean) Name() string {
	return "clean"
}

func (c *Clean) Description() string {
	return fmt.Sprintf("<descriptions>")
}

func (c *Clean) Usage() string {
	return "<usage>"
}

func (c *Clean) Flags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{Name: "bind", Value: ":3838", Usage: "<usage>"},
		cli.StringFlag{Name: "db", Value: "dockpit.db", Usage: "<usage>"},
		cli.StringFlag{Name: "docker-host", Value: os.Getenv("DOCKER_HOST"), Usage: "<usage>"},
		cli.StringFlag{Name: "docker-cert-path", Value: os.Getenv("DOCKER_CERT_PATH"), Usage: "<usage>"},
	}
}

func (c *Clean) Action() func(ctx *cli.Context) {
	return c.command.Action(c.Run)
}

func (c *Clean) Run(ctx *cli.Context) error {
	m, err := model.NewModel(ctx.String("db"))
	if err != nil {
		return errwrap.Wrapf("Failed open model: {{err}}", err)
	}

	docker, err := client.NewDocker(m, ctx.String("docker-host"), ctx.String("docker-cert-path"))
	if err != nil {
		return errwrap.Wrapf("Failed to connect to Docker host: {{err}}", err)
	}

	err = docker.RemoveAll()
	if err != nil {
		return errwrap.Wrapf("Failed remove all running Dockpit containers: {{err}}", err)
	}

	return nil
}
