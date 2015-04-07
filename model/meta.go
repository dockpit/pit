package model

import (
	"encoding/json"

	"code.google.com/p/go-uuid/uuid"
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

//stats contain db lifetime counters for certain actions
type Stats struct {
	NrOfDepsCreated int `json:"nr_of_deps_created"`
}

func NewStats() (*Stats, error) {
	return &Stats{}, nil
}

func (s *Stats) Handle(ev Event) bool {
	if _, ok := ev.Subject.(*Dep); ok && ev.Action == ActionCreated {
		s.NrOfDepsCreated++
		return true
	}

	return false
}

func NewStatsFromSerialized(data []byte) (*Stats, error) {
	var stats *Stats
	err := json.Unmarshal(data, &stats)

	return stats, err
}

func (s *Stats) Serialize() ([]byte, error) {
	return json.Marshal(s)
}
