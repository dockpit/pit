package contract

import (
	"net/http"

	"github.com/dockpit/contrast/assert"
)

type Given struct {
	Name string `json:"name"`
}

type When struct {
	Method  string      `json:"method"`
	Path    string      `json:"path"`
	Headers http.Header `json:"headers"`
	Body    string      `json:"body"`
}
type Then struct {
	StatusCode int         `json:"code"`
	Status     string      `json:"status"`
	Headers    http.Header `json:"headers"`
	Body       string      `json:"body"`
}

type While struct {
	ID     string `json:"id"`
	Method string `json:"method"`
	Path   string `json:"path"`
}

type ResourceData struct {
	Pattern string      `json:"pattern"`
	Cases   []*CaseData `json:"cases"`
}

type CaseData struct {
	Name  string           `json:"name"`
	Given map[string]Given `json:"given"`
	When  When             `json:"when"`
	Then  Then             `json:"then"`
	While []While          `json:"while"`
}

//
//
//
type ContractData struct {
	Name       string              `json:"name"`
	Resources  []*ResourceData     `json:"resources"`
	Archetypes []*assert.Archetype `json:"archetypes"`
}
