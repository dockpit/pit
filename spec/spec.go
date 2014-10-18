package spec

import (
	"encoding/json"
)

type Conditions struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}

type Expectations struct {
	StatusCode int `json:"status_code"`
}

type Dependency struct {
	Service string `json:"service"`
}

//Service endpoint case specification
type Case struct {
	caseData
}

func (c *Case) Study() (*Study, error) {
	return NewStudy(c)
}

func (c *Case) Name() string {
	return c.caseData.Name
}

func (c *Case) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &c.caseData)
}

type caseData struct {
	Name  string        `json:"name"`
	When  *Conditions   `json:"when"`
	Then  *Expectations `json:"then"`
	While []*Dependency `json:"while"`
}

// Service endpoint specification
type endpointData struct {
	Name  string  `json:"name"`
	Cases []*Case `json:"cases"`
}

type Endpoint struct {
	endpointData
}

func (e *Endpoint) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &e.endpointData)
}

func (e *Endpoint) Name() string {
	return e.endpointData.Name
}

func (e *Endpoint) Cases() []C {
	res := []C{}
	for _, c := range e.endpointData.Cases {
		res = append(res, c)
	}

	return res
}

// Service specification
type Spec struct {
	specData
}

type specData struct {
	Endpoints []*Endpoint `json:"endpoints"`
}

func (s *Spec) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &s.specData)
}

func (s *Spec) Endpoints() []EP {
	res := []EP{}
	for _, ep := range s.specData.Endpoints {
		res = append(res, ep)
	}

	return res
}

func (s *Spec) Mock() (M, error) {
	return NewMock(s)
}
