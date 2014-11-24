package contract

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zenazn/goji/web"
)

type Recording struct {
	Count int `json:"count"`
}

//
//
// Represents a mocked contract
type Mock struct {
	contract C

	Recordings map[string]map[string]*Recording
}

func NewMock(c C) *Mock {
	return &Mock{c, make(map[string]map[string]*Recording)}
}

//allows the mock to return current recording through a http response
func (m *Mock) ListRecordings(c web.C, w http.ResponseWriter, r *http.Request) {
	pattern := r.URL.Query().Get("pattern")
	method := r.URL.Query().Get("method")

	//pattern&method is mandatory for retrieving recording
	if pattern == "" || method == "" {
		http.NotFound(w, r)
		return
	}

	var rec *Recording
	var resr map[string]*Recording
	var ok bool

	if resr, ok = m.Recordings[pattern]; !ok {
		http.NotFound(w, r)
		return
	}

	if rec, ok = resr[method]; !ok {
		http.NotFound(w, r)
		return
	}

	//reset after the data is provided
	defer func() {
		resr[method] = &Recording{}
	}()

	encoder := json.NewEncoder(w)
	err := encoder.Encode(rec)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// walk resources, actions and pairs to map all necessary states to
// create a router that mocks the contract
func (m *Mock) Mux() (*web.Mux, error) {
	mux := web.New()

	res, err := m.contract.Resources()
	if err != nil {
		return mux, err
	}

	//look at each resource
	for _, r := range res {

		//lazily create resource recordings
		var resr map[string]*Recording
		var ok bool

		if resr, ok = m.Recordings[r.Pattern()]; !ok {
			resr = make(map[string]*Recording)
			m.Recordings[r.Pattern()] = resr
		}

		//we use the actions to create dynamic middleware
		acs, err := r.Actions()
		if err != nil {
			return mux, err
		}

		//create middleware that routes to the correct example
		fn := func(ctx web.C, w http.ResponseWriter, r *http.Request) {

			//match the request method tot the correct action
			for _, a := range acs {

				//match by method
				if a.Method() == r.Method {

					//lazily create recording
					var rec *Recording
					if rec, ok = resr[a.Method()]; !ok {
						rec = &Recording{}
						resr[a.Method()] = rec
					}

					//ask the action for a handle
					h, err := a.Handler(r)
					if err != nil {
						http.Error(w, fmt.Sprintf("%s", err.Error()), http.StatusInternalServerError)
						return
					}

					//increment recording count
					rec.Count++

					//serve mock and return
					h.ServeHTTP(w, r)
					return
				}
			}
		}

		//route everything that matches the resource pattern to the dynamic middleware
		mux.Handle(r.Pattern(), fn)
	}

	//allow the mock to report resource recordings
	mux.Get("/_recordings", m.ListRecordings)

	return mux, nil
}
