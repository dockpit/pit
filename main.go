package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

var Version = "0.0.0-DEV"
var Build = "unbuild"

func main() {
	app := cli.NewApp()
	app.Name = "dockpit"
	app.Usage = "docker based microservice development environment"
	app.Version = fmt.Sprintf("%s (%s)", Version, Build)

	app.Run(os.Args)
}
