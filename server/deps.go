package server

import (
	"encoding/json"
	"fmt"
	"net/http"

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

func (s *Server) RemoveDep(c web.C, w http.ResponseWriter, r *http.Request) {
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

	err = s.model.RemoveDep(dep)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed remove dep with name '%s': %s", name, err), http.StatusInternalServerError)
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
