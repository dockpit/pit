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
	"github.com/dockpit/lang/manifest"
	"github.com/dockpit/pit/config"
	"github.com/dockpit/pit/reporter"
	"github.com/dockpit/state"
)

var ManifestStatesPath = filepath.Join(".manifest", "states")
var ManifestExamplesPath = filepath.Join(".manifest", "examples")

// CLI Command interface
type C interface {
	Name() string
	Description() string
	Usage() string
	Run(c *cli.Context) error
	Action() func(ctx *cli.Context)
	Flags() []cli.Flag
}

type cmd struct {
	out io.Writer
	reporter.R
}

func newCmd(out io.Writer) *cmd {

	//if no specific output, use stdout
	if out == nil {
		out = os.Stdout
	}

	//redirect default log ouput
	log.SetOutput(out)

	//@todo make the command reporter configurable
	r := reporter.NewTerminal(out)
	return &cmd{out, r}
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
		wd = filepath.Join(wd, ManifestStatesPath)
	} else {
		wd = fmt.Sprintf("[%s]", err.Error())
	}

	return []cli.Flag{
		cli.StringFlag{Name: "states, s", Value: wd, Usage: fmt.Sprintf(" Specify where to look for states.")},
	}
}

func (c *cmd) ConfigFlags() []cli.Flag {

	//get working dir
	wd, err := os.Getwd()
	if err == nil {
		wd = filepath.Join(wd)
	} else {
		wd = fmt.Sprintf("[%s]", err.Error())
	}

	return []cli.Flag{
		cli.StringFlag{Name: "config", Value: wd, Usage: fmt.Sprintf("path to the the directory that contains the configuration file.")},
	}
}

func (c *cmd) LoadConfig(ctx *cli.Context) (config.C, error) {

	//get working dir
	confdir := strings.TrimSpace(ctx.String("config"))
	if confdir == "" {
		return nil, fmt.Errorf("Please specify directory where the configuration file is located (-config)")
	}

	l := config.NewLoader(confdir)
	conf, err := l.Load()
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("Could not find dockpit configuration file (dockpit.json) in '%s'", confdir)
		}
		return nil, err
	}

	return conf, err
}

func (c *cmd) ParseExampleFlags() []cli.Flag {

	//get working dir
	wd, err := os.Getwd()
	if err == nil {
		wd = filepath.Join(wd, ManifestExamplesPath)
	} else {
		wd = fmt.Sprintf("[%s]", err.Error())
	}

	return []cli.Flag{
		cli.StringFlag{Name: "examples, e", Value: wd, Usage: fmt.Sprintf(" Specify where to look for examples.")},
	}
}

func (c *cmd) DockerHostCertArguments(ctx *cli.Context) (string, string, error) {

	//try docker host retrieval
	host := strings.TrimSpace(ctx.String("docker"))
	if host == "" {
		host = os.Getenv("DOCKER_HOST")
		if host == "" {
			return "", "", fmt.Errorf("Could not retrieve DOCKER_HOST, not provided as option and not in env")
		}
	}

	//try docker cert retrieval
	cert := strings.TrimSpace(ctx.String("docker-cert"))
	if cert == "" {
		cert = os.Getenv("DOCKER_CERT_PATH")
		if cert == "" {
			return "", "", fmt.Errorf("Could not retrieve DOCKER_CERT_PATH, not provided as option and not in env")
		}
	}

	return host, cert, nil
}

func (c *cmd) StateManager(ctx *cli.Context) (*state.Manager, error) {

	//path to state context folders
	path := strings.TrimSpace(ctx.String("states"))
	host, cert, err := c.DockerHostCertArguments(ctx)
	if err != nil {
		return nil, err
	}

	conf, err := c.LoadConfig(ctx)
	if err != nil {
		return nil, err
	}

	//create state manager
	m, err := state.NewManager(host, cert, path, conf)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (c *cmd) ExamplesPath(ctx *cli.Context) string {
	return strings.TrimSpace(ctx.String("examples"))
}

func (c *cmd) ParseExamples(ctx *cli.Context) (manifest.M, error) {

	//retrieve path
	path := c.ExamplesPath(ctx)

	//parse the spec
	p := lang.NewParser(path)
	md, err := p.Parse()
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("Failed to open examples in '%s', is this a Dockpit project?", path)
		}

		return nil, fmt.Errorf("Parsing error: %s", err)
	}

	//create manifest from data
	m, err := manifest.NewManifest(md)
	if err != nil {
		return nil, fmt.Errorf("Failed to create manifest from parsed data: %s", err)
	}

	return m, nil
}

func (c *cmd) toAction(fn func(c *cli.Context) error) func(ctx *cli.Context) {
	return func(ctx *cli.Context) {
		//@todo, remove with duplication in constructor
		log.SetOutput(c.out)

		err := fn(ctx)
		if err != nil {

			log.Printf("%s - Command: %s %s", err, ctx.Command.Name, ctx.Args())

			//@todo set exit code to non-zero?
			return
		}

	}
}
