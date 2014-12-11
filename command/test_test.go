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
	"time"

	"labix.org/v2/mgo"
	// "github.com/bmizerany/assert"

	"github.com/dockpit/pit/command"
)

func mock(t *testing.T, expath, confpath string) {
	out := bytes.NewBuffer(nil)
	cmd := command.NewMock(out, command.NewInstall(out))
	AssertCommand(t, cmd, []string{"-examples", expath, "-config", confpath}, `(?s)Mocking.*done!.*http`, out)
}

func unmock(t *testing.T, expath string) {
	out := bytes.NewBuffer(nil)
	cmd2 := command.NewUnmock(out)
	AssertCommand(t, cmd2, []string{"-examples", expath}, `(?s)Unmocked`, out)
}

func TestTest(t *testing.T) {

	//simulate an implementation that complies the example contract
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/users" {

			//call dependency
			_, err := http.Get("http://192.168.59.103:4321/customers")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			//simulate state access
			_, err = mgo.DialWithTimeout("mongodb://192.168.59.103:31000", time.Millisecond*100)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			fmt.Fprintf(w, `[]`+"\n")
		} else if r.URL.Path == "/users/21" {
			fmt.Fprintf(w, `{"id": 21}`+"\n")
		}

	}))

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

	//run test
	AssertCommand(t, test, []string{"-examples", expath, "-states", stpath, "-config", confpath, svr.URL}, `(?s)Tested.*successful`, out)

	//unmock
	unmock(t, expath)

}
