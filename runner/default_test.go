package runner_test

import (
	"os"
	"testing"

	"github.com/bmizerany/assert"

	"github.com/dockpit/pit/runner"
)

func TestCreation(t *testing.T) {
	r, err := runner.Create("default", os.Stdout)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "default", r.Name())
}
