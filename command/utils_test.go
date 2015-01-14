package command_test

import (
	"bytes"
	"regexp"
	"strings"
	"testing"

	"github.com/dockpit/pit/command"

	"github.com/codegangsta/cli"
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

	//no error @todo inelegant
	if strings.Contains(output, "error") {
		t.Fatalf("Command Output contained an unexpected error: %s", out)
	}

	//match regexp
	m, err := regexp.MatchString(pattern, output)
	if err != nil {
		t.Fatal(err)
	}

	if !m {
		t.Errorf("Out didn't match expected pattern /%s/, received: '%s'", pattern, out)
	}

}
