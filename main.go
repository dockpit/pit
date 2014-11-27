package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"

	"github.com/dockpit/pit/command"
)

var Version = "0.0.0-DEV"
var Build = "unbuild"

func main() {
	app := cli.NewApp()
	app.Name = "dockpit"
	app.Usage = "docker based microservice development environment"
	app.Version = fmt.Sprintf("%s (%s)", Version, Build)

	//specify output
	out := os.Stdout

	//create reused commands
	icmd := command.NewInstall(out)

	//init micros commands
	cmds := []command.C{
		command.NewBuild(out),
		icmd,
		command.NewMock(out, icmd),
		command.NewUnmock(out),
	}

	//append to app
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
