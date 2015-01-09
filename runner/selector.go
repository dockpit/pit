package runner

import (
	"fmt"

	"github.com/dockpit/lang"
	"github.com/dockpit/pit/contract"
)

//
type AllSelector struct{}

func (a *AllSelector) ShouldRun(*contract.Pair) bool {
	return true
}

//
type CaseSelector struct {
	Case string
}

func (a *CaseSelector) ShouldRun(p *contract.Pair) bool {
	return p.Name == a.Case
}

func Parse(sel string) (Selector, error) {
	if sel == "*" {
		return &AllSelector{}, nil
	}

	m := lang.CaseEX.FindStringSubmatch(sel)
	if m == nil {
		return nil, fmt.Errorf("Selectur must be a casename enclused in quotes (e.g 'my case'), got: %s", sel)
	}

	return &CaseSelector{m[1]}, nil
}
