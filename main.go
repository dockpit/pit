package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"

	"github.com/dockpit/pit/command"
)

type Command interface {
	Name() string
	Description() string
	Usage() string
	Run(c *cli.Context) error
	Action() func(ctx *cli.Context)
	Flags() []cli.Flag
}

var Version = "0.0.0"
var Build = "gobuild"

func main() {
	app := cli.NewApp()
	app.Author = "Ad van der Veer"
	app.Email = "ad@dockpit.io"
	app.Name = "pit"
	app.Usage = "develop your (micro)service in isolation"
	app.Version = fmt.Sprintf("%s (%s)", Version, Build)

	cmds := []Command{
		command.NewStart(Version),
		command.NewClean(),
	}

	for _, c := range cmds {
		app.Commands = append(app.Commands, cli.Command{
			Name:        c.Name(),
			Usage:       c.Usage(),
			Action:      c.Action(),
			Description: c.Description(),
			Flags:       c.Flags(),
		})
	}

	app.Run(os.Args)
}
