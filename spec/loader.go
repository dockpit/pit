package spec

import (
	"io"
	"net/http"
	"net/url"
	"os"
)

type Loader struct{}

func NewLoader() *Loader {
	return &Loader{}
}

func (l *Loader) Load(loc string) (io.ReadCloser, error) {
	var r io.ReadCloser

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
		//else fetch using the default http agent
		resp, err := http.Get(u.String())
		if err != nil {
			return nil, err
		}

		r = resp.Body
	}

	return r, nil
}
