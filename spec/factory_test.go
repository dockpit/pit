package spec_test

import (
	"testing"

	"github.com/dockpit/pit/spec"
)

func TestFactoryLoadingLocalFile(t *testing.T) {
	var f spec.F
	l := spec.NewLoader()
	f = spec.NewFactory(l)

	s, err := f.Create("../examples/notes.json")
	if err != nil {
		t.Fatal(err)
	}

	if len(s.Endpoints()) == 0 {
		t.Fatalf("Expected %d endpoints, received %d", 3, len(s.Endpoints()))
	}

	if len(s.Endpoints()[0].Cases()) == 0 {
		t.Fatalf("Expected endpoint to have %d cases, received %d", 1, len(s.Endpoints()[0].Cases()))
	}

	if s.Endpoints()[0].Name() != "list_notes" {
		t.Fatalf("Expected correct endpoint name")
	}
}
