package command

import (
	"fmt"

	"github.com/codegangsta/cli"
)

type Init struct {
	*command
}

func NewInit() *Init {
	return &Init{newCommand()}
}

func (c *Init) Name() string {
	return "init"
}

func (c *Init) Description() string {
	return fmt.Sprintf("Initializes Dockpit into the current direcotry")
}

func (c *Init) Usage() string {
	return "Initialize Dockpit in the current Directory"
}

func (c *Init) Flags() []cli.Flag {
	return []cli.Flag{
		DBFlag,
		DockerHostFlag,
		DockerCertFlag,
	}
}

func (c *Init) Action() func(ctx *cli.Context) {
	return c.command.Action(c.Run)
}

func (c *Init) Run(ctx *cli.Context) error {

	fmt.Println("Hello world")

	return nil
}
