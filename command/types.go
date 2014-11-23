package command

import (
	"fmt"
	"io"
	"log"
	"os"
	"text/template"

	"github.com/codegangsta/cli"

	"github.com/dockpit/pit/spec"
)

var SpecFilename = "dockpit.json"

// Docker client abstraction interface
type D interface {
	RemoveAll() error
	Start(*spec.Dependencies) error
}

// CLI Command interface
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

	//if no specific output, use stdout
	if out == nil {
		out = os.Stdout
	}

	//redirect default log ouput
	log.SetOutput(out)

	return &cmd{out}
}

func (c *cmd) Run(ctx *cli.Context) (*template.Template, interface{}, error) {
	return nil, nil, fmt.Errorf("Command '%s' is not yet implemented", ctx.Command.Name)
}

func (c *cmd) templated(fn func(c *cli.Context) (*template.Template, interface{}, error)) func(ctx *cli.Context) {
	return func(ctx *cli.Context) {
		t, data, err := fn(ctx)
		if err != nil {
			log.Println(err, ", Command: '", ctx.Command.Name, "' Args: ", ctx.Args())

			//@todo set exit code to non-zero?
			return
		}

		err = t.Execute(c.out, data)
		if err != nil {
			log.Fatal("Template Error: ", err)
		}
	}
}
