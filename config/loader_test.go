package config_test

import (
	"os"
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
	deppc := c.PortsForDependency("github.com/dockpit/pit-token")

	//assert binding
	assert.Equal(t, 1, len(deppc))
	assert.Equal(t, "4321", deppc[0].Host)

	//state provider spec

	// assert mysql state provider
	mysqlc := c.StateProviderConfig("mysql")

	assert.Equal(t, "default", mysqlc.DefaultState())

	// assert mongo state provider
	mongoc := c.StateProviderConfig("mongodb")

	assert.Equal(t, "no users", mongoc.DefaultState())
	assert.Equal(t, time.Second*10, mongoc.ReadyTimeout())
	assert.Equal(t, []string{"--nojournal"}, mongoc.Cmd())
	assert.Equal(t, true, mongoc.ReadyExp().MatchString("is waiting for conn"))

	mongpc := c.PortsForStateProvider("mongodb")
	assert.Equal(t, 1, len(mongpc))
	assert.Equal(t, "27017", mongpc[0].Host)

	// assert run command with nil (assume config has a run command for us configured)
	rconf := c.RunConfig()

	assert.Equal(t, rconf.Cmd, []string{"ls", "-la"})
	assert.Equal(t, rconf.Dir, "..")
	assert.Equal(t, rconf.ReadyTimeout, time.Second*2)
	assert.Equal(t, true, rconf.ReadyExp.MatchString("2015/01/08 12:15:17.019258 Starting Goji on [::]:8000"))

}
