package spec

import (
	"fmt"
	"io"
	"net/url"
	"os"
)

type Loader struct{}

func NewLoader() *Loader {
	return &Loader{}
}

func (l *Loader) Load(loc string) (io.Reader, error) {
	var r io.Reader

	//check if its an url
	u, err := url.Parse(loc)
	if err != nil {
		return nil, err
	}

	//if given 'url' is not absolute cosider it
	//as a file location local to workingdir
	if !u.IsAbs() {
		r, err = os.Open(loc)
		if err != nil {
			return nil, err
		}
	} else {
		//@todo implement loading from url
		return nil, fmt.Errorf("Not implemented yet")
	}

	return r, nil
}
