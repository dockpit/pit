package command

import (
	"fmt"
	"io"
	"net/url"
	"strings"

	"github.com/codegangsta/cli"

	"github.com/dockpit/pit/runner"
)

type Test struct {
	*cmd
}

func NewTest(out io.Writer) *Test {
	return &Test{
		cmd: newCmd(out),
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

	//get and parse subject url
	host := ctx.Args().First()
	if host == "" {
		return fmt.Errorf("Please provide the address (e.g http://localhost:8000) to test as a first argument")
	}

	hurl, err := url.Parse(host)
	if err != nil {
		return err
	}

	//get manifest
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
	selin := ctx.Args().Get(1)
	if selin == "" {
		selin = "*"
	}

	sel, err := runner.Parse(selin)
	if err != nil {
		return err
	}

	//load configuration
	conf, err := c.LoadConfig(ctx)
	if err != nil {
		return err
	}

	//set the special docker_host configuration
	//@todo this is pretty hackish, but use env to set
	//docker hostname (ip), possibly move this to command types
	cdata := conf.Data()
	cdata.DockerHostname = strings.SplitN(durl.Host, ":", 2)[0]

	//create runner
	r, err := runner.Create("default", c.out)
	if err != nil {
		return err
	}

	//run tests
	err = r.Run(conf, m, sel, sm, hurl, durl)
	if err != nil {
		return err
	}

	return nil
}
