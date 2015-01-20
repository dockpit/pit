package command_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/dockpit/pit/command"
	"github.com/dockpit/pit/reporter"

	"github.com/stretchr/testify/assert"
)

func TestInitIntoTmp(t *testing.T) {

	r := reporter.NewTerminal(os.Stdout)
	cmd := command.NewInit(r)
	tdir, err := ioutil.TempDir("", "dp_tinittmp")
	if err != nil {
		t.Fatal(err)
	}

	AssertCommandNoError(t, cmd, []string{tdir}, `(?s)Initialized configuration file`, r)

	//should be able to read config file from installation
	_, err = ioutil.ReadFile(filepath.Join(tdir, "dockpit.json"))
	if err != nil {
		t.Fatal(err)
	}

	_, err = ioutil.ReadDir(filepath.Join(tdir, ".manifest", "examples"))
	assert.NoError(t, err)

	_, err = ioutil.ReadDir(filepath.Join(tdir, ".manifest", "states"))
	assert.NoError(t, err)
}
