package command

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/hashicorp/errwrap"

	"github.com/dockpit/pit/client"
	"github.com/dockpit/pit/command/ui"
	"github.com/dockpit/pit/model"
	"github.com/dockpit/pit/server"
)

type Start struct {
	*command
}

func NewStart() *Start {
	return &Start{newCommand()}
}

func (c *Start) Name() string {
	return "start"
}

func (c *Start) Description() string {
	return fmt.Sprintf("<descriptions>")
}

func (c *Start) Usage() string {
	return "<usage>"
}

func (c *Start) Flags() []cli.Flag {

	return []cli.Flag{
		cli.StringFlag{Name: "bind", Value: ":3838", Usage: "<usage>"},
		cli.StringFlag{Name: "db", Value: "dockpit.db", Usage: "<usage>"},
		cli.StringFlag{Name: "docker-host", Value: os.Getenv("DOCKER_HOST"), Usage: "<usage>"},
		cli.StringFlag{Name: "docker-cert-path", Value: os.Getenv("DOCKER_CERT_PATH"), Usage: "<usage>"},
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

	docker, err := client.NewDocker(m, ctx.String("docker-host"), ctx.String("docker-cert-path"))
	if err != nil {
		return errwrap.Wrapf("Failed to connect to Docker host: {{err}}", err)
	}

	svr, err := server.New(ctx.String("bind"), m, docker)
	if err != nil {
		return errwrap.Wrapf("Failed to init server: {{err}}", err)
	}

	box := ui.NewBox(m, svr, docker)

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
