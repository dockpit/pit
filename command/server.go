package command

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/codegangsta/cli"
	"github.com/zenazn/goji"

	"github.com/dockpit/pit/spec"
)

var tmpl_server = ``

type Server struct {
	*cmd
}

func NewServer(out io.Writer) *Server {
	return &Server{
		cmd: newCmd(out),
	}
}

func (c *Server) Name() string {
	return "server"
}

func (c *Server) Description() string {
	return fmt.Sprintf("Looks at the given specification to create a mock server that adhers to the given spec. If no spec is provided, the %s file in the current working directory is used.", SpecFilename)
}

func (c *Server) Usage() string {
	return "launches a server that mocks a service specification"
}

func (c *Server) Flags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{Name: "bind, b", Value: ":8000", Usage: fmt.Sprintf("Specify on which address the server will listen.")},
	}
}

func (c *Server) Action() func(ctx *cli.Context) {
	return c.templated(c.Run)
}

func (c *Server) Run(ctx *cli.Context) (*template.Template, interface{}, error) {

	//get working dir
	wd, err := os.Getwd()
	if err != nil {
		return nil, nil, err
	}

	//transfer bind from ctx
	flag.Set("bind", strings.TrimSpace(ctx.String("bind")))

	//first may provide path to spec src
	src := strings.TrimSpace(ctx.Args().First())
	if src == "" {

		//defaults to file in working dir
		src = path.Join(wd, SpecFilename)
	}

	//create services
	loader := spec.NewLoader()
	factory := spec.NewFactory(loader)

	//create spec
	spec, err := factory.Create(src)
	if err != nil {
		return nil, nil, err
	}

	//create mock from spec
	mock, err := spec.Mock()
	if err != nil {
		return nil, nil, err
	}

	//start goij server, reuse logging stack
	goji.DefaultMux.Handle("/*", mock.Mux())
	goji.Serve()

	return template.Must(template.New("server.success").Parse(tmpl_server)), nil, nil
}
