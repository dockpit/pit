package model

import (
	"encoding/json"
	"fmt"
)

var DepBucketName = "deps"

type Dep struct {
	Name   string   `json:"name"`
	States []*State `json:"states"`
}

func NewDep(name string) (*Dep, error) {
	if name == "" {
		return nil, fmt.Errorf("Dep Validation Error: Name cannot be empty")
	}

	return &Dep{
		Name:   name,
		States: []*State{},
	}, nil
}

func NewDepFromSerialized(data []byte) (*Dep, error) {
	var dep *Dep
	err := json.Unmarshal(data, &dep)

	return dep, err
}

func (i *Dep) Serialize() ([]byte, error) {
	return json.Marshal(i)
}

func (d *Dep) AddState(s *State) {
	d.States = append(d.States, s)
}

func (d *Dep) GetState(sname string) *State {
	for _, state := range d.States {
		if state.Name == sname {
			return state
		}
	}
	return nil
}
