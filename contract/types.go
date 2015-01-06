package contract

import (
	"fmt"
	"io"
	"net/http"

	"github.com/zenazn/goji/web"

	"github.com/dockpit/pit/config"
	"github.com/dockpit/state"
)

func UnexpectedResponseError(req http.Request, exp *http.Response, got *http.Response, err error) error {
	return TestingError(fmt.Sprintf("Unexpected Response from '%s': %s", req.URL.Path, err.Error()))
}

func TestingError(msg string) error { return fmt.Errorf("Test: %s", msg) }
func MockingError(msg string) error { return fmt.Errorf("Mock: %s", msg) }

//
type StateManager interface {
	Build(pname, sname string, out io.Writer) (string, error)
	Start(pname, sname string) (*state.StateContainer, error)
	Stop(pname, sname string) error
}

//
type TestFunc func(host, dhost string, c *http.Client, sm StateManager, conf config.C) error

//
//
//
type A interface {
	AddPair(p *Pair)
	Pairs() []*Pair
	Method() string
	Handler(*http.Request) (web.Handler, error)
	Tests() []TestFunc
}

//
//
//
type R interface {
	Actions() ([]A, error)
	Pattern() string
}

//
//
//
type C interface {
	Name() string
	Resources() ([]R, error)
	States() (map[string][]string, error)
	Dependencies() (map[string][]string, error)
}
