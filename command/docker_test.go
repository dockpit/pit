package command_test

import (
	"net/url"
	"os"
	"testing"

	"github.com/dockpit/pit/command"
	"github.com/dockpit/pit/spec"
)

func getclient(t *testing.T) command.D {
	h := os.Getenv("DOCKER_HOST")
	if h == "" {
		t.Skip("No DOCKER_HOST env variable setup")
	}

	host, err := url.Parse(h)
	if err != nil {
		t.Fatal(err)
	}

	//change to http connection
	host.Scheme = "https"

	cert := os.Getenv("DOCKER_CERT_PATH")
	if cert == "" {
		t.Skip("No DOCKER_CERT_PATH env variable setup")
	}

	d, err := command.NewDocker(host.String(), cert, "latest")
	if err != nil {
		t.Fatal(err)
	}

	return d
}

func TestRemoveStart(t *testing.T) {
	var d command.D
	d = getclient(t)
	_ = d

	//get speck
	l := spec.NewLoader()
	f := spec.NewFactory(l)
	s, err := f.Create("../examples/notes.json")
	if err != nil {
		t.Fatal(err)
	}

	//fetch service deps
	deps, err := s.Dependencies()
	if err != nil {
		t.Fatal(err)
	}

	//remove all dockpit containers
	err = d.RemoveAll()
	if err != nil {
		t.Fatal(err)
	}

	//start new ones
	err = d.Start(deps)
	if err != nil {
		t.Fatal(err)
	}

}
