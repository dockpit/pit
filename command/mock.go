package command

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/codegangsta/cli"

	"github.com/dockpit/pit/spec"
)

var tmpl_mock = ``

type Mock struct {
	*cmd
	version string

	Docker D
}

func NewMock(out io.Writer, version string) *Mock {
	return &Mock{
		version: version,
		cmd:     newCmd(out),
	}
}

func (c *Mock) Name() string {
	return "mock"
}

func (c *Mock) Description() string {
	return fmt.Sprintf("Mock dependencies of the %s in the current working directory, or the one specifies as the first argument", SpecFilename)
}

func (c *Mock) Usage() string {
	return "mock service dependencies"
}

func (c *Mock) Flags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{Name: "docker, d", Value: "", Usage: fmt.Sprintf("The Docker host location, defaults to reading from DOCKER_HOST environment variable.")},
		cli.StringFlag{Name: "docker-cert, c", Value: "", Usage: fmt.Sprintf("The Docker Certificates Path location, defaults to reading from DOCKER_CERT_PATH environment variable.")},
	}
}

func (c *Mock) Action() func(ctx *cli.Context) {
	return c.templated(c.Run)
}

func (c *Mock) Run(ctx *cli.Context) (*template.Template, interface{}, error) {

	//get working dir
	wd, err := os.Getwd()
	if err != nil {
		return nil, nil, err
	}

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

	//ask spec to map its dependencies
	deps, err := spec.Dependencies()
	if err != nil {
		return nil, nil, err
	}

	//try docker host retrieval
	host := strings.TrimSpace(ctx.String("docker"))
	if host == "" {
		host = os.Getenv("DOCKER_HOST")
		if host == "" {
			return nil, nil, fmt.Errorf("Could not retrieve DOCKER_HOST, not provided as option and not in env")
		}
	}

	//try docker cert retrieval
	cert := strings.TrimSpace(ctx.String("docker-cert"))
	if cert == "" {
		cert = os.Getenv("DOCKER_CERT_PATH")
		if cert == "" {
			return nil, nil, fmt.Errorf("Could not retrieve DOCKER_CERT_PATH, not provided as option and not in env")
		}
	}

	//create docker client
	docker := c.Docker
	if docker == nil {

		//create docker instance
		docker, err = NewDocker(host, cert, c.version)
		if err != nil {
			return nil, nil, err
		}
	}

	//remove all running dockpit containers
	fmt.Fprint(c.out, "Removing old dependency mock containers...")
	err = docker.RemoveAll()
	if err != nil {
		return nil, nil, err
	}

	//launch dependency containers
	fmt.Fprint(c.out, "done!\nStarting new dependency mock containers...")
	err = docker.Start(deps)
	if err != nil {
		return nil, nil, err
	}

	fmt.Fprint(c.out, "done!\n")
	return template.Must(template.New("mock.success").Parse(tmpl_mock)), nil, nil
}
