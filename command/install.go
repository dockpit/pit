package command

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"text/template"

	"github.com/codegangsta/cli"

	"github.com/dockpit/debs"
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

	//parse spec for dependencies

	//install all dependencies

	m := debs.NewManager(filepath.Join(wd, ".dockpit"))

	_ = m

	//download or update if dependency is already installed
	// tool.DownloadPackage(p, filepath.Join(wd, ".dockpit"))

	return template.Must(template.New("install.success").Parse(tmpl_install)), nil, nil
}
