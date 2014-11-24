package command_test

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
	"time"

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

		//@todo test server

		p.Signal(os.Interrupt)
	}()

	AssertCommand(t, cmd, []string{"-examples", filepath.Join(wd, "..", ".dockpit", "examples")}, `(?s)gracefully shutting down`, out)

}
