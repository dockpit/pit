package command

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"text/template"
	"time"

	"github.com/codegangsta/cli"
	"github.com/fsouza/go-dockerclient"

	"github.com/dockpit/dirtar"
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
	ids := map[string]*bytes.Buffer{}
	for dep, dir := range installation {

		iname := fmt.Sprintf("%s:%s", "dockpit/pit", "latest")
		fmt.Fprintf(c.out, "Mocking %s using image %s...\n", dep, iname)

		//containers are created from the dockpit image
		copts := docker.CreateContainerOptions{
			// Name: fmt.Sprintf("pitmock_%s", filepath.Base(dep)),
			Config: &docker.Config{
				Image: iname,
				Cmd:   []string{"serve"},
			},
		}

		//contairs are started with published public ports
		sopts := &docker.HostConfig{
			PublishAllPorts: true,
		}

		fmt.Fprintf(c.out, "\tcreating container...")

		//create container
		container, err := client.CreateContainer(copts)
		if err != nil {
			return nil, nil, err
		}

		fmt.Fprintf(c.out, "done!\n\tstarting container %s...", container.ID)

		//start container
		err = client.StartContainer(container.ID, sopts)
		if err != nil {
			return nil, nil, err
		}

		//wait for container to start
		//@todo very dirty business here
		<-time.After(time.Millisecond * 200)

		//tar examples
		data := bytes.NewBuffer(nil)
		tpath := filepath.Join(dir, ".dockpit", "examples")

		fmt.Fprintf(c.out, "done!\n\ttarring examples @ %s...", tpath)

		err = dirtar.Tar(tpath, data)
		if err != nil {
			return nil, nil, err
		}

		ids[container.ID] = data
		if err != nil {
			if os.IsNotExist(err) {
				//@todo, the installed dependency doesnt have any examples, what shall we do?
				return nil, nil, fmt.Errorf("The dependency %s did not provide any examples (%s)", dep, dir)
			}

			return nil, nil, err
		}

		fmt.Fprintf(c.out, "done!\n")
	}

	//prepare url to post context to
	host, _, err := c.DockerHostCertArguments(ctx)
	if err != nil {
		return nil, nil, err
	}

	loc, err := url.Parse(host)
	if err != nil {
		return nil, nil, err
	}

	loc.Scheme = "http"
	loc.Path = "/_examples"

	//get status
	cs, err := client.ListContainers(docker.ListContainersOptions{})
	if err != nil {
		return nil, nil, err
	}

	//upload examples
	for _, container := range cs {

		//we only care about containers we just launched
		var tar *bytes.Buffer
		var ok bool
		if tar, ok = ids[container.ID]; !ok {
			continue
		}

		fmt.Fprintf(c.out, "Reload Examples of %s...\n", container.ID)

		//get the external port for 8000
		//@todo make 8000 (private port) configurable
		//@todo make :2376 (docker host port) configurable
		for _, pconfig := range container.Ports {
			if pconfig.PrivatePort == 8000 {
				loc.Host = strings.Replace(loc.Host, ":2376", fmt.Sprintf(":%d", pconfig.PublicPort), 1)
			}
		}

		fmt.Fprintf(c.out, "\tUploading %s...", container.ID)

		//send request to upload
		resp, err := http.Post(loc.String(), "application/x-tar", tar)
		if err != nil {
			return nil, nil, err
		}

		//were the examples uploaded correctly?
		if resp.StatusCode != 201 {
			return nil, nil, fmt.Errorf("Failed to upload examples to container '%s' %s", container.ID, container.Names)
		}

		fmt.Fprintf(c.out, "done!\n\tSending HUP signal...")

		//wait for container to settle?
		//@todo very dirty business here
		<-time.After(time.Millisecond * 200)

		//send HUP signal to 'root' process to reload examples
		//@from http://stackoverflow.com/questions/25687131/how-to-send-signal-to-program-run-in-a-docker-container
		err = client.KillContainer(docker.KillContainerOptions{
			ID:     container.ID,
			Signal: docker.Signal(syscall.SIGHUP),
		})

		if err != nil {
			return nil, nil, err
		}

		//report progress
		fmt.Fprintf(c.out, "done!\n\tMock %s complete - %s://%s\n\n", container.Names, loc.Scheme, loc.Host)
	}

	return template.Must(template.New("mock.success").Parse(tmpl_mock)), nil, nil
}
