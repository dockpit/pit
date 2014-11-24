package contract_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bmizerany/assert"

	. "github.com/dockpit/pit/contract"
)

var req_userA, _ = http.NewRequest("GET", "/users/21", nil)
var resp_userA = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader(`{"id": "21"}`))}

var req_userB, _ = http.NewRequest("GET", "/users/11", nil)
var resp_userB = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader(`{"id": "11"}`))}

var req_userC, _ = http.NewRequest("GET", "/users/444", nil)
var resp_userC = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader(`{"id": "444"}`))}

var req_userD, _ = http.NewRequest("POST", "/users/13", nil)
var resp_userD = &http.Response{StatusCode: 201}

func TestMapping(t *testing.T) {
	var r R

	//given a set of http cases
	r = NewResource(
		"/users/:user_id",
		&Pair{req_userA, resp_userA, []While{}, map[string]Given{}},
		&Pair{req_userB, resp_userB, []While{}, map[string]Given{}},
		&Pair{req_userC, resp_userC, []While{}, map[string]Given{}},
		&Pair{req_userD, resp_userD, []While{}, map[string]Given{}},
	)

	assert.Equal(t, "/users/:user_id", r.Pattern())
}

func TestHandler(t *testing.T) {

	//given a set of http cases
	r := NewResource(
		"/users/:user_id",
		&Pair{req_userA, resp_userA, []While{}, map[string]Given{}},
		&Pair{req_userB, resp_userB, []While{}, map[string]Given{}},
		&Pair{req_userC, resp_userC, []While{}, map[string]Given{}},
		&Pair{req_userD, resp_userD, []While{}, map[string]Given{}},
	)

	// get actions
	as, err := r.Actions()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2, len(as))

	//Handlers should write correctly
	h1, err := as[1].Handler(nil)
	if err != nil {
		t.Fatal(err)
	}

	svr1 := httptest.NewServer(h1)
	resp, _ := http.Post(svr1.URL, "application/json", nil)
	assert.Equal(t, 201, resp.StatusCode)

	//Handlers should write correctly
	h2, err := as[0].Handler(nil)
	if err != nil {
		t.Fatal(err)
	}

	svr2 := httptest.NewServer(h2)
	resp, _ = http.Get(svr2.URL)

	assert.Equal(t, 200, resp.StatusCode)

}

func TestHandlerContent(t *testing.T) {

	//given a set of http cases
	r := NewResource(
		"/users/:user_id",
		&Pair{req_userA, resp_userA, []While{}, map[string]Given{}},
		&Pair{req_userB, resp_userB, []While{}, map[string]Given{}},
		&Pair{req_userC, resp_userC, []While{}, map[string]Given{}},
		&Pair{req_userD, resp_userD, []While{}, map[string]Given{}},
	)

	// get actions
	as, err := r.Actions()
	if err != nil {
		t.Fatal(err)
	}

	var a A
	a = as[0]

	h1, err := a.Handler(nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.Get(httptest.NewServer(h1).URL)
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	//handler should have written content
	assert.Equal(t, `{"id": "21"}`, string(body))
}

func TestTests(t *testing.T) {

	//given a set of http cases
	r := NewResource(
		"/users/:user_id",
		&Pair{req_userA, resp_userA, []While{}, map[string]Given{}},
		&Pair{req_userB, resp_userB, []While{}, map[string]Given{}},
		&Pair{req_userC, resp_userC, []While{}, map[string]Given{}},
		&Pair{req_userD, resp_userD, []While{}, map[string]Given{}},
	)

	// get actions
	as, err := r.Actions()
	if err != nil {
		t.Fatal(err)
	}

	// get tests
	ts := as[1].Tests()

	// should return one test func
	assert.Equal(t, 1, len(ts))

	//success test
	success := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(201)
	}))

	err = ts[0](success.URL, http.DefaultClient)
	assert.Equal(t, nil, err)

	//failing test
	failing := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}))

	err = ts[0](failing.URL, http.DefaultClient)
	assert.NotEqual(t, nil, err)

}

func TestTestsContent(t *testing.T) {

	//given a set of http cases
	r := NewResource(
		"/users/:user_id",
		&Pair{req_userB, resp_userB, []While{}, map[string]Given{}},
		&Pair{req_userC, resp_userC, []While{}, map[string]Given{}},
		&Pair{req_userD, resp_userD, []While{}, map[string]Given{}},
	)

	// get actions
	as, err := r.Actions()
	if err != nil {
		t.Fatal(err)
	}

	// get tests
	ts := as[0].Tests()

	// should return one test func
	assert.Equal(t, 2, len(ts))

	//success test
	success := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Fprint(w, `{"id": "11"}`)
	}))

	err = ts[0](success.URL, http.DefaultClient)
	assert.Equal(t, nil, err)

	//failing test
	failing := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Fprint(w, `{"id":"11"}`) //very strict byte by byte check
	}))

	err = ts[0](failing.URL, http.DefaultClient)
	assert.NotEqual(t, nil, err)

}
