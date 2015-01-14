package command_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/dockpit/pit/command"
)

func mock(t *testing.T, expath, confpath string) {
	out := bytes.NewBuffer(nil)
	cmd := command.NewMock(out, command.NewInstall(out))
	AssertCommandNoError(t, cmd, []string{"-examples", expath, "-config", confpath}, `(?s)Mocking.*done!.*http`, out)
}

func unmock(t *testing.T, expath string) {
	out := bytes.NewBuffer(nil)
	cmd2 := command.NewUnmock(out)
	AssertCommandNoError(t, cmd2, []string{"-examples", expath}, `(?s).*`, out)
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

	confpath := filepath.Join(wd, "test_example")
	expath := filepath.Join(confpath, command.ManifestExamplesPath)
	stpath := filepath.Join(confpath, command.ManifestStatesPath)

	out := bytes.NewBuffer(nil)
	test := command.NewTest(out)

	tdir, err := ioutil.TempDir("", "dp_tinstallinfotmp")
	if err != nil {
		t.Fatal(err)
	}

	os.Setenv("PIT_PATH", tdir)

	//mock
	mock(t, expath, confpath)
	defer unmock(t, expath)

	//run test
	AssertCommandNoError(t, test, []string{"-examples", expath, "-states", stpath, "-config", confpath, "http://localhost:8000"}, `(?s).*`, out)

}

func TestSingleRunTest(t *testing.T) {
	compileSubject(t)

	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	confpath := filepath.Join(wd, "test_example")
	expath := filepath.Join(confpath, command.ManifestExamplesPath)
	stpath := filepath.Join(confpath, command.ManifestStatesPath)

	out := bytes.NewBuffer(nil)
	test := command.NewTest(out)

	tdir, err := ioutil.TempDir("", "dp_tinstallinfotmp")
	if err != nil {
		t.Fatal(err)
	}

	os.Setenv("PIT_PATH", tdir)

	//mock
	mock(t, expath, confpath)
	defer unmock(t, expath)

	//run test
	AssertCommandNoError(t, test, []string{"-examples", expath, "-states", stpath, "-config", confpath, "http://localhost:8000", "'list a single user'"}, `(?s).*skipping.*`, out)
}
