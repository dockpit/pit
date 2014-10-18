package spec_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

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

func parse(data string, t *testing.T) *spec.Case {
	c := &spec.Case{}
	err := json.Unmarshal([]byte(data), c)
	if err != nil {
		t.Fatal(err)
	}

	return c
}

func TestFactoryWhenThenBasicCaseGeneration(t *testing.T) {
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

	err = st.Test(ts.URL, http.DefaultClient)
	if err != nil {
		t.Fatalf("Expected test to succeed, but got %s", err)
	}
}
