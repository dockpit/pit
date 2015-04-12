package command

import (
	"fmt"
	"strconv"

	"github.com/codegangsta/cli"
	"github.com/hashicorp/errwrap"
	"github.com/samalba/dockerclient"

	"github.com/dockpit/pit/client"
	"github.com/dockpit/pit/model"
)

type Locate struct {
	*command
}

func NewLocate() *Locate {
	return &Locate{newCommand()}
}

func (c *Locate) Name() string {
	return "locate"
}

func (c *Locate) Description() string {
	return fmt.Sprintf("Allows you to quickly grab the location of a running dependeny, most often used for generating arguments or configurations for the service being developed. The command takes the dependeny name and private port you're locating as its first and second argument. Example: executing `pit locate etcd 4001` might return '192.168.99.100:49226'")
}

func (c *Locate) Usage() string {
	return "Print the address and public port of running dependency"
}

func (c *Locate) Flags() []cli.Flag {
	return []cli.Flag{
		DBFlag,
		DockerHostFlag,
		DockerCertFlag,
	}
}

func (c *Locate) Action() func(ctx *cli.Context) {
	return c.command.Action(c.Run)
}

func (c *Locate) Run(ctx *cli.Context) error {
	m, err := model.NewModel(ctx.String("db"))
	if err != nil {
		return errwrap.Wrapf("Failed open model: {{err}}", err)
	}

	docker, err := client.NewDocker(m, ctx.String("docker-host"), ctx.String("docker-cert-path"))
	if err != nil {
		return errwrap.Wrapf("Failed to connect to Docker host: {{err}}", err)
	}

	isos, err := m.GetAllIsolations()
	if err != nil {
		return errwrap.Wrapf("Failed to retrieve all isolations: {{err}}", err)
	}

	deps, err := m.GetAllDeps()
	if err != nil {
		return errwrap.Wrapf("Failed to retrieve all dependencies: {{err}}", err)
	}

	running, err := docker.Running(isos, deps)
	if err != nil {
		return errwrap.Wrapf("Failed to retrieve inlocateion on running Dockpit containers: {{err}}", err)
	}

	//for now only one running isolation is supported
	if len(running) > 1 {
		return fmt.Errorf("Unexpected: More then one isolation detected to be running, locating is not supported - maybe you need to run the 'clean' command?")
	} else if len(running) < 1 {
		return fmt.Errorf("Unexpected: No isolation seems to be running, did you start one from in terminal UI?")
	}

	//below only makes sense if there is one
	//isolation running
	var riso *model.Isolation
	var rdeps map[*model.Dep]dockerclient.Container

	for iso, deps := range running {
		riso = iso
		rdeps = deps
	}

	dname := ctx.Args().First()
	if dname == "" {
		return fmt.Errorf("Please provide the dependency name and container port. Example: `pit locate postgres 2222`")
	} else {
		cportStr := ctx.Args().Get(1)
		if cportStr == "" {
			return fmt.Errorf("Please provide the container port for locating dependency '%s' as a second argument. Example: `pit locate %s 2222`", dname, dname)
		}

		cport, err := strconv.Atoi(cportStr)
		if err != nil {
			return errwrap.Wrapf(fmt.Sprintf("Failed to convert port argument '%s' to a number: {{err}}", cportStr), err)
		}

		var foundC dockerclient.Container
		var location string
		for rdep, c := range rdeps {
			if rdep.Name == dname {
				foundC = c
				for _, p := range c.Ports {
					if p.PrivatePort == cport {
						location = fmt.Sprintf("%s:%d", docker.HostIP, p.PublicPort)
					}
				}
			}
		}

		if foundC.Id == "" {
			return fmt.Errorf("Could not find a running dependency named '%s' in the current isolation: '%s'", dname, riso.Name)
		}

		if location == "" {
			return fmt.Errorf("Could not find an exposed port for container port '%d', exposed ports: %v", cport, foundC.Ports)
		}

		fmt.Println(location)
		return nil
	}

	return nil
}
