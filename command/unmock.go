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

type Unmock struct {
	*cmd
}

func NewUnmock(r reporter.R) *Unmock {
	return &Unmock{
		cmd: newCmd(r),
	}
}

func (c *Unmock) Name() string {
	return "unmock"
}

func (c *Unmock) Description() string {
	return fmt.Sprintf("...")
}

func (c *Unmock) Usage() string {
	return "remove mocks of a given set of examples"
}

func (c *Unmock) Flags() []cli.Flag {
	fs := []cli.Flag{}
	fs = append(fs, c.ParseExampleFlags()...)
	fs = append(fs, c.DockerFlags()...)

	return fs
}

func (c *Unmock) Action() func(ctx *cli.Context) {
	return c.toAction(c.Run)
}

func (c *Unmock) Run(ctx *cli.Context) error {
	c.Enter(MockPart, MockPart.StoppingMocks)
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	//get pit path
	pp := os.Getenv("PIT_PATH")
	if pp == "" {
		return fmt.Errorf("Couldn't read 'PIT_PATH' environment variable, is it set?")
	}

	//create Unmock manager
	host, cert, err := c.DockerHostCertArguments(ctx)
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

	mm, err := manager.NewManager(host, cert)
	if err != nil {
		return err
	}

	//start the mock of each installation
	c.Report(MockPart.UnmockingFrom, pp)
	dm := deps.NewManager(pp)
	for dep, _ := range depl {
		c.Enter(DepPart, DepPart.UnmockingDep, dep)

		in, err := dm.Locate(dep)
		if err != nil {
			return err
		}

		err = mm.Stop(filepath.Join(in, ManifestExamplesPath))
		if err != nil {
			return err
		}

		c.Success(DepPart.UnmockedDep)
		c.Exit()
	}

	c.Exit()
	return nil
}
