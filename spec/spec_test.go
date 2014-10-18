package spec_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zenazn/goji/web"

	"github.com/dockpit/pit/spec"
)

var case1 = `{
	"when": {
		"method": "POST",
		"path": "/notes"
	},
	"then": {
		"status_code": 201
	}
}`

var case2 = `{
	"name": "success",
	"when": {
		"method": "POST",
		"path": "/notes"
	},
	"then": {
		"status_code": 200
	},
	"while": [
		{"user_service": ["show_user/success"]}
	]
}`

func parse(data string, t *testing.T) *spec.Case {
	c := &spec.Case{}
	err := json.Unmarshal([]byte(data), c)
	if err != nil {
		t.Fatal(err)
	}

	return c
}

func TestWhenThenBasicTestGeneration(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "test", 201)
	}))

	defer ts.Close()

	var c spec.C
	c = parse(case1, t)

	st, err := c.Study()
	if err != nil {
		t.Fatal(err)
	}

	err = st.Test(ts.URL, http.DefaultClient, map[string]string{})
	if err != nil {
		t.Fatalf("Expected test to succeed, but got %s", err)
	}
}

func TestWhenThenBasicMockGeneration(t *testing.T) {
	var c spec.C
	c = parse(case1, t)

	st, err := c.Study()
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewServer(st.Handler)
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != 201 {
		t.Fatal("Expected mock to simulate correctly")
	}

	if st.Assert(resp) != nil {
		t.Fatal("Asserting response with own assert should succeed")
	}

	if st.Pattern.Match(st.Request, &web.C{}) != true {
		t.Fatal("Pattern Should match its own request")
	}
}

func TestWhileGeneration(t *testing.T) {
	rects := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"count": %d}`, 1)
	}))

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "test", 200)
	}))

	defer ts.Close()

	var c spec.C
	c = parse(case2, t)

	st, err := c.Study()
	if err != nil {
		t.Fatal(err)
	}

	err = st.Test(ts.URL, http.DefaultClient, map[string]string{"user_service": rects.URL})
	if err != nil {
		t.Fatalf("Expected test to succeed, but got %s", err)
	}
}
