package model

type Template struct {
	ID            string           `json:"id"`
	Name          string           `json:"name"`
	DefaultName   string           `json:"default_name"`
	Category      string           `json:"category"`
	Files         map[string]*File `json:"files"`
	StateSettings *StateSettings   `json:"state_settings"`
}
