package command_test

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/dockpit/pit/command"
)

func TestBuildingOfStates(t *testing.T) {

	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	out := bytes.NewBuffer(nil)
	cmd := command.NewBuild(out)

	AssertCommand(t, cmd, []string{
		"-config", filepath.Join(wd, "test_example"),
		"-examples", filepath.Join(wd, "test_example", command.ManifestExamplesPath),
		"-states", filepath.Join(wd, "test_example", command.ManifestStatesPath),
	}, `(?s)pitstate_mongo_52f3cc.*pitstate_mongo_1815bee1bdd121d.*successful`, out)
}
