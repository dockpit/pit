package command

import (
	"fmt"
	"io"

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

func NewBuild(out io.Writer) *Build {
	return &Build{
		cmd: newCmd(out),
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

	//get manifest
	c.Report(ManifestPart.ParsingExamples)
	m, err := c.ParseExamples(ctx)
	if err != nil {
		return err
	}

	//get all states in the manifest
	c.Report(ManifestPart.RetrievingStates)
	states, err := m.States()
	if err != nil {
		return err
	}

	//load configuration
	c.Report(ConfigPart.LoadingConfig)
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

	//get the state manager
	c.Report(StatePart.CreatingManager)
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
