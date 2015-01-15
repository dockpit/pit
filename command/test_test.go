package command_test

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/dockpit/pit/command"
	"github.com/dockpit/pit/reporter"
)

func mock(t *testing.T, expath, confpath string) {
	r := reporter.NewTerminal(os.Stdout)
	cmd := command.NewMock(r, command.NewInstall(r))
	AssertCommandNoError(t, cmd, []string{"-examples", expath, "-config", confpath}, `(?s)Mocked.*success.*http`, r)
}

func unmock(t *testing.T, expath string) {
	r := reporter.NewTerminal(os.Stdout)
	cmd2 := command.NewUnmock(r)
	AssertCommandNoError(t, cmd2, []string{"-examples", expath}, `(?s).*`, r)
}

func compileSubject(t *testing.T) {
	bcmd := exec.Command("go", "build", "-o=/tmp/test", "test_example/main.go")
	bcmd.Stdout = os.Stdout
	bcmd.Stderr = os.Stderr
	err := bcmd.Run()
	if err != nil {
		t.Fatal(err)
	}
}

func TestTest(t *testing.T) {
	compileSubject(t)

	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	r := reporter.NewTerminal(os.Stdout)
	confpath := filepath.Join(wd, "test_example")
	expath := filepath.Join(confpath, command.ManifestExamplesPath)
	stpath := filepath.Join(confpath, command.ManifestStatesPath)

	test := command.NewTest(r)

	tdir, err := ioutil.TempDir("", "dp_tinstallinfotmp")
	if err != nil {
		t.Fatal(err)
	}

	os.Setenv("PIT_PATH", tdir)

	//mock
	mock(t, expath, confpath)
	defer unmock(t, expath)

	//run test
	AssertCommandNoError(t, test, []string{"-examples", expath, "-states", stpath, "-config", confpath}, `(?s).*`, r)
}

func TestSingleRunTest(t *testing.T) {
	compileSubject(t)

	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	r := reporter.NewTerminal(os.Stdout)
	confpath := filepath.Join(wd, "test_example")
	expath := filepath.Join(confpath, command.ManifestExamplesPath)
	stpath := filepath.Join(confpath, command.ManifestStatesPath)

	test := command.NewTest(r)

	tdir, err := ioutil.TempDir("", "dp_tinstallinfotmp")
	if err != nil {
		t.Fatal(err)
	}

	os.Setenv("PIT_PATH", tdir)

	//mock
	mock(t, expath, confpath)
	defer unmock(t, expath)

	//run test
	AssertCommandNoError(t, test, []string{"-examples", expath, "-states", stpath, "-config", confpath, "'list a single user'"}, `(?s).*skipped.*`, r)
}

// @todo write test to test failing test
