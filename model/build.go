package model

import (
	"bytes"
	"strconv"

	"code.google.com/p/go-uuid/uuid"
)

type Output struct {
	Buffer *bytes.Buffer
}

func NewOutput() *Output {
	return &Output{bytes.NewBuffer(nil)}
}

func (o *Output) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(o.Buffer.String())), nil
}

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
