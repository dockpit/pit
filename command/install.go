package command

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/codegangsta/cli"

	"github.com/dockpit/debs"
	"github.com/dockpit/lang"
	"github.com/dockpit/pit/contract"
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
	return "Install a dependency"
}

func (c *Install) Flags() []cli.Flag {

	//get working dir
	wd, err := os.Getwd()
	if err == nil {
		wd = filepath.Join(wd, ".dockpit", "examples")
	} else {
		wd = fmt.Sprintf("[%s]", err.Error())
	}

	return []cli.Flag{
		cli.StringFlag{Name: "path, p", Value: wd, Usage: fmt.Sprintf(" Specify where to look for examples.")},
	}
}

func (c *Install) Action() func(ctx *cli.Context) {
	return c.templated(c.Run)
}

func (c *Install) Run(ctx *cli.Context) (*template.Template, interface{}, error) {

	//retrieve path
	path := strings.TrimSpace(ctx.String("path"))

	//get pit path
	pp := os.Getenv("PIT_PATH")
	if pp == "" {
		return nil, nil, fmt.Errorf("Couldn't read 'PIT_PATH' environment variable, is it set?")
	}

	//parse the spec
	p := lang.NewParser(path)
	cd, err := p.Parse()
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil, fmt.Errorf("Failed to open .dockpit/examples in '%s', is this a Dockpit project?", path)
		}

		return nil, nil, fmt.Errorf("Parsing error: %s", err)
	}

	//create contract from data
	contract, err := contract.NewContract(cd)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to create contract from parsed data: %s", err)
	}

	//retrieve all dependencies
	deps, err := contract.Dependencies()
	if err != nil {
		return nil, nil, err
	}

	//use the manager to install all dependencies into pit path
	m := debs.NewManager(pp)
	for dep, _ := range deps {
		fmt.Fprintf(c.out, "Installing %s:\n", dep)

		err := m.Install(dep)
		if err != nil {
			fmt.Fprintf(c.out, "ERROR \n")
			return nil, nil, err
		}

		fmt.Fprintf(c.out, "done!\n\n")
	}

	return template.Must(template.New("install.success").Parse(tmpl_install)), nil, nil
}
