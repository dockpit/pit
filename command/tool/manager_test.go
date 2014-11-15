package tool_test

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/bmizerany/assert"

	"github.com/dockpit/pit/command/tool"
)

func TestInstall(t *testing.T) {

	dir, err := ioutil.TempDir("", "dockpit_test")
	if err != nil {
		t.Fatal(err)
	}

	m := tool.NewManager(dir)

	err = m.Install("github.com/golang/example")
	if err != nil {
		t.Fatal(err)
	}

	//check if we can open it
	fname := filepath.Join(dir, "deps", "github.com", "golang", "example", "LICENSE")
	_, err = ioutil.ReadFile(fname)
	if err != nil {
		t.Fatal(err)
	}

	//corrupt the git repository
	err = ioutil.WriteFile(filepath.Join(dir, "deps", "github.com", "golang", "example", ".git", "HEAD"), []byte("a"), 0777)
	if err != nil {
		t.Fatal(err)
	}

	//another call to install
	err = m.Install("github.com/golang/example")
	if err != nil {
		t.Fatal(err)
	}

	// should not throw error since no git interaction is expected
	assert.Equal(t, nil, err)
}

func TestUpsert(t *testing.T) {

	dir, err := ioutil.TempDir("", "dockpit_test")
	if err != nil {
		t.Fatal(err)
	}

	m := tool.NewManager(dir)

	err = m.Upsert("github.com/golang/example")
	if err != nil {
		t.Fatal(err)
	}

	//check if we can open it
	fname := filepath.Join(dir, "deps", "github.com", "golang", "example", "LICENSE")
	_, err = ioutil.ReadFile(fname)
	if err != nil {
		t.Fatal(err)
	}

	//corrupt the git repository
	err = ioutil.WriteFile(filepath.Join(dir, "deps", "github.com", "golang", "example", ".git", "HEAD"), []byte("a"), 0777)
	if err != nil {
		t.Fatal(err)
	}

	//another call to Upsert
	err = m.Upsert("github.com/golang/example")

	// should throw error since git interaction is expected
	assert.NotEqual(t, nil, err)
}
