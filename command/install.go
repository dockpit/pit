package command

import (
	"fmt"
	"io"
	"os"
	"text/template"

	"github.com/codegangsta/cli"

	"github.com/dockpit/debs"
)

var tmpl_install = `Installation successful!`

type Install struct {
	*cmd
}

func NewInstall(out io.Writer) *Install {
	return &Install{
		cmd: newCmd(out),
	}
}

func (c *Install) Name() string {
	return "install"
}

func (c *Install) Description() string {
	return fmt.Sprintf("Parse examples into a contract and extract all dependencies. Install each dependency into the Dockpit workspace which is expected to be specified as the PIT_PATH environment variable")
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
	return c.templated(c.Run)
}

func (c *Install) Run(ctx *cli.Context) (*template.Template, interface{}, error) {

	//get pit path
	pp := os.Getenv("PIT_PATH")
	if pp == "" {
		return nil, nil, fmt.Errorf("Couldn't read 'PIT_PATH' environment variable, is it set?")
	}

	//get contract
	contract, err := c.ParseExamples(ctx)
	if err != nil {
		return nil, nil, err
	}

	//retrieve all dependencies
	deps, err := contract.Dependencies()
	if err != nil {
		return nil, nil, err
	}

	//use the manager to install all dependencies into pit path
	m := debs.NewManager(pp)
	installation := map[string]string{}
	for dep, _ := range deps {
		fmt.Fprintf(c.out, "Installing %s:\n", dep)

		err := m.Install(dep)
		if err != nil {
			fmt.Fprintf(c.out, "ERROR \n")
			return nil, nil, err
		}

		//add location installation
		installation[dep], err = m.Locate(dep)
		if err != nil {
			return nil, nil, err
		}

		fmt.Fprintf(c.out, "done!\n\n")
	}

	return template.Must(template.New("install.success").Parse(tmpl_install)), installation, nil
}
