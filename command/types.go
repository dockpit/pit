package command

import (
	"fmt"
	"io"
	"log"
	"os"
	"text/template"

	"github.com/codegangsta/cli"

	"github.com/advanderveer/micros/loader"
)

type C interface {
	Name() string
	Description() string
	Usage() string
	Run(c *cli.Context) (*template.Template, interface{}, error)
	Action() func(ctx *cli.Context)
	Flags() []cli.Flag
}

type cmd struct {
	out io.Writer
}

func newCmd(out io.Writer) *cmd {
	if out == nil {
		out = os.Stdout
	}

	return &cmd{out}
}

func (c *cmd) Run(ctx *cli.Context) (*template.Template, interface{}, error) {
	return nil, nil, fmt.Errorf("Command '%s' is not yet implemented", ctx.Command.Name)
}

func (c *cmd) loadSpec(path string) (*loader.Spec, error) {
	sfile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer sfile.Close()

	//load service spec
	bl := loader.NewBasic()
	spec, err := bl.Load(sfile)
	if err != nil {
		return nil, err
	}

	return spec, nil
}

func (c *cmd) templated(fn func(c *cli.Context) (*template.Template, interface{}, error)) func(ctx *cli.Context) {
	return func(ctx *cli.Context) {
		t, data, err := fn(ctx)
		if err != nil {
			log.Fatal(err, ", Command: '", ctx.Command.Name, "' Args: ", ctx.Args())
		}

		err = t.Execute(c.out, data)
		if err != nil {
			log.Fatal("Template Error: ", err)
		}
	}
}
