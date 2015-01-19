package command_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/dockpit/pit/command"
	"github.com/dockpit/pit/reporter"
)

func TestMock(t *testing.T) {
	r := reporter.NewTerminal(os.Stdout)
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	expath := filepath.Join(wd, "test_example", command.ManifestExamplesPath)
	cmd := command.NewMock(r, command.NewInstall(r))

	tdir, err := ioutil.TempDir("", "dp_tinstallinfotmp")
	if err != nil {
		t.Fatal(err)
	}

	os.Setenv("PIT_PATH", tdir)

	AssertCommandNoError(t, cmd, []string{
		"-config", filepath.Join(wd, "test_example"),
		"-examples", expath,
	}, `(?s)Installed.*Mocked.*success.*http`, r)

	//should have called install
	_, err = ioutil.ReadFile(filepath.Join(tdir, "github.com", "dockpit", "pit-token", "main.go"))
	if err != nil {
		t.Fatal(err)
	}

	//unmock afterwards
	cmd2 := command.NewUnmock(r)

	AssertCommandNoError(t, cmd2, []string{"-examples", expath}, `(?s).*`, r)
}
