package command

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/codegangsta/cli"

	"github.com/dockpit/lang"
	"github.com/dockpit/pit/contract"
	"github.com/dockpit/pit/spec"
	"github.com/dockpit/state"
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

func (c *cmd) DockerFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{Name: "docker, d", Value: "", Usage: fmt.Sprintf("The Docker host location, defaults to reading from DOCKER_HOST environment variable.")},
		cli.StringFlag{Name: "docker-cert, c", Value: "", Usage: fmt.Sprintf("The Docker Certificates Path location, defaults to reading from DOCKER_CERT_PATH environment variable.")},
	}
}

func (c *cmd) BuildStatesFlags() []cli.Flag {

	//get working dir
	wd, err := os.Getwd()
	if err == nil {
		wd = filepath.Join(wd, ".dockpit", "states")
	} else {
		wd = fmt.Sprintf("[%s]", err.Error())
	}

	return []cli.Flag{
		cli.StringFlag{Name: "states, s", Value: wd, Usage: fmt.Sprintf(" Specify where to look for states.")},
	}
}

func (c *cmd) ParseExampleFlags() []cli.Flag {

	//get working dir
	wd, err := os.Getwd()
	if err == nil {
		wd = filepath.Join(wd, ".dockpit", "examples")
	} else {
		wd = fmt.Sprintf("[%s]", err.Error())
	}

	return []cli.Flag{
		cli.StringFlag{Name: "examples, e", Value: wd, Usage: fmt.Sprintf(" Specify where to look for examples.")},
	}
}

func (c *cmd) StateManager(ctx *cli.Context) (*state.Manager, error) {

	//path to state context folders
	path := strings.TrimSpace(ctx.String("states"))

	//try docker host retrieval
	host := strings.TrimSpace(ctx.String("docker"))
	if host == "" {
		host = os.Getenv("DOCKER_HOST")
		if host == "" {
			return nil, fmt.Errorf("Could not retrieve DOCKER_HOST, not provided as option and not in env")
		}
	}

	//try docker cert retrieval
	cert := strings.TrimSpace(ctx.String("docker-cert"))
	if cert == "" {
		cert = os.Getenv("DOCKER_CERT_PATH")
		if cert == "" {
			return nil, fmt.Errorf("Could not retrieve DOCKER_CERT_PATH, not provided as option and not in env")
		}
	}

	//create state manager
	m, err := state.NewManager(host, cert, path)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (c *cmd) ParseExamples(ctx *cli.Context) (contract.C, error) {

	//retrieve path
	path := strings.TrimSpace(ctx.String("examples"))

	//parse the spec
	p := lang.NewParser(path)
	cd, err := p.Parse()
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("Failed to open .dockpit/examples in '%s', is this a Dockpit project?", path)
		}

		return nil, fmt.Errorf("Parsing error: %s", err)
	}

	//create contract from data
	contract, err := contract.NewContract(cd)
	if err != nil {
		return nil, fmt.Errorf("Failed to create contract from parsed data: %s", err)
	}

	return contract, nil
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
