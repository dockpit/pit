package contract

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/zenazn/goji/web"
)

//
//
//
type Pair struct {
	Request  *http.Request
	Response *http.Response
	While    []While
	Given    map[string]Given
}

func NewPairFromData(data *CaseData) (*Pair, error) {

	//@todo add header
	//@todo add request body
	//@todo add while

	//create request from data
	req, err := http.NewRequest(data.When.Method, data.When.Path, nil)
	if err != nil {
		return nil, err
	}

	//creat response from data
	resp := &http.Response{}
	resp.StatusCode = data.Then.StatusCode
	resp.Body = ioutil.NopCloser(strings.NewReader(data.Then.Body))

	return &Pair{req, resp, data.While, data.Given}, nil
}

func (p *Pair) BelongsToAction(a A) bool {

	//compare HTTP method
	if p.Request.Method == a.Method() {
		return true
	}

	return false
}

func (p *Pair) IsExpectedResponse(resp *http.Response) error {
	var err error

	//assert response code
	if p.Response.StatusCode != resp.StatusCode {
		return fmt.Errorf("StatusCode not equal, expected %d got: %d", p.Response.StatusCode, resp.StatusCode)
	}

	//get expected content
	c1 := []byte{}
	if p.Response.Body != nil {
		buff := bytes.NewBuffer(nil)
		r := io.TeeReader(p.Response.Body, buff)

		c1, err = ioutil.ReadAll(r)
		if err != nil {
			return err
		}

		p.Response.Body = ioutil.NopCloser(buff)
	}

	//get actualy content
	c2 := []byte{}
	if resp.Body != nil {
		buff := bytes.NewBuffer(nil)
		r := io.TeeReader(resp.Body, buff)

		c2, err = ioutil.ReadAll(r)
		if err != nil {
			return err
		}

		resp.Body = ioutil.NopCloser(buff)
	}

	if !bytes.Equal(c1, c2) {
		return fmt.Errorf("Content not equal, expected %s got: %s", string(c1), string(c2))
	}

	// @todo assert other headers

	return nil
}

func (p *Pair) IsSuccessLike() bool {
	return p.Response.StatusCode >= 200 && p.Response.StatusCode < 300
}

func (p *Pair) GenerateHandler() web.Handler {
	return web.HandlerFunc(func(ctx web.C, w http.ResponseWriter, r *http.Request) {

		//write headers by example
		w.WriteHeader(p.Response.StatusCode)

		//copy body without consuming the original
		if p.Response.Body != nil {
			buff := bytes.NewBuffer(nil)
			r := io.TeeReader(p.Response.Body, buff)
			io.Copy(w, r)

			//'reset' original response body
			p.Response.Body = ioutil.NopCloser(buff)
		}

		// @todo write other headers

	})
}

func (p *Pair) GenerateTest() TestFunc {
	return func(host string, client *http.Client) error {

		//copy request from example pair
		req := *p.Request

		//parse overwrite host url
		h, err := url.Parse(host)
		if err != nil {
			return err
		}

		//overwrite generated with test specific host/scheme
		req.URL.Host = h.Host
		req.URL.Scheme = h.Scheme

		//do the actual request
		resp, err := client.Do(&req)
		if err != nil {
			return err
		}

		//let the pair assert itself
		if err := p.IsExpectedResponse(resp); err != nil {
			return UnexpectedResponseError(p.Response, resp, err)
		}

		return nil
	}
}

//
//
//
type Action struct {
	pairs  []*Pair
	method string
}

func NewAction(p *Pair) *Action {
	return &Action{
		pairs:  []*Pair{p},
		method: p.Request.Method,
	}
}

func (a *Action) Pairs() []*Pair {
	return a.pairs
}

func (a *Action) AddPair(p *Pair) {
	a.pairs = append(a.pairs, p)
}

func (a *Action) Method() string {
	return a.method
}

func (a *Action) Tests() []TestFunc {
	tests := []TestFunc{}

	for _, p := range a.pairs {
		tests = append(tests, p.GenerateTest())
	}

	return tests
}

func (a *Action) Handler() (web.Handler, error) {

	//pick the first example that specified a success like response
	var ex *Pair
	for _, p := range a.pairs {
		if p.IsSuccessLike() {
			ex = p
			break
		}
	}

	if ex == nil {
		return nil, MockingError(fmt.Sprintf("%s Action has no 'success-like' example", a.method))
	}

	//return a handler that simply returns the example
	return ex.GenerateHandler(), nil
}

//
//
//
type Resource struct {
	pattern string
	cases   []*Pair
}

func NewResource(pattern string, cases ...*Pair) *Resource {
	return &Resource{pattern, cases}
}

func (r *Resource) Pattern() string {
	return r.pattern
}

func (r *Resource) Actions() ([]A, error) {
	actions := []A{}

Cases:
	for _, pair := range r.cases {
		for _, a := range actions {
			if pair.BelongsToAction(a) {

				//add to action
				a.AddPair(pair)

				// and continue to next pair
				continue Cases
			}
		}

		//no existing action was matched, create new one from pair
		actions = append(actions, NewAction(pair))
	}

	return actions, nil
}
