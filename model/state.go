package model

import (
	"fmt"
)

const (
	TypeCustom = "CUSTOM"
)

type State struct {
	Name      string            `json:"name"`
	Type      string            `json:"type"`
	Files     map[string]string `json:"files"`
	ImageName string            `json:"image_name"`
}

func NewState(name string, t string) (*State, error) {
	if name == "" || t == "" {
		return nil, fmt.Errorf("State Validation Error: Name and Type cannot be empty")
	}

	return &State{
		Name:  name,
		Files: map[string]string{"Dockerfile": "FROM scratch"},
		Type:  t,
	}, nil
}
