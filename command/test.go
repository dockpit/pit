package command

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/codegangsta/cli"

	"github.com/dockpit/exec"
)

var tmpl_test_success = `Tested successful!
`
var tmpl_test_failed = `Test failed...
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

	host := ctx.Args().First()
	if host == "" {
		return nil, nil, fmt.Errorf("Please provide the address (e.g http://localhost:8000) to test as a first argument")
	}

	//get contract
	contract, err := c.ParseExamples(ctx)
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

	//load configuration
	conf, err := c.LoadConfig(ctx)
	if err != nil {
		return nil, nil, err
	}

	//get command from configuration
	run := conf.RunConfig().Cmd

	//create command that can has timedout start& stop
	tcmd := exec.Command(run[0], run[1:]...)
	tcmd.Stdout = os.Stdout
	tcmd.Stderr = os.Stderr

	err = tcmd.StartWithTimeout(conf.RunConfig().ReadyTimeout, conf.RunConfig().ReadyExp)
	if err != nil {
		return nil, nil, err
	}

	//@todo make this configurable
	defer tcmd.StopWithTimeout(time.Second)

	//run all tests, for all resources
	res, err := contract.Resources()
	if err != nil {
		return nil, nil, err
	}

	errs := []error{}
	for _, r := range res {
		acs, err := r.Actions()
		if err != nil {
			return nil, nil, err
		}

		for _, a := range acs {
			for _, tt := range a.Tests() {

				err := tt(host, dhost, http.DefaultClient, sm, conf)
				if err != nil {
					//@todo handle failed tests better

					//gather testing errors
					errs = append(errs, err)

					fmt.Fprintf(c.out, "%s\n", err.Error())
				}
			}
		}
	}

	//we got some testing errs
	if len(errs) > 0 {
		return template.Must(template.New("test.failed").Parse(tmpl_test_failed)), errs, fmt.Errorf("Test failed")
	}

	return template.Must(template.New("test.success").Parse(tmpl_test_success)), nil, nil
}
