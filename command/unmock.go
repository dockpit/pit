package command

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"

	"github.com/dockpit/debs"
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

	mm, err := manager.NewManager(host, cert)
	if err != nil {
		return err
	}

	//use the manager to locate dependencies
	dm := debs.NewManager(pp)
	for dep, _ := range deps {
		in, err := dm.Locate(dep)
		if err != nil {
			return err
		}

		//@todo centralize subpath .dockpit/examples
		err = mm.Stop(filepath.Join(in, ManifestExamplesPath))
		if err != nil {
			return err
		}
	}

	return nil
}
