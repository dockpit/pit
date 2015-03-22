package model

import (
	"encoding/json"

	"code.google.com/p/go-uuid/uuid"
)

var MetaBucketName = "meta"
var DatabaseMetaKey = "_db"

type Meta struct {
	ID       string            `json:"name"`
	Machines map[string]string `json:"machines"`
}

func NewMeta() (*Meta, error) {
	return &Meta{
		ID:       uuid.New(),
		Machines: map[string]string{},
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
