package model

import (
	"encoding/json"
	"fmt"

	"code.google.com/p/go-uuid/uuid"
)

var IsolationBucketName = "isolations"
var DefaultIsolationID = "default"

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

func (i *Isolation) AddDep(dep *Dep, sid string) {
	i.States[dep.ID] = sid
}

func (i *Isolation) RemoveDep(dep *Dep) {
	delete(i.States, dep.ID)
}

func (i *Isolation) HasDepState(dep *Dep, sid string) bool {
	if id, ok := i.States[dep.ID]; !ok {
		return false
	} else {
		if id != sid {
			return false
		}
	}

	return true
}

func (i *Isolation) HasDep(dep *Dep) bool {
	if _, ok := i.States[dep.ID]; !ok {
		return false
	}

	return true
}
