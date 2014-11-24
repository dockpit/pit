package command

import (
	"flag"
	"fmt"
	"io"
	"strings"
	"text/template"

	"github.com/codegangsta/cli"
	"github.com/zenazn/goji"
)

var tmpl_serve = `Served successful!`

type Serve struct {
	*cmd
}

func NewServe(out io.Writer) *Serve {
	return &Serve{
		cmd: newCmd(out),
	}
}

func (c *Serve) Name() string {
	return "serve"
}

func (c *Serve) Description() string {
	return fmt.Sprintf("...")
}

func (c *Serve) Usage() string {
	return "serve a folder of examples as being a web service"
}

func (c *Serve) Flags() []cli.Flag {
	fs := []cli.Flag{}
	fs = append(fs, c.ParseExampleFlags()...)
	fs = append(fs, cli.StringFlag{Name: "bind, b", Value: ":8000", Usage: fmt.Sprintf("Specify on which address the server will listen.")})

	return fs
}

func (c *Serve) Action() func(ctx *cli.Context) {
	return c.templated(c.Run)
}

func (c *Serve) Run(ctx *cli.Context) (*template.Template, interface{}, error) {

	//get contract
	contract, err := c.ParseExamples(ctx)
	if err != nil {
		return nil, nil, err
	}

	//transfer bind from ctx
	flag.Set("bind", strings.TrimSpace(ctx.String("bind")))

	//@todo create mock
	fmt.Println(contract)

	//@todo turn contract into mux

	//start goij server, reuse logging stack
	// goji.DefaultMux.Handle("/*", mock.Mux())
	goji.Serve()

	return template.Must(template.New("serve.success").Parse(tmpl_serve)), nil, nil
}
