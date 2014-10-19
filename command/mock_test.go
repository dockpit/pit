package command_test

import (
	"bytes"
	"testing"

	"github.com/dockpit/pit/command"
)

func TestSingleSpecMocking(t *testing.T) {

	out := bytes.NewBuffer(nil)
	cmd := command.NewMock(out)

	//@todo implement check
	AssertCommand(t, cmd, []string{"../examples/notes.json"}, `(?s).*`, out)
}
