package command_test

import (
	"bytes"
	"os"
	"testing"
	"time"

	"github.com/dockpit/pit/command"
)

func TestSingleSpecServing(t *testing.T) {

	out := bytes.NewBuffer(nil)
	cmd := command.NewServer(out)

	//send interrupt signal to intself to gracefully shut server server
	pid := os.Getpid()
	p, err := os.FindProcess(pid)
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		<-time.After(time.Millisecond * 100)
		p.Signal(os.Interrupt)
	}()

	//mock specific spec
	AssertCommand(t, cmd, []string{"--bind", ":9000", "../examples/notes.json"}, `(?s).*9000.*Received interrupt.*`, out)

}
