package runner

import (
	"github.com/dockpit/lang/manifest"
)

type AllSelector struct{}

func (a *AllSelector) ShouldRun(*manifest.Pair) bool {
	return true
}

type CaseSelector struct {
	Case string
}

func (a *CaseSelector) ShouldRun(p *manifest.Pair) bool {
	return p.Name == a.Case
}

func Parse(sel string) (Selector, error) {
	if sel == "*" {
		return &AllSelector{}, nil
	}

	return &CaseSelector{sel}, nil
}
