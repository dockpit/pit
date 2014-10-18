package spec

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/zenazn/goji/web"
)

type TestFunc func(host string, c *http.Client, depAddr map[string]string) error
type AssertFunc func(resp *http.Response) error

//Generated pattern for matching mock requests to the correct mock handler
type Pattern struct {
	Prototype *http.Request
}

func (p *Pattern) Prefix() string {
	return "/"
}

// Match the received request to the handler by examining the prototype request
func (p *Pattern) Match(r *http.Request, c *web.C) bool {

	if r.URL.Path != p.Prototype.URL.Path {
		return false
	}

	if r.Method != p.Prototype.Method {
		return false
	}

	return true
}

func (p *Pattern) Run(r *http.Request, c *web.C) {}

//allows for analysing service endpoint cases
type Study struct {
	Request *http.Request
	Assert  AssertFunc
	Test    TestFunc
	Handler web.HandlerFunc
	Pattern web.Pattern
}

func NewStudy(c *Case) (*Study, error) {
	r, err := generateRequest(c)
	if err != nil {
		return nil, err
	}

	a, err := generateAssert(c)
	if err != nil {
		return nil, err
	}

	t, err := generateTest(c, r, a)
	if err != nil {
		return nil, err
	}

	h, err := generateHandler(c)
	if err != nil {
		return nil, err
	}

	p, err := generatePattern(r)
	if err != nil {
		return nil, err
	}

	return &Study{r, a, t, h, p}, nil
}

// Generate a routing pattern that always matches the given request
func generatePattern(r *http.Request) (*Pattern, error) {
	return &Pattern{r}, nil
}

// Generate the http handler function that writes the expected response based on the specification
func generateHandler(c *Case) (web.HandlerFunc, error) {
	return web.HandlerFunc(func(ctx web.C, w http.ResponseWriter, r *http.Request) {

		// @todo write mock handlers more sofisticaed according to spec

		w.WriteHeader(c.Then.StatusCode)

	}), nil
}

// Generates an request to be send to the subject based on the case provided by the spec
func generateRequest(c *Case) (*http.Request, error) {
	l, err := url.Parse("/")
	if err != nil {
		return nil, err
	}

	// @todo specify more then just the path

	l.Path = c.When.Path

	return http.NewRequest(c.When.Method, l.String(), nil)
}

// Generate an assertion function that checks the response returned by the subject
func generateAssert(c *Case) (AssertFunc, error) {
	return func(resp *http.Response) error {

		// @todo make checks more sofisticated based on spec

		if c.caseData.Then.StatusCode != resp.StatusCode {
			return fmt.Errorf("Receiver status code %d, expected %d", c.caseData.Then.StatusCode, resp.StatusCode)
		}

		return nil
	}, nil
}

// Generates a testing function using the request, and assert func
func generateTest(c *Case, req *http.Request, a AssertFunc) (TestFunc, error) {
	return func(host string, client *http.Client, depAddr map[string]string) error {

		//parse overwrite host url
		h, err := url.Parse(host)
		if err != nil {
			return err
		}

		//overwrite generated with test specific host/scheme
		req.URL.Host = h.Host
		req.URL.Scheme = h.Scheme

		//do the actual request
		resp, err := client.Do(req)
		if err != nil {
			return err
		}

		//check dependency recordings
		for _, dep := range c.caseData.While {
			for loc, paths := range dep {
				var addr string
				var ok bool

				if addr, ok = depAddr[loc]; !ok {
					return fmt.Errorf("Test for case '%s' didn't get an addr for dependency '%s'", c.Name(), loc)
				}

				for _, path := range paths {

					//get recorded access
					resp, err = http.Get(fmt.Sprintf("%s/_recordings/%s", addr, path))
					if err != nil {
						return err
					}

					if resp.StatusCode > 201 {
						return fmt.Errorf("Mock recording doesn't have data for '%s', returned %d", path, resp.StatusCode)
					}

					rec := &Recording{}
					dec := json.NewDecoder(resp.Body)
					err = dec.Decode(rec)
					if err != nil {
						return err
					}

					//dont actually count count
					if rec.Count < 1 {
						return fmt.Errorf("Expected dependency '%s#%s' to be called at least once", loc, path)
					}
				}

			}
		}

		//and assert resp
		return a(resp)
	}, nil
}
