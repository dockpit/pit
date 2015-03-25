package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zenazn/goji/web"
)

func (s *Server) ListIsolations(c web.C, w http.ResponseWriter, r *http.Request) {
	isos, err := s.model.GetAllIsolations()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get all isolations: %s", err), http.StatusInternalServerError)
		return
	}

	enc := json.NewEncoder(w)
	err = enc.Encode(isos)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode: {{err}}", err), http.StatusInternalServerError)
	}
}

func (s *Server) RemoveIsolation(c web.C, w http.ResponseWriter, r *http.Request) {
	name := c.URLParams["name"]
	iso, err := s.model.FindIsolationByName(name)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to find isolation with name '%s': %s", name, err), http.StatusBadRequest)
		return
	}

	if iso == nil {
		http.NotFound(w, r)
		return
	}

	err = s.model.RemoveIsolation(iso)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed remove isolation with name '%s': %s", name, err), http.StatusInternalServerError)
		return
	}
}

func (s *Server) RenderIsolation(c web.C, w http.ResponseWriter, r *http.Request) {
	name := c.URLParams["name"]
	iso, err := s.model.FindIsolationByName(name)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to find isolation with name '%s': %s", name, err), http.StatusBadRequest)
		return
	}

	if iso == nil {
		http.NotFound(w, r)
		return
	}

	s.view.RenderOneIsolation(w, iso)
}
