package command

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"

	"github.com/dockpit/deps"
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
	return fmt.Sprintf("Installs all other (micro)servicea your current .manfest depends on into your workspace, if the first argument is specified it will install only the specified (micro)service")
}

func (c *Install) Usage() string {
	return "Install (micro)services into your workspace"
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
	c.Enter(InstallPart, InstallPart.StartingInstall)
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	mrel, err := filepath.Rel(wd, ctx.String("examples"))
	if err != nil {
		return err
	}

	//get pit path
	pp := os.Getenv("PIT_PATH")
	if pp == "" {
		return fmt.Errorf("Couldn't read 'PIT_PATH' environment variabl	e, is it set?")
	}

	c.Report(InstallPart.InstallingInto, pp)
	dm := deps.NewManager(pp)

	//holds list of dependencies
	depl := make(map[string][]string)

	//do we want to install a specific package
	dep := ctx.Args().First()
	if dep != "" {

		//install specific dep
		depl[dep] = []string{}
	} else {

		//get manifest
		c.Report(ManifestPart.ParsingExamples, mrel)
		m, err := c.ParseExamples(ctx)
		if err != nil {
			return err
		}

		//retrieve all dependencies
		depl, err = m.Dependencies()
		if err != nil {
			return err
		}
	}

	//use the manager to install all dependencies into pit path
	for dep, _ := range depl {

		c.Enter(DepPart, DepPart.InstallingDep, dep)

		err := dm.Install(dep, c.Pipe())
		if err != nil {
			return err
		}

		c.Success(DepPart.InstalledDep)
		c.Exit()
	}

	c.Exit()
	return nil
}
