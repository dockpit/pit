package command

import (
	"io"
	"text/template"

	"github.com/codegangsta/cli"
)

var tmpl_mock = `mocked!`

type Mock struct {
	*cmd
}

func NewMock(out io.Writer) *Mock {
	return &Mock{
		cmd: newCmd(out),
	}
}

func (c *Mock) Name() string {
	return "mock"
}

func (c *Mock) Description() string {
	return "Mock service specification(s)"
}

func (c *Mock) Usage() string {
	return "Mock service specification(s)."
}

func (c *Mock) Flags() []cli.Flag {
	return []cli.Flag{}
}

func (c *Mock) Action() func(ctx *cli.Context) {
	return c.templated(c.Run)
}

func (c *Mock) Run(ctx *cli.Context) (*template.Template, interface{}, error) {
	return template.Must(template.New("switch.success").Parse(tmpl_mock)), nil, nil
}
