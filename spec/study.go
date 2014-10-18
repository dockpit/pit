package spec

import (
	"fmt"
	"net/http"
	"net/url"
)

type TestFunc func(host string, c *http.Client) error
type AssertFunc func(resp *http.Response) error

type Study struct {
	Request *http.Request
	Assert  AssertFunc
	Test    TestFunc
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

	t, err := generateTest(r, a)
	if err != nil {
		return nil, err
	}

	return &Study{r, a, t}, nil
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
func generateTest(req *http.Request, a AssertFunc) (TestFunc, error) {
	return func(host string, c *http.Client) error {

		//parse overwrite host url
		h, err := url.Parse(host)
		if err != nil {
			return err
		}

		//overwrite generated with test specific host/scheme
		req.URL.Host = h.Host
		req.URL.Scheme = h.Scheme

		//do the actual request
		resp, err := c.Do(req)
		if err != nil {
			return err
		}

		//and assert resp
		return a(resp)
	}, nil
}
