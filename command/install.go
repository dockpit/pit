package command

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"text/template"

	"github.com/codegangsta/cli"

	"github.com/dockpit/pit/command/tool"
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

	//define package by importlink
	p := tool.NewPackage("github.com/dockpit/pit")

	tool.DownloadPackage(p, filepath.Join(wd, ".dockpit"))

	return template.Must(template.New("install.success").Parse(tmpl_install)), nil, nil
}
