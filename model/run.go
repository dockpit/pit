package model

import (
	"code.google.com/p/go-uuid/uuid"
)

type Run struct {
	ID          string  `json:"id"`
	State       State   `json:"state"`
	ContainerID string  `json:"container_id"`
	Error       error   `json:"error"`
	IsReady     bool    `json:"is_ready"`
	Output      *Output `json:"output"`
}

func NewRun(s State) (*Run, error) {
	return &Run{
		ID:     uuid.New(),
		State:  s,
		Output: NewOutput(),
	}, nil
}
