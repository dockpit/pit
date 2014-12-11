package command_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/dockpit/pit/command"
)

func TestInstallIntoTmp(t *testing.T) {

	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	out := bytes.NewBuffer(nil)
	cmd := command.NewInstall(out)

	tdir, err := ioutil.TempDir("", "dp_tinstallinfotmp")
	if err != nil {
		t.Fatal(err)
	}

	os.Setenv("PIT_PATH", tdir)

	AssertCommand(t, cmd, []string{"-examples", filepath.Join(wd, "..", command.ManifestExamplesPath)}, `(?s)Installing github\.com.*done\!.*successful`, out)

	//should be able to read file from installation
	_, err = ioutil.ReadFile(filepath.Join(tdir, "deps", "github.com", "dockpit", "ex-store-customers", "main.go"))
	if err != nil {
		t.Fatal(err)
	}

}
