package command_test

import (
	"bytes"
	"testing"

	"github.com/dockpit/pit/command"

	"github.com/codegangsta/cli"
	"github.com/stretchr/testify/assert"
)

func AssertCommandNoError(t *testing.T, cmd command.C, args []string, pattern string, out *bytes.Buffer) {
	app := cli.NewApp()
	app.Flags = cmd.Flags()
	app.Action = cmd.Action()

	//prepend zero-length string
	args = append([]string{""}, args...)

	err := app.Run(args)
	if err != nil {
		t.Fatal(err)
	}

	output := out.String()

	assert.NotContains(t, output, "error", "Error", "ERROR")
	assert.Regexp(t, pattern, output)
}
