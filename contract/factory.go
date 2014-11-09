package contract

import (
	"encoding/json"
)

//
type Factory struct{}

func NewFactory() *Factory {
	return &Factory{}
}

func (f *Factory) Draft(loc string) (*Contract, error) {
	data, err := f.Load(loc)
	if err != nil {
		return nil, err
	}

	return NewContract(data)
}

func (f *Factory) Load(loc string) (*ContractData, error) {

	//create as link
	link, err := NewLink(loc)
	if err != nil {
		return nil, err
	}

	//creat reader
	r, err := link.Load()
	if err != nil {
		return nil, err
	}
	defer r.Close()

	//decode into data struct
	c := &ContractData{}
	dec := json.NewDecoder(r)
	err = dec.Decode(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
