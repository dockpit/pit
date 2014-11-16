package contract

import (
	"net/http"
)

type Given map[string]string

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
	Method string `json:"method"`
	ID     string `json:"id"`
}

type ResourceData struct {
	Name    string      `json:"name"`
	Pattern string      `json:"pattern"`
	Cases   []*CaseData `json:"cases"`
}

type CaseData struct {
	Given Given   `json:"given"`
	When  When    `json:"when"`
	Then  Then    `json:"then"`
	While []While `json:"while"`
}

//
//
//
type ContractData struct {
	Name      string          `json:"name"`
	Resources []*ResourceData `json:"resources"`
}
