package model

import (
	"code.google.com/p/go-uuid/uuid"
	"encoding/json"
)

var MetaBucketName = "meta"
var DatabaseMetaKey = "_db"
var StatsMetaKey = "_stats"

//db meta struct contains meta information about the db itself
type Meta struct {
	ID string `json:"name"`
}

func NewMeta() (*Meta, error) {
	return &Meta{
		ID: uuid.New(),
	}, nil
}

func NewMetaFromSerialized(data []byte) (*Meta, error) {
	var meta *Meta
	err := json.Unmarshal(data, &meta)

	return meta, err
}

func (m *Meta) Serialize() ([]byte, error) {
	return json.Marshal(m)
}

type Stats struct {
	NrOfDepsCreated int                  `json:"nr_of_deps_created"`
	Achievements    map[string]Milestone `json:"achievements"`
}

func NewStats() (*Stats, error) {
	return &Stats{
		Achievements: map[string]Milestone{},
	}, nil
}

func (s *Stats) UpdateAchievements() {
	for _, ms := range Milestones() {
		unlocked := ms.IsUnlocked(s)
		if _, ok := s.Achievements[ms.ID]; ok && !unlocked {
			//achievement is no longer valid
			delete(s.Achievements, ms.ID)
		} else if !ok && unlocked {
			//unloacked achievement
			s.Achievements[ms.ID] = ms
		}
	}
}

func (s *Stats) Handle(ev Event) bool {
	if _, ok := ev.Subject.(*Dep); ok && ev.Action == ActionCreated {
		s.NrOfDepsCreated++
		return true
	}

	s.UpdateAchievements()
	return false
}

func NewStatsFromSerialized(data []byte) (*Stats, error) {
	var stats *Stats
	err := json.Unmarshal(data, &stats)

	stats.UpdateAchievements()
	return stats, err
}

func (s *Stats) Serialize() ([]byte, error) {
	return json.Marshal(s)
}
