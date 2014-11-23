package command

import (
	"fmt"
	"io"

	"text/template"

	"github.com/codegangsta/cli"
)

var tmpl_build = `State building successful!`

type Build struct {
	*cmd
}

func NewBuild(out io.Writer) *Build {
	return &Build{
		cmd: newCmd(out),
	}
}

func (c *Build) Name() string {
	return "build"
}

func (c *Build) Description() string {
	return fmt.Sprintf("...")
}

func (c *Build) Usage() string {
	return "Build the states using Docker"
}

func (c *Build) Flags() []cli.Flag {
	fs := []cli.Flag{}
	fs = append(fs, c.ParseExampleFlags()...)
	fs = append(fs, c.BuildStatesFlags()...)

	return fs
}

func (c *Build) Action() func(ctx *cli.Context) {
	return c.templated(c.Run)
}

func (c *Build) Run(ctx *cli.Context) (*template.Template, interface{}, error) {

	//get contract
	contract, err := c.ParseExamples(ctx)
	if err != nil {
		return nil, nil, err
	}

	_ = contract

	return template.Must(template.New("build.success").Parse(tmpl_build)), nil, nil
}
