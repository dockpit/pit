package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"

	"github.com/dockpit/pit/command"
	"github.com/dockpit/pit/reporter"
)

var Version = "0.0.0-DEV"
var Build = "unbuild"

func main() {
	app := cli.NewApp()
	app.Name = "dockpit"
	app.Usage = "docker based microservice development environment"
	app.Version = fmt.Sprintf("%s (%s)", Version, Build)

	//create the reporter
	r := reporter.NewTerminal(os.Stdout)

	//create reused commands
	icmd := command.NewInstall(r)

	//init micros commands
	cmds := []command.C{
		command.NewBuild(r),
		icmd,
		command.NewMock(r, icmd),
		command.NewUnmock(r),
		command.NewTest(r),
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
