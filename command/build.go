package command

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dockpit/pit/reporter"

	"github.com/codegangsta/cli"
)

var BuildPart = &reporter.Build{}
var ManifestPart = &reporter.Manifest{}
var ConfigPart = &reporter.Config{}
var StatePart = &reporter.State{}

type Build struct {
	*cmd
}

func NewBuild(r reporter.R) *Build {
	return &Build{
		cmd: newCmd(r),
	}
}

func (c *Build) Name() string {
	return "build"
}

func (c *Build) Description() string {
	return fmt.Sprintf("...")
}

func (c *Build) Usage() string {
	return "Build the states using Docker"
}

func (c *Build) Flags() []cli.Flag {
	fs := []cli.Flag{}
	fs = append(fs, c.ConfigFlags()...)
	fs = append(fs, c.ParseExampleFlags()...)
	fs = append(fs, c.BuildStatesFlags()...)
	fs = append(fs, c.DockerFlags()...)

	return fs
}

func (c *Build) Action() func(ctx *cli.Context) {
	return c.toAction(c.Run)
}

func (c *Build) Run(ctx *cli.Context) error {
	c.Enter(BuildPart, BuildPart.StartingBuild)
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	mrel, err := filepath.Rel(wd, ctx.String("examples"))
	if err != nil {
		return err
	}

	//get manifest
	c.Report(ManifestPart.ParsingExamples, mrel)
	m, err := c.ParseExamples(ctx)
	if err != nil {
		return err
	}

	//get all states in the manifest
	states, err := m.States()
	if err != nil {
		return err
	}

	confrel, err := filepath.Rel(wd, ctx.String("config"))
	if err != nil {
		return err
	}

	//load configuration
	c.Report(ConfigPart.LoadingConfig, confrel)
	conf, err := c.LoadConfig(ctx)
	if err != nil {
		return err
	}

	//add default states for each provider
	for _, pc := range conf.ProviderConfigs() {
		if sts, ok := states[pc.Name()]; ok {
			states[pc.Name()] = append(sts, pc.DefaultState())
		}
	}

	staterel, err := filepath.Rel(wd, ctx.String("states"))
	if err != nil {
		return err
	}

	//get the state manager
	c.Report(StatePart.CreatingManager, staterel)
	sm, err := c.StateManager(ctx)
	if err != nil {
		return err
	}

	//loop over states and build them
	for pname, snames := range states {
		c.Enter(StatePart, StatePart.BuildingProvider, pname)
		for _, sname := range snames {
			c.Enter(StatePart, StatePart.BuildingState, sname)
			iname, err := sm.Build(pname, sname, c.Pipe())
			if err != nil {
				return err
			}

			c.Success(StatePart.BuiltState, iname)
			c.Exit()
		}
		c.Exit()
	}

	c.Exit()
	return nil
}
