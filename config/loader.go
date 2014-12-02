package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Loader struct {
	dir string
}

func NewLoader(dir string) *Loader {
	return &Loader{dir}
}

func (l *Loader) Load() (C, error) {

	//open file
	f, err := os.Open(filepath.Join(l.dir, "dockpit.json"))
	if err != nil {
		return nil, err
	}

	//parse json
	dec := json.NewDecoder(f)
	cd := &ConfigData{}
	err = dec.Decode(cd)
	if err != nil {
		return nil, err
	}

	//domain specific parsing
	c, err := NewConfig(cd)
	if err != nil {
		return nil, err
	}

	return c, nil
}
