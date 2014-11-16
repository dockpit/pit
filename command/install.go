package command

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"text/template"

	"github.com/codegangsta/cli"

	"github.com/dockpit/debs"
	"github.com/dockpit/lang"
	"github.com/dockpit/pit/contract"
)

var tmpl_install = ``

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
	return fmt.Sprintf("...")
}

func (c *Install) Usage() string {
	return "Install a dependency"
}

func (c *Install) Flags() []cli.Flag {
	return []cli.Flag{}
}

func (c *Install) Action() func(ctx *cli.Context) {
	return c.templated(c.Run)
}

func (c *Install) Run(ctx *cli.Context) (*template.Template, interface{}, error) {

	//get working dir
	wd, err := os.Getwd()
	if err != nil {
		return nil, nil, err
	}

	//parse the examples
	p := lang.NewParser(filepath.Join(wd, ".dockpit", "examples"))
	cd, err := p.Parse()
	if err != nil {
		return nil, nil, err
	}

	//create contract from data
	contract, err := contract.NewContract(cd)
	if err != nil {
		return nil, nil, err
	}

	//retrieve all dependencies
	deps, err := contract.Dependencies()
	if err != nil {
		return nil, nil, err
	}

	//use the manager to install all dependencies
	m := debs.NewManager(filepath.Join(wd, ".dockpit"))
	for dep, _ := range deps {
		err := m.Install(dep)
		if err != nil {
			return nil, nil, err
		}
	}

	return template.Must(template.New("install.success").Parse(tmpl_install)), nil, nil
}
