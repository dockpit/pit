package runner

import (
	"fmt"
	"io"
	"net/url"

	"github.com/dockpit/pit/config"
	"github.com/dockpit/pit/contract"
	"github.com/dockpit/state"
)

type Suite interface {
	SetUp() error
	Run() error
	TearDown() error
}

type Runner interface {
	Name() string
	Run(config.C, contract.C, *state.Manager, *url.URL, *url.URL) error
}

func Create(name string, w io.Writer) (Runner, error) {
	switch name {
	case "default":
		return NewDefault(w), nil
	default:
		return nil, fmt.Errorf("No runner named '%s'", name)
	}
}
