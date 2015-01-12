package command_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	// "github.com/bmizerany/assert"

	"github.com/dockpit/pit/command"
)

func TestMock(t *testing.T) {

	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	expath := filepath.Join(wd, "test_example", command.ManifestExamplesPath)
	out := bytes.NewBuffer(nil)
	cmd := command.NewMock(out, command.NewInstall(out))

	tdir, err := ioutil.TempDir("", "dp_tinstallinfotmp")
	if err != nil {
		t.Fatal(err)
	}

	os.Setenv("PIT_PATH", tdir)

	AssertCommand(t, cmd, []string{
		"-config", filepath.Join(wd, "test_example"),
		"-examples", expath,
	}, `(?s)Mocking.*done!.*http`, out)

	//should have called install
	_, err = ioutil.ReadFile(filepath.Join(tdir, "deps", "github.com", "dockpit", "pit-token", "main.go"))
	if err != nil {
		t.Fatal(err)
	}

	//unmock afterwards
	cmd2 := command.NewUnmock(out)

	AssertCommand(t, cmd2, []string{"-examples", expath}, `(?s)Unmocked`, out)
}
