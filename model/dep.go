package model

import (
	"encoding/json"
	"fmt"

	"code.google.com/p/go-uuid/uuid"
)

var DepBucketName = "deps"

type Dep struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	States   []*State  `json:"states"`
	Template *Template `json:"template"`
}

func NewDep(name string, t *Template) (*Dep, error) {
	if name == "" {
		return nil, fmt.Errorf("Dep Validation Error: Name cannot be empty")
	}

	return &Dep{
		ID:       uuid.New(),
		Name:     name,
		States:   []*State{},
		Template: t,
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

func (d *Dep) GetState(sid string) *State {
	for _, state := range d.States {
		if state.ID == sid {
			return state
		}
	}
	return nil
}

func (d *Dep) UpdateState(oldstate *State, newstate *State) {
	newstates := []*State{}
	for _, state := range d.States {
		if state.ID == oldstate.ID {
			newstates = append(newstates, newstate)
			continue
		}

		newstates = append(newstates, state)
	}

	d.States = newstates
}
