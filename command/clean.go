package command

import (
	"fmt"

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
	return fmt.Sprintf("Sometimes Docker containers remain after testing, this command removes all containers (and the data in them) that begin their name with 'dp-'. It does not remove any of your Docker images.")
}

func (c *Clean) Usage() string {
	return "Remove all Dockpit containers"
}

func (c *Clean) Flags() []cli.Flag {
	return []cli.Flag{
		DBFlag,
		DockerHostFlag,
		DockerCertFlag,
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
