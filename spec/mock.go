package spec

import (
	"encoding/json"
	"net/http"

	"github.com/zenazn/goji/web"
)

type Recording struct {
	Count int `json:"count"`
}

type Mock struct {
	mux *web.Mux

	Recordings map[string]map[string]*Recording
}

func (m *Mock) Mux() *web.Mux {
	return m.mux
}

func NewMock(s S) (*Mock, error) {
	m := &Mock{
		mux:        web.New(),
		Recordings: make(map[string]map[string]*Recording),
	}

	for _, ep := range s.Endpoints() {

		//lazily create endpoint recordings
		var epr map[string]*Recording
		var ok bool

		if epr, ok = m.Recordings[ep.Name()]; !ok {
			epr = make(map[string]*Recording)
			m.Recordings[ep.Name()] = epr
		}

		for _, c := range ep.Cases() {

			//lazily create recording
			var rec *Recording
			if rec, ok = epr[c.Name()]; !ok {
				rec = &Recording{}
				epr[c.Name()] = rec
			}

			//create study
			st, err := c.Study()
			if err != nil {
				return m, err
			}

			//per case 'middleware' that records request
			fn := func(ctx web.C, w http.ResponseWriter, r *http.Request) {

				//increment count
				rec.Count++

				//server mock handler
				st.Handler.ServeHTTP(w, r)
			}

			m.mux.Handle(st.Pattern, fn)
		}
	}

	m.mux.Get("/_recordings/:endpoint/:case", m.ListRecordings)

	return m, nil
}

func (m *Mock) ListRecordings(c web.C, w http.ResponseWriter, r *http.Request) {
	epname := c.URLParams["endpoint"]
	cname := c.URLParams["case"]

	var rec *Recording
	var ep map[string]*Recording
	var ok bool

	if ep, ok = m.Recordings[epname]; !ok {
		http.NotFound(w, r)
		return
	}

	if rec, ok = ep[cname]; !ok {
		http.NotFound(w, r)
		return
	}

	encoder := json.NewEncoder(w)
	err := encoder.Encode(rec)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
