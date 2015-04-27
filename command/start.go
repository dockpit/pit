package command

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/hashicorp/errwrap"

	"github.com/dockpit/pit/client"
	"github.com/dockpit/pit/command/ui"
	"github.com/dockpit/pit/host"
	"github.com/dockpit/pit/model"
	"github.com/dockpit/pit/server"
)

type Start struct {
	*command
	version string
}

func NewStart(v string) *Start {
	return &Start{newCommand(), v}
}

func (c *Start) Name() string {
	return "start"
}

func (c *Start) Description() string {
	return fmt.Sprintf("Starts serving the web UI over HTTP and launches the terminal UI in your current shell. It will create a new dockpit.db if it doesn't exists yet and connect to the configured Docker host for creating images and running containers.")
}

func (c *Start) Usage() string {
	return "Start the development interface"
}

func (c *Start) Flags() []cli.Flag {

	return []cli.Flag{
		cli.StringFlag{Name: "bind", Value: ":3838", Usage: "The address and port on which the process will listen to serve the web UI"},
		DBFlag,
		DockerHostFlag,
		DockerCertFlag,
	}
}

func (c *Start) Action() func(ctx *cli.Context) {
	return c.command.Action(c.Run)
}

func (c *Start) Run(ctx *cli.Context) error {
	m, err := model.NewModel(ctx.String("db"))
	if err != nil {
		return errwrap.Wrapf("Failed to start: {{err}}", err)
	}

	lhost := host.NewLocal(ctx.String("docker-host"), ctx.String("docker-cert-path"))
	dphost := host.NewDockpit(lhost)

	err = dphost.Init(ctx)
	if err != nil {
		return errwrap.Wrapf("Failed to initialize Docker host connection: {{err}}", err)
	}

	docker, err := client.NewDocker(m, dphost)
	if err != nil {
		return errwrap.Wrapf("Failed to connect to Docker host: {{err}}", err)
	}

	svr, err := server.New(c.version, ctx.String("bind"), m, docker)
	if err != nil {
		return errwrap.Wrapf("Failed to init server: {{err}}", err)
	}

	store, err := ui.NewStore(m, docker)
	if err != nil {
		return errwrap.Wrapf("Failed to create store: {{err}}", err)
	}

	box := ui.NewBox(m, svr, docker, store)

	err = box.Init()
	if err != nil {
		return errwrap.Wrapf("Failed to start: {{err}}", err)
	}

	go func() {
		err := svr.Start()
		if err != nil {
			c.Fatal(errwrap.Wrapf("Server failed: {{err}}", err))
		}
	}()

	err = box.Run()
	if err != nil {
		return errwrap.Wrapf("Termbox failed: {{err}}", err)
	}

	err = svr.Stop()
	if err != nil {
		return errwrap.Wrapf("Failed to stop server: {{err}}", err)
	}

	return nil
}
