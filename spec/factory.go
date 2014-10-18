package spec

import (
	"encoding/json"
)

type Factory struct {
	Loader L
}

func NewFactory(l L) *Factory {
	return &Factory{l}
}

func (f *Factory) Create(loc string) (S, error) {
	r, err := f.Loader.Load(loc)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	s := &Spec{}

	dec := json.NewDecoder(r)
	err = dec.Decode(s)
	if err != nil {
		return nil, err
	}

	return s, nil
}
