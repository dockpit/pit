package command

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"text/template"

	"github.com/codegangsta/cli"

	"github.com/dockpit/debs"
	"github.com/dockpit/mock/manager"
)

var tmpl_unmock = `Unmocked successful!`

type Unmock struct {
	*cmd
}

func NewUnmock(out io.Writer) *Unmock {
	return &Unmock{
		cmd: newCmd(out),
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
	return c.templated(c.Run)
}

func (c *Unmock) Run(ctx *cli.Context) (*template.Template, interface{}, error) {

	//get pit path
	pp := os.Getenv("PIT_PATH")
	if pp == "" {
		return nil, nil, fmt.Errorf("Couldn't read 'PIT_PATH' environment variable, is it set?")
	}

	//create Unmock manager
	host, cert, err := c.DockerHostCertArguments(ctx)
	if err != nil {
		return nil, nil, err
	}

	//get manifest
	m, err := c.ParseExamples(ctx)
	if err != nil {
		return nil, nil, err
	}

	//retrieve all dependencies
	deps, err := m.Dependencies()
	if err != nil {
		return nil, nil, err
	}

	mm, err := manager.NewManager(host, cert)
	if err != nil {
		return nil, nil, err
	}

	//use the manager to locate dependencies
	dm := debs.NewManager(pp)
	for dep, _ := range deps {

		in, err := dm.Locate(dep)
		if err != nil {
			return nil, nil, err
		}

		//@todo centralize subpath .dockpit/examples
		err = mm.Stop(filepath.Join(in, ManifestExamplesPath))
		if err != nil {
			return nil, nil, err
		}

	}

	return template.Must(template.New("unmock.success").Parse(tmpl_unmock)), nil, nil
}
