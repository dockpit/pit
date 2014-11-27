package command_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	// "github.com/bmizerany/assert"

	"github.com/dockpit/pit/command"
)

func TestTest(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//@todo, call mocked dependencies
		fmt.Fprintf(w, "[]")
	}))

	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	expath := filepath.Join(wd, "..", ".dockpit", "examples")
	stpath := filepath.Join(wd, "..", ".dockpit", "states")
	out := bytes.NewBuffer(nil)

	mock := command.NewMock(out, command.NewInstall(out))
	build := command.NewBuild(out)
	test := command.NewTest(out, mock, build)

	tdir, err := ioutil.TempDir("", "dp_tinstallinfotmp")
	if err != nil {
		t.Fatal(err)
	}

	os.Setenv("PIT_PATH", tdir)

	AssertCommand(t, test, []string{"-examples", expath, "-states", stpath, svr.URL}, `(?s)Tested.*successful`, out)

}
