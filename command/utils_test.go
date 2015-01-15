package command_test

import (
	"testing"

	"github.com/dockpit/pit/command"

	"github.com/codegangsta/cli"
	"github.com/dockpit/pit/reporter"
	"github.com/stretchr/testify/assert"
)

func AssertCommandNoError(t *testing.T, cmd command.C, args []string, pattern string, r reporter.R) {
	app := cli.NewApp()
	app.Flags = cmd.Flags()
	app.Action = cmd.Action()

	//prepend zero-length string
	args = append([]string{""}, args...)

	err := app.Run(args)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 0, r.StatusCode(), "Expected status code 0")
	assert.Regexp(t, pattern, r.String())
}
