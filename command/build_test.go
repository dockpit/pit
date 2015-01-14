package command_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/dockpit/pit/command"
	"github.com/dockpit/pit/reporter"
)

func TestBuildingOfStates(t *testing.T) {

	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	r := reporter.NewTerminal(os.Stdout)
	cmd := command.NewBuild(r)

	AssertCommandNoError(t, cmd, []string{
		"-config", filepath.Join(wd, "test_example"),
		"-examples", filepath.Join(wd, "test_example", command.ManifestExamplesPath),
		"-states", filepath.Join(wd, "test_example", command.ManifestStatesPath),
	}, `(?s)pitstate_mongo_52f3cc.*pitstate_mongo_1815bee1bdd121d.*`, r)
}
