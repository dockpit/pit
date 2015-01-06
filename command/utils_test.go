package command_test

import (
	"bytes"
	"regexp"
	"testing"

	"github.com/dockpit/pit/command"

	"github.com/codegangsta/cli"
)

func AssertCommand(t *testing.T, cmd command.C, args []string, pattern string, out *bytes.Buffer) {
	app := cli.NewApp()
	app.Flags = cmd.Flags()
	app.Action = cmd.Action()

	//prepend zero-length string
	args = append([]string{""}, args...)

	err := app.Run(args)
	if err != nil {
		t.Fatal(err)
	}

	m, err := regexp.MatchString(pattern, out.String())
	if err != nil {
		t.Fatal(err)
	}

	if !m {
		t.Errorf("Out didn't match expected pattern /%s/, received: '%s'", pattern, out)
	}

}
