package runner

import (
	"fmt"
	"net/url"

	"github.com/dockpit/lang/manifest"
	"github.com/dockpit/mock/manager"
	"github.com/dockpit/pit/config"
	"github.com/dockpit/pit/reporter"
	"github.com/dockpit/state"
)

type Selector interface {
	ShouldRun(*manifest.Pair) bool
}

type Runner interface {
	Name() string
	Run(config.C, manifest.M, Selector, *state.Manager, *manager.Manager, *url.URL, *url.URL) (*reporter.Result, error)
}

func Create(name string, r reporter.R) (Runner, error) {
	switch name {
	case "default":
		return NewDefault(r), nil
	default:
		return nil, fmt.Errorf("No runner named '%s'", name)
	}
}
