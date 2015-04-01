package model

import (
	"code.google.com/p/go-uuid/uuid"
)

type Build struct {
	ID        string  `json:"id"`
	State     State   `json:"state"`
	Dep       Dep     `json:"dep"`
	Error     error   `json:"error"`
	ImageName string  `json:"image_name"`
	IsRunning bool    `json:"is_running"`
	Output    *Output `json:"output"`
}

func NewBuild(d Dep, s State) (*Build, error) {
	return &Build{
		ID:     uuid.New(),
		Dep:    d,
		State:  s,
		Output: NewOutput(),
	}, nil
}
