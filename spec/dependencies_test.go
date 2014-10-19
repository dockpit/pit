package spec_test

import (
	"testing"

	"github.com/dockpit/pit/spec"
)

func TestDependencyMapping(t *testing.T) {

	l := spec.NewLoader()
	f := spec.NewFactory(l)

	s, err := f.Create("../examples/notes.json")
	if err != nil {
		t.Fatal(err)
	}

	deps, err := s.Dependencies()
	if err != nil {
		t.Fatal(err)
	}

	if len(deps.Map) != 1 {
		t.Fatal("Expected one service it depends on")
	}

	if len(deps.Map["user_service"]) != 1 {
		t.Fatal("Expected user_service dep to have one endpoint")
	}

	if len(deps.Map["user_service"]["show_user"]) != 1 {
		t.Fatal("Expected user_service/show_user dep to have one case")
	}

	if deps.Map["user_service"]["show_user"]["success"] != true {
		t.Fatal("Expected user_service/show_user/success to be true")
	}

}
