package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dockpit/pit/model"

	"github.com/zenazn/goji/web"
)

func (s *Server) OneDepState(c web.C, w http.ResponseWriter, r *http.Request) {
	name := c.URLParams["name"]
	dep, err := s.model.FindDepByName(name)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to find dep with name '%s': %s", name, err), http.StatusBadRequest)
		return
	}

	if dep == nil {
		http.NotFound(w, r)
		return
	}

	sname := c.URLParams["state_name"]
	state := dep.GetState(sname)
	if state == nil {
		http.NotFound(w, r)
		return
	}

	enc := json.NewEncoder(w)
	err = enc.Encode(state)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode: {{err}}", err), http.StatusInternalServerError)
	}
}

func (s *Server) UpdateDepState(c web.C, w http.ResponseWriter, r *http.Request) {
	name := c.URLParams["name"]
	dep, err := s.model.FindDepByName(name)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to find dep with name '%s': %s", name, err), http.StatusBadRequest)
		return
	}

	if dep == nil {
		http.NotFound(w, r)
		return
	}

	sname := c.URLParams["state_name"]
	state := dep.GetState(sname)
	if state == nil {
		http.NotFound(w, r)
		return
	}

	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var newstate *model.State
	err = dec.Decode(&newstate)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed decode state for updating '%s': %s", sname, err), http.StatusBadRequest)
		return
	}

	dep.UpdateState(newstate)

	err = s.model.UpdateDep(dep)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to update dep: %s", err), http.StatusInternalServerError)
		return
	}
}

func (s *Server) RemoveDepState(c web.C, w http.ResponseWriter, r *http.Request) {
	name := c.URLParams["name"]
	dep, err := s.model.FindDepByName(name)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to find dep with name '%s': %s", name, err), http.StatusBadRequest)
		return
	}

	if dep == nil {
		http.NotFound(w, r)
		return
	}

	sname := c.URLParams["state_name"]
	err = s.model.RemoveDepStateByName(dep, sname)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to remove state from dep: %s", err), http.StatusInternalServerError)
	}
}
