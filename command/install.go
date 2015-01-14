package command

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"

	"github.com/dockpit/debs"
	"github.com/dockpit/pit/reporter"
)

type Install struct {
	*cmd
}

func NewInstall(r reporter.R) *Install {
	return &Install{
		cmd: newCmd(r),
	}
}

func (c *Install) Name() string {
	return "install"
}

func (c *Install) Description() string {
	return fmt.Sprintf("Parse examples into a manifest and extract all dependencies. Install each dependency into the Dockpit workspace which is expected to be specified as the PIT_PATH environment variable")
}

func (c *Install) Usage() string {
	return "Install dependencies"
}

func (c *Install) Flags() []cli.Flag {
	fs := []cli.Flag{}
	fs = append(fs, c.ParseExampleFlags()...)

	return fs
}

func (c *Install) Action() func(ctx *cli.Context) {
	return c.toAction(c.Run)
}

func (c *Install) Run(ctx *cli.Context) error {

	//get pit path
	pp := os.Getenv("PIT_PATH")
	if pp == "" {
		return fmt.Errorf("Couldn't read 'PIT_PATH' environment variable, is it set?")
	}

	//get manifest
	m, err := c.ParseExamples(ctx)
	if err != nil {
		return err
	}

	//retrieve all dependencies
	deps, err := m.Dependencies()
	if err != nil {
		return err
	}

	//use the manager to install all dependencies into pit path
	dm := debs.NewManager(pp)
	installs := map[string]string{}
	for dep, _ := range deps {
		fmt.Fprintf(c.Pipe(), "Installing %s:\n", dep)

		err := dm.Install(dep)
		if err != nil {
			fmt.Fprintf(c.Pipe(), "ERROR \n")
			return err
		}

		//add location installation
		installs[dep], err = dm.Locate(dep)
		if err != nil {
			return err
		}

		fmt.Fprintf(c.Pipe(), "done!\n\n")
	}

	return nil
}
