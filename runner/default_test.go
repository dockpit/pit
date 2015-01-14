package runner_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dockpit/pit/config"
	"github.com/dockpit/pit/reporter"
	"github.com/dockpit/pit/runner"
)

func TestCreation(t *testing.T) {
	reporter := reporter.NewTerminal(os.Stdout)
	r, err := runner.Create("default", reporter)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "default", r.Name())
}

func TestCommandArgsTemplating(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	reporter := reporter.NewTerminal(os.Stdout)
	r := runner.NewDefault(reporter)
	l := config.NewLoader(filepath.Join(wd, "..", "command", "test_example"))
	conf, err := l.Load()
	if err != nil {
		t.Fatal(err)
	}

	//set docker hostname for test
	data := conf.Data()
	data.DockerHostname = "localhost"

	args, err := r.TemplatedCommandArgs(conf, "{{ .DockerHostname }}", "{{ .DockerHostname }}")
	if err != nil {
		t.Fatal(err)
	}

	//hostname is not configured in this case
	assert.Equal(t, "localhost", args[0])
	assert.Equal(t, "localhost", args[1])
}
