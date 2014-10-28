package spec_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dockpit/pit/spec"
)

func TestMockMux(t *testing.T) {
	l := spec.NewLoader()
	f := spec.NewFactory(l)

	s, err := f.Create("../examples/notes/dockpit.json")
	if err != nil {
		t.Fatal(err)
	}

	m, err := s.Mock()
	if err != nil {
		t.Fatal(err)
	}

	svr := httptest.NewServer(m.Mux())

	//reacts normally to request
	resp, err := http.Get(fmt.Sprintf("%s/notes", svr.URL))
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode > 201 {
		t.Fatal("Expected mock to succeed")
	}

	//recorded access
	resp, err = http.Get(fmt.Sprintf("%s/_recordings/list_notes/success", svr.URL))
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode > 201 {
		t.Fatal("Expected mock to record")
	}

	rec := &spec.Recording{}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(rec)
	if err != nil {
		t.Fatal(err)
	}

	if rec.Count != 1 {
		t.Fatalf("Expected recording count to be 1, got %d", rec.Count)
	}

	//second call should return 0, at it resets the recordings
	resp, err = http.Get(fmt.Sprintf("%s/_recordings/list_notes/success", svr.URL))
	if err != nil {
		t.Fatal(err)
	}

	dec = json.NewDecoder(resp.Body)
	err = dec.Decode(rec)
	if err != nil {
		t.Fatal(err)
	}

	if rec.Count != 0 {
		t.Fatalf("Expected recording count to be 0, got %d", rec.Count)
	}

	//recorded of non called
	resp, err = http.Get(fmt.Sprintf("%s/_recordings/delete_notes/success", svr.URL))
	if err != nil {
		t.Fatal(err)
	}

	dec = json.NewDecoder(resp.Body)
	err = dec.Decode(rec)
	if err != nil {
		t.Fatal(err)
	}

	if rec.Count != 0 {
		t.Fatalf("Expected recording count to be 0, got %d", rec.Count)
	}

}
