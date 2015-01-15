package command

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/codegangsta/cli"

	"github.com/dockpit/pit/reporter"
	"github.com/dockpit/pit/runner"
)

type Test struct {
	*cmd
}

func NewTest(r reporter.R) *Test {
	return &Test{
		cmd: newCmd(r),
	}
}

func (c *Test) Name() string {
	return "test"
}

func (c *Test) Description() string {
	return fmt.Sprintf("...")
}

func (c *Test) Usage() string {
	return "the the implementation against current examples"
}

func (c *Test) Flags() []cli.Flag {
	fs := []cli.Flag{}

	fs = append(fs, c.ConfigFlags()...)
	fs = append(fs, c.ParseExampleFlags()...)
	fs = append(fs, c.BuildStatesFlags()...)

	return fs
}

func (c *Test) Action() func(ctx *cli.Context) {
	return c.toAction(c.Run)
}

func (c *Test) Run(ctx *cli.Context) error {
	c.Enter(TestPart, TestPart.StartingTests)
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

	//get the state manager
	sm, err := c.StateManager(ctx)
	if err != nil {
		return err
	}

	//fetch docker host
	dhost, _, err := c.DockerHostCertArguments(ctx)
	if err != nil {
		return err
	}

	durl, err := url.Parse(dhost)
	if err != nil {
		return err
	}

	//see if we want to run a subset
	selin := ctx.Args().First()
	if selin == "" {
		selin = "*"
	}

	c.Report(ConfigPart.ParsingSelector, selin)
	sel, err := runner.Parse(selin)
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

	hurl, err := url.Parse(conf.Subject())
	if err != nil {
		return err
	}

	//set the special docker_host configuration
	//@todo this is pretty hackish, but use env to set
	//docker hostname (ip), possibly move this to command types
	dhostname := strings.SplitN(durl.Host, ":", 2)[0]
	c.Report(ConfigPart.SettingDockerHostname, dhostname)
	cdata := conf.Data()
	cdata.DockerHostname = dhostname

	//create runner
	r, err := runner.Create("default", c)
	if err != nil {
		return err
	}

	//run tests
	res, err := r.Run(conf, m, sel, sm, hurl, durl)
	if err != nil {
		return err
	}

	//report results
	if res.Failed > 0 {
		c.Error(TestPart.SomeTestsFailed, res)
		c.SetStatusCode(2)
	} else {
		if res.Skipped > 0 {
			c.Warning(TestPart.SomeTestsSkipped, res)
			c.Success(TestPart.SomeTestsPassed, res)
		} else {
			c.Success(TestPart.AllTestsPassed, res)
		}
	}

	c.Exit()
	return nil
}
