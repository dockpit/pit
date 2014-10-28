package spec

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"net/http"
	"net/url"
	"os"
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

	//expect to be turned into name
	//	//@todo figure naming from link:
	//	- read from json file
	//	- read from json link
	h := md5.New()
	io.WriteString(h, loc)
	name := hex.EncodeToString(h.Sum(nil))

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
