package command

import (
	"fmt"
	"io"
	"net/url"
	"text/template"

	"github.com/codegangsta/cli"

	"github.com/dockpit/pit/runner"
)

var tmpl_test_success = `Tested successful!
`

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
	return c.templated(c.Run)
}

func (c *Test) Run(ctx *cli.Context) (*template.Template, interface{}, error) {

	//get and parse subject url
	host := ctx.Args().First()
	if host == "" {
		return nil, nil, fmt.Errorf("Please provide the address (e.g http://localhost:8000) to test as a first argument")
	}

	hurl, err := url.Parse(host)
	if err != nil {
		return nil, nil, err
	}

	//get manifest
	m, err := c.ParseExamples(ctx)
	if err != nil {
		return nil, nil, err
	}

	//get the state manager
	sm, err := c.StateManager(ctx)
	if err != nil {
		return nil, nil, err
	}

	//fetch docker host
	dhost, _, err := c.DockerHostCertArguments(ctx)
	if err != nil {
		return nil, nil, err
	}

	durl, err := url.Parse(dhost)
	if err != nil {
		return nil, nil, err
	}

	//see if we want to run a subset
	selin := ctx.Args().Get(1)
	if selin == "" {
		selin = "*"
	}

	sel, err := runner.Parse(selin)
	if err != nil {
		return nil, nil, err
	}

	//load configuration
	conf, err := c.LoadConfig(ctx)
	if err != nil {
		return nil, nil, err
	}

	//create runner
	r, err := runner.Create("default", c.out)
	if err != nil {
		return nil, nil, err
	}

	//run tests
	err = r.Run(conf, m, sel, sm, hurl, durl)
	if err != nil {
		return nil, nil, err
	}

	return template.Must(template.New("test.success").Parse(tmpl_test_success)), nil, nil
}
