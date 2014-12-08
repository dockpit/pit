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

func TestTest(t *testing.T) {

	//simulate an implementation that complies the example contract
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/users" {

			//simulate state access
			_, err := mgo.DialWithTimeout("mongodb://192.168.59.103:31000", time.Millisecond*100)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			fmt.Fprintf(w, `[]`+"\n")
		} else if r.URL.Path == "/users/21" {
			fmt.Fprintf(w, `{"id": 21}`+"\n")
		}

		//@todo, call mocked dependencies

	}))

	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	expath := filepath.Join(wd, "..", ".manifest", "examples")
	stpath := filepath.Join(wd, "..", ".manifest", "states")
	out := bytes.NewBuffer(nil)

	test := command.NewTest(out)

	tdir, err := ioutil.TempDir("", "dp_tinstallinfotmp")
	if err != nil {
		t.Fatal(err)
	}

	os.Setenv("PIT_PATH", tdir)

	AssertCommand(t, test, []string{"-examples", expath, "-states", stpath, svr.URL}, `(?s)Tested.*successful`, out)

}
