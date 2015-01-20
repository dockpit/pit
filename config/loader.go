package config

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

type Loader struct {
	dir string
}

func NewLoader(dir string) *Loader {
	return &Loader{dir}
}

func (l *Loader) Path() string {
	return filepath.Join(l.dir, ConfigFile)
}

func (l *Loader) LoadData(r io.Reader) (C, error) {
	//parse json
	dec := json.NewDecoder(r)
	cd := &ConfigData{}
	err := dec.Decode(cd)
	if err != nil {
		return nil, err
	}

	//domain specific parsing
	c, err := Parse(cd)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (l *Loader) Load() (C, error) {

	//open file
	f, err := os.Open(l.Path())
	if err != nil {
		return nil, err
	}

	defer f.Close()
	return l.LoadData(f)
}

func (l *Loader) InitData() ([]byte, error) {
	return []byte(InitConfigData), nil
}
