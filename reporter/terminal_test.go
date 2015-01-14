package reporter_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dockpit/pit/reporter"
)

func TestTerminalOutput(t *testing.T) {
	var r reporter.R

	r = reporter.NewTerminal(os.Stdout)

	r.Printf("before enter")
	r.Enter(reporter.Build{}, nil)
	r.Printf("enter once")
	r.Enter(reporter.Build{}, nil)
	r.Printf("second enter")
	r.Exit()
	r.Printf("exit once")
	r.Exit()
	r.Printf("second exit")

	assert.Contains(t, r.String(), "      second enter")
}
