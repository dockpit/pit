package contract

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
)

type Link struct {
	Raw string
	URL *url.URL

	name string
}

func NewLink(loc string) (*Link, error) {

	//expect an url to parse
	u, err := url.Parse(loc)
	if err != nil {
		return nil, err
	}

	//only support naming by dir for now
	name := path.Base(path.Dir(loc))
	if name == "" {
		return nil, fmt.Errorf("Could not extract name from location %s, expected .../<name>/dockpit.json", loc)
	}

	return &Link{
		Raw:  loc,
		URL:  u,
		name: name,
	}, nil
}

func (l *Link) Name() string {
	return l.name
}

func (l *Link) Load() (io.ReadCloser, error) {
	var err error
	var r io.ReadCloser

	//if given 'url' is not absolute cosider it
	//as a file location local to workingdir
	if !l.URL.IsAbs() {
		r, err = os.Open(l.Raw)
		if err != nil {
			return nil, err
		}
	} else {
		//else fetch using the default http agent
		resp, err := http.Get(l.URL.String())
		if err != nil {
			return nil, err
		}

		r = resp.Body
	}

	return r, nil
}
