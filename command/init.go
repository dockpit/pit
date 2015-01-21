package command

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"

	"github.com/dockpit/pit/config"
	"github.com/dockpit/pit/reporter"
)

type Init struct {
	*cmd
}

func NewInit(r reporter.R) *Init {
	return &Init{
		cmd: newCmd(r),
	}
}

func (c *Init) Name() string {
	return "Init"
}

func (c *Init) Description() string {
	return fmt.Sprintf("Creates a minimal set of files and directories to jumpstart a new Dockpit project in the currend directory.")
}

func (c *Init) Usage() string {
	return "Bootstrap a new Dockpit project"
}

func (c *Init) Flags() []cli.Flag {
	fs := []cli.Flag{}

	return fs
}

func (c *Init) Action() func(ctx *cli.Context) {
	return c.toAction(c.Run)
}

func (c *Init) Run(ctx *cli.Context) error {
	c.Enter(InitPart, InitPart.StartingInit)
	var err error

	dir := ctx.Args().First()
	if dir == "" {
		dir, err = os.Getwd()
		if err != nil {
			return err
		}
	}

	c.Report(InitPart.InitittingInto, dir)
	l := config.NewLoader(dir)

	//create dockpit.json
	_, err = os.Open(l.Path())
	if os.IsNotExist(err) {
		d, err := l.InitData()
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(l.Path(), d, 0644)
		if err != nil {
			return err
		}

		c.Success(InitPart.InitializedConfigFile, l.Path())
	} else {
		c.Warning(InitPart.ConfigFileAlreadyExists, l.Path())
	}

	//create .manifest/examples file
	err = os.MkdirAll(filepath.Join(dir, ManifestExamplesPath), 0755)
	if err != nil {
		return err
	}

	//create .manifest/states file
	err = os.MkdirAll(filepath.Join(dir, ManifestStatesPath), 0755)
	if err != nil {
		return err
	}

	c.Success(InitPart.InitializedManifestDir, ".manifest")
	c.Exit()
	return nil
}
