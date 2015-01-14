package command

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"

	"github.com/dockpit/debs"
	"github.com/dockpit/mock/manager"
)

type Mock struct {
	*cmd

	install *Install //the install command
}

func NewMock(out io.Writer, install *Install) *Mock {
	return &Mock{
		cmd:     newCmd(out),
		install: install,
	}
}

func (c *Mock) Name() string {
	return "mock"
}

func (c *Mock) Description() string {
	return fmt.Sprintf("...")
}

func (c *Mock) Usage() string {
	return "Mock all dependencies of a given set of examples"
}

func (c *Mock) Flags() []cli.Flag {
	fs := []cli.Flag{}
	fs = append(fs, c.install.Flags()...)
	fs = append(fs, c.ConfigFlags()...)
	fs = append(fs, c.DockerFlags()...)

	return fs
}

func (c *Mock) Action() func(ctx *cli.Context) {
	return c.toAction(c.Run)
}

func (c *Mock) Run(ctx *cli.Context) error {

	//run install command
	//@todo make optional?
	err := c.install.Run(ctx)
	if err != nil {
		return err
	}

	//get pit path
	pp := os.Getenv("PIT_PATH")
	if pp == "" {
		return fmt.Errorf("Couldn't read 'PIT_PATH' environment variable, is it set?")
	}

	//get manifest
	m, err := c.ParseExamples(ctx)
	if err != nil {
		return err
	}

	//retrieve all dependencies
	deps, err := m.Dependencies()
	if err != nil {
		return err
	}

	//create mock manager
	host, cert, err := c.DockerHostCertArguments(ctx)
	if err != nil {
		return err
	}

	//load configuration
	conf, err := c.LoadConfig(ctx)
	if err != nil {
		return err
	}

	mm, err := manager.NewManager(host, cert)
	if err != nil {
		return err
	}

	//start the mock of each installation
	dm := debs.NewManager(pp)
	for dep, _ := range deps {
		in, err := dm.Locate(dep)
		if err != nil {
			return err
		}

		fmt.Fprintf(c.out, "Mocking %s...", dep)

		//get first prot binding from configuration
		ports := conf.PortsForDependency(dep)

		// @todo centralize this?
		mc, err := mm.Start(filepath.Join(in, ManifestExamplesPath), ports[0].Host)
		if err != nil {
			return err
		}

		fmt.Fprintf(c.out, " done! (%s)\n", mc.Endpoint)

	}

	return nil
}
