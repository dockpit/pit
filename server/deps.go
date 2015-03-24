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
