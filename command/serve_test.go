package command_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/bmizerany/assert"

	"github.com/dockpit/pit/command"
)

func TestServe(t *testing.T) {

	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	out := bytes.NewBuffer(nil)
	cmd := command.NewServe(out)

	//send interrupt signal to intself to gracefully shut server server
	pid := os.Getpid()
	p, err := os.FindProcess(pid)
	if err != nil {
		t.Fatal(err)
	}

	//a routine that stops the server after some time
	go func() {
		<-time.After(time.Millisecond * 100)
		defer func() { p.Signal(os.Interrupt) }()

		resp, err := http.Get("http://localhost:9000/users")
		if err != nil {
			t.Error(err)
		}
		defer resp.Body.Close()

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, 200, resp.StatusCode, fmt.Sprintf("Expected server to return %d, but got %s: %s", 200, resp.Status, string(b)))
		assert.Equal(t, string(b), "[]")
	}()

	AssertCommand(t, cmd, []string{"--bind", ":9000", "-examples", filepath.Join(wd, "..", ".dockpit", "examples")}, `(?s)gracefully shutting down`, out)

}
