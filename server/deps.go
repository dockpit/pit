package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dockpit/pit/model"

	"github.com/zenazn/goji/web"
)

func (s *Server) ListDeps(c web.C, w http.ResponseWriter, r *http.Request) {
	deps, err := s.model.GetAllDeps()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get all dependencies: %s", err), http.StatusInternalServerError)
		return
	}

	enc := json.NewEncoder(w)
	err = enc.Encode(deps)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode: {{err}}", err), http.StatusInternalServerError)
	}
}

func (s *Server) ListTemplates(c web.C, w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	err := enc.Encode(s.templates)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode: {{err}}", err), http.StatusInternalServerError)
	}
}

func (s *Server) CreateDep(c web.C, w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var newdep *model.Dep
	err := dec.Decode(&newdep)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed decode bew dep: %s", err), http.StatusBadRequest)
		return
	}

	dep, err := model.NewDep(newdep.Name, newdep.Template)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create dep: %s", err), http.StatusBadRequest)
		return
	}

	err = s.model.InsertDep(dep)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to insert dep: %s", err), http.StatusBadRequest)
		return
	}
}

func (s *Server) RemoveDep(c web.C, w http.ResponseWriter, r *http.Request) {
	id := c.URLParams["dep_id"]
	dep, err := s.model.FindDepByID(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to find dep with id '%s': %s", id, err), http.StatusBadRequest)
		return
	}

	if dep == nil {
		http.NotFound(w, r)
		return
	}

	err = s.model.RemoveDep(dep)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed remove dep with id '%s': %s", id, err), http.StatusInternalServerError)
		return
	}
}
