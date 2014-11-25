package command

import (
	"fmt"
	"io"
	"text/template"

	"github.com/codegangsta/cli"
	"github.com/fsouza/go-dockerclient"
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

	//get docker client
	client, err := c.DockerClient(ctx)
	if err != nil {
		return nil, nil, err
	}

	//@todo remove all old containers

	//loop over installation and create pit containers
	for _, _ = range installation {

		//@todo inject correct contract data

		//containers are created from the dockpit image
		copts := docker.CreateContainerOptions{
			Config: &docker.Config{
				Image: fmt.Sprintf("%s:%s", "dockpit/pit", "latest"),
				Cmd:   []string{"serve"},
			},
		}

		//contairs are started with published public ports
		sopts := &docker.HostConfig{
			PublishAllPorts: true,
		}

		//create container
		c, err := client.CreateContainer(copts)
		if err != nil {
			return nil, nil, err
		}

		//start container
		err = client.StartContainer(c.ID, sopts)
		if err != nil {
			return nil, nil, err
		}

	}

	return template.Must(template.New("Mock.success").Parse(tmpl_mock)), nil, nil
}
