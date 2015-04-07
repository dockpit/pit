package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dockpit/pit/model"

	"github.com/zenazn/goji/web"
)

func (s *Server) GetStats(c web.C, w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	err := enc.Encode(s.model.Stats)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode: {{err}}", err), http.StatusInternalServerError)
	}
}

func (s *Server) PutStats(c web.C, w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var newstats *model.Stats
	err := dec.Decode(&newstats)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed decode new stats: %s", err), http.StatusBadRequest)
		return
	}

	for _, na := range newstats.Achievements {
		for _, naa := range s.model.Stats.Achievements {
			if naa.ID == na.ID {
				naa.Shown = na.Shown

				//update memory stats
				s.model.Stats.Achievements[naa.ID] = naa
			}
		}
	}

	err = s.model.PersistStats()
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to persist new stats: %s", err), http.StatusInternalServerError)
	}

	enc := json.NewEncoder(w)
	err = enc.Encode(s.model.Stats)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode: %s", err), http.StatusInternalServerError)
	}
}
