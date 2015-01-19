package command

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"

	"github.com/dockpit/deps"
	"github.com/dockpit/mock/manager"
	"github.com/dockpit/pit/reporter"
)

type Mock struct {
	*cmd
	install *Install
}

func NewMock(r reporter.R, install *Install) *Mock {
	return &Mock{
		cmd:     newCmd(r),
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
	c.Enter(MockPart, MockPart.StartingMocks)
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	//get pit path
	pp := os.Getenv("PIT_PATH")
	if pp == "" {
		return fmt.Errorf("Couldn't read 'PIT_PATH' environment variable, is it set?")
	}

	//run install command
	//@todo make optional?
	err = c.install.Run(ctx)
	if err != nil {
		return err
	}

	mrel, err := filepath.Rel(wd, ctx.String("examples"))
	if err != nil {
		return err
	}

	//get manifest
	c.Report(ManifestPart.ParsingExamples, mrel)
	m, err := c.ParseExamples(ctx)
	if err != nil {
		return err
	}

	//retrieve all dependencies
	depl, err := m.Dependencies()
	if err != nil {
		return err
	}

	//create mock manager
	host, cert, err := c.DockerHostCertArguments(ctx)
	if err != nil {
		return err
	}

	confrel, err := filepath.Rel(wd, ctx.String("config"))
	if err != nil {
		return err
	}

	//load configuration
	c.Report(ConfigPart.LoadingConfig, confrel)
	conf, err := c.LoadConfig(ctx)
	if err != nil {
		return err
	}

	mm, err := manager.NewManager(host, cert)
	if err != nil {
		return err
	}

	//start the mock of each installation
	c.Report(MockPart.MockingFrom, pp)
	dm := deps.NewManager(pp)
	for dep, _ := range depl {
		c.Enter(DepPart, DepPart.MockingDep, dep)

		in, err := dm.Locate(dep)
		if err != nil {
			return err
		}

		//get first prot binding from configuration
		ports := conf.PortsForDependency(dep)

		//start dependency
		mc, err := mm.Start(filepath.Join(in, ManifestExamplesPath), ports[0].Host)
		if err != nil {
			return err
		}

		c.Success(DepPart.MockedDep, mc.Endpoint)
		c.Exit()
	}

	return nil
}
