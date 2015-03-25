package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dockpit/pit/model"

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

func (s *Server) CreateIsolation(c web.C, w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var newiso *model.Isolation
	err := dec.Decode(&newiso)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed decode new isolation: %s", err), http.StatusBadRequest)
		return
	}

	i, err := model.NewIsolation(newiso.Name)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create isolation: %s", err), http.StatusBadRequest)
		return
	}

	err = s.model.InsertIsolation(i)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to insert isolation: %s", err), http.StatusBadRequest)
		return
	}
}

func (s *Server) RemoveIsolation(c web.C, w http.ResponseWriter, r *http.Request) {
	ID := c.URLParams["id"]
	iso, err := s.model.FindIsolationByID(ID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to find isolation with ID '%s': %s", ID, err), http.StatusBadRequest)
		return
	}

	if iso == nil {
		http.NotFound(w, r)
		return
	}

	err = s.model.RemoveIsolation(iso)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed remove isolation with ID '%s': %s", ID, err), http.StatusInternalServerError)
		return
	}
}

func (s *Server) UpdateIsolation(c web.C, w http.ResponseWriter, r *http.Request) {
	ID := c.URLParams["id"]
	iso, err := s.model.FindIsolationByID(ID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to find isolation with ID '%s': %s", ID, err), http.StatusBadRequest)
		return
	}

	if iso == nil {
		http.NotFound(w, r)
		return
	}

	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var newiso *model.Isolation
	err = dec.Decode(&newiso)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed decode isolation with ID '%s': %s", ID, err), http.StatusBadRequest)
		return
	}

	err = s.model.UpdateIsolation(newiso)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to update isolation with name '%s': %s", ID, err), http.StatusInternalServerError)
		return
	}
}
