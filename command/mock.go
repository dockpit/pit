package command

import (
	"fmt"
	"io"
	"path/filepath"
	"text/template"

	"github.com/codegangsta/cli"

	"github.com/dockpit/mock/manager"
)

var tmpl_mock = `Mocked successful!`

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
	fs = append(fs, c.DockerFlags()...)

	return fs
}

func (c *Mock) Action() func(ctx *cli.Context) {
	return c.templated(c.Run)
}

func (c *Mock) Run(ctx *cli.Context) (*template.Template, interface{}, error) {

	//run install command
	//@todo make optional?
	_, res, err := c.install.Run(ctx)
	if err != nil {
		return nil, nil, err
	}

	//assert installation to map
	installation, ok := res.(map[string]string)
	if !ok {
		return nil, nil, fmt.Errorf("Unexpected response from install command: %s", res)
	}

	//create mock manager
	host, cert, err := c.DockerHostCertArguments(ctx)
	if err != nil {
		return nil, nil, err
	}

	m, err := manager.NewManager(host, cert)
	if err != nil {
		return nil, nil, err
	}

	//start the mock of each installation
	for dep, in := range installation {

		fmt.Fprintf(c.out, "Mocking %s...", dep)

		//@todo centralize this?
		mc, err := m.Start(filepath.Join(in, ".dockpit", "examples"))
		if err != nil {
			return nil, nil, err
		}

		fmt.Fprintf(c.out, " done! (%s)\n", mc.Endpoint)

	}

	return template.Must(template.New("mock.success").Parse(tmpl_mock)), nil, nil
}
