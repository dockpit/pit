package command_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/dockpit/pit/command"
	"github.com/dockpit/pit/reporter"
)

func TestInstallIntoTmp(t *testing.T) {

	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	r := reporter.NewTerminal(os.Stdout)
	cmd := command.NewInstall(r)

	tdir, err := ioutil.TempDir("", "dp_tinstallinfotmp")
	if err != nil {
		t.Fatal(err)
	}

	os.Setenv("PIT_PATH", tdir)

	AssertCommandNoError(t, cmd, []string{
		"-examples", filepath.Join(wd, "test_example", command.ManifestExamplesPath),
	}, `(?s)Installing.*github\.com.*success`, r)

	//should be able to read file from installation
	_, err = ioutil.ReadFile(filepath.Join(tdir, "github.com", "dockpit", "pit-token", "main.go"))
	if err != nil {
		t.Fatal(err)
	}

}

func TestInstallSpecificPackage(t *testing.T) {

	r := reporter.NewTerminal(os.Stdout)
	cmd := command.NewInstall(r)

	tdir, err := ioutil.TempDir("", "dp_tinstallinfotmp")
	if err != nil {
		t.Fatal(err)
	}

	os.Setenv("PIT_PATH", tdir)

	AssertCommandNoError(t, cmd, []string{
		"github.com/dockpit/pit-token",
	}, `(?s)Installing.*github\.com.*success`, r)

	//should be able to read file from installation
	_, err = ioutil.ReadFile(filepath.Join(tdir, "github.com", "dockpit", "pit-token", "main.go"))
	if err != nil {
		t.Fatal(err)
	}

}
