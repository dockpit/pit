package spec

import (
	"io"
)

type Loader struct{}

func NewLoader() *Loader {
	return &Loader{}
}

func (l *Loader) Load(loc string) (io.ReadCloser, error) {
	link, err := NewLink(loc)
	if err != nil {
		return nil, err
	}

	return link.Load()
}
