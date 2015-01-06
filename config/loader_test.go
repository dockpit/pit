package config_test

import (
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/bmizerany/assert"

	"github.com/dockpit/pit/config"
)

func TestLoaderLoad(t *testing.T) {
	var c config.C

	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	l := config.NewLoader(wd)

	c, err = l.Load()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, len(c.DependencyConfigs()))
	assert.Equal(t, 2, len(c.ProviderConfigs()))

	//test portconfig fetching
	deppc := c.PortBindingsForDep("github.com/dockpit/ex-store-customers")

	//assert binding
	assert.Equal(t, 1, len(deppc["8000/tcp"]))
	assert.Equal(t, "4321", deppc["8000/tcp"][0].HostPort)

	//state provider spec

	// assert bindings

	// assert mongo state provider
	mongoc := c.StateProviderConfig("mongodb")

	assert.Equal(t, time.Second*10, mongoc.ReadyTimeout())
	assert.Equal(t, []string{"--nojournal"}, mongoc.Cmd())
	assert.Equal(t, true, mongoc.ReadyExp().MatchString("is waiting for conn"))

	mongpc := mongoc.PortBindings()
	assert.Equal(t, 1, len(mongpc["27017/tcp"]))
	assert.Equal(t, "27017", mongpc["27017/tcp"][0].HostPort)

	// assert run command with nil (assume config has a run command for us configured)
	cmd, err := c.RunCommand(nil)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, cmd.Args, []string{"ls", "-la"})

}

func TestCommandOverwrite(t *testing.T) {
	var c config.C

	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	l := config.NewLoader(wd)

	c, err = l.Load()
	if err != nil {
		t.Fatal(err)
	}

	cmd, err := c.RunCommand(exec.Command("ls", "-l"))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, cmd.Args, []string{"ls", "-l"})
}
