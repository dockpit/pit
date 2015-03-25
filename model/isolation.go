package model

import (
	"encoding/json"
	"fmt"

	"code.google.com/p/go-uuid/uuid"
)

var IsolationBucketName = "isolations"

type Isolation struct {
	ID     string            `json:"id"`
	Name   string            `json:"name"`
	States map[string]string `json:"states"`
}

func NewIsolation(name string) (*Isolation, error) {
	if name == "" {
		return nil, fmt.Errorf("Isolation Validation Error: Name cannot be empty")
	}

	return &Isolation{
		ID:     uuid.New(),
		Name:   name,
		States: map[string]string{},
	}, nil
}

func NewIsolationFromSerialized(data []byte) (*Isolation, error) {
	var iso *Isolation
	err := json.Unmarshal(data, &iso)

	return iso, err
}

func (i *Isolation) Serialize() ([]byte, error) {
	return json.Marshal(i)
}

func (i *Isolation) AddDep(dep *Dep, state string) {
	i.States[dep.Name] = state
}

func (i *Isolation) RemoveDep(dep *Dep) {
	delete(i.States, dep.Name)
}

func (i *Isolation) HasDepState(dep *Dep, sname string) bool {
	if name, ok := i.States[dep.Name]; !ok {
		return false
	} else {
		if sname != name {
			return false
		}
	}

	return true
}

func (i *Isolation) HasDep(dep *Dep) bool {
	if _, ok := i.States[dep.Name]; !ok {
		return false
	}

	return true
}
