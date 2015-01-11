package runner

import (
	"fmt"
	"io"
	"net/url"

	"github.com/dockpit/lang/manifest"
	"github.com/dockpit/pit/config"
	"github.com/dockpit/state"
)

type Selector interface {
	ShouldRun(*manifest.Pair) bool
}

type Runner interface {
	Name() string
	Run(config.C, manifest.C, Selector, *state.Manager, *url.URL, *url.URL) error
}

func Create(name string, w io.Writer) (Runner, error) {
	switch name {
	case "default":
		return NewDefault(w), nil
	default:
		return nil, fmt.Errorf("No runner named '%s'", name)
	}
}
