package config_test

import (
	"os"
	"testing"

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
	mysqlpc := c.PortBindingsForState("mysql")

	// assert bindings
	assert.Equal(t, 1, len(mysqlpc["3306/tcp"]))
	assert.Equal(t, "3306", mysqlpc["3306/tcp"][0].HostPort)

}
