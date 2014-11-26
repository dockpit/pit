package command_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/bmizerany/assert"

	"github.com/dockpit/pit/command"
)

func TestServe(t *testing.T) {

	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	out := bytes.NewBuffer(nil)
	cmd := command.NewServe(out)

	//send interrupt signal to intself to gracefully shut server server
	pid := os.Getpid()
	p, err := os.FindProcess(pid)
	if err != nil {
		t.Fatal(err)
	}

	//a routine that stops the server after some time
	go func() {
		<-time.After(time.Millisecond * 100)
		defer func() { p.Signal(os.Interrupt) }()

		//call the mock
		resp, err := http.Get("http://localhost:9000/users")
		if err != nil {
			t.Error(err)
		}
		defer resp.Body.Close()

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}

		//assert
		assert.Equal(t, 200, resp.StatusCode, fmt.Sprintf("Expected server to return %d, but got %s: %s", 200, resp.Status, string(b)))
		assert.Equal(t, string(b), "[]")

		//fetch recording
		recresp, err := http.Get("http://localhost:9000/_recordings?pattern=%2Fusers&method=GET")
		if err != nil {
			t.Error(err)
		}
		defer recresp.Body.Close()

		b, err = ioutil.ReadAll(recresp.Body)
		if err != nil {
			t.Error(err)
		}

		//assert recording
		assert.Equal(t, 200, recresp.StatusCode)
		assert.Equal(t, string(b), "{\"count\":1}\n")

	}()

	AssertCommand(t, cmd, []string{"--bind", ":9000", "-examples", filepath.Join(wd, "..", ".dockpit", "examples")}, `(?s)Stopping`, out)

}

func TestUploading(t *testing.T) {

	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	out := bytes.NewBuffer(nil)
	cmd := command.NewServe(out)

	//send interrupt signal to intself to gracefully shut server server
	pid := os.Getpid()
	p, err := os.FindProcess(pid)
	if err != nil {
		t.Fatal(err)
	}

	//create an alternative path to read examples from since we will overwrite them
	dir, err := ioutil.TempDir("", "dp_upl_")
	if err != nil {
		t.Fatal()
	}

	//a routine that stops the server after some time
	go func() {
		<-time.After(time.Millisecond * 100)
		defer func() { p.Signal(os.Interrupt) }()

		//tar current wd for kicks
		data, err := command.Tar(wd)
		if err != nil {
			t.Fatal(err)
		}

		//upload a tar
		resp, err := http.Post("http://localhost:9000/_examples", "application/x-tar", data)
		if err != nil {
			t.Error(err)
		}
		defer resp.Body.Close()

		//decode and assert response
		files := []string{}
		dec := json.NewDecoder(resp.Body)
		err = dec.Decode(&files)
		if err != nil {
			t.Fatal(err)
		}

		//assert
		assert.Equal(t, 201, resp.StatusCode, fmt.Sprintf("Expected server to return %d, but got %s", 200, resp.Status))
		assert.NotEqual(t, 0, len(files))

		//check if each file actually was created an exists
		for _, p := range files {
			fi, err := os.Stat(p)
			if err != nil {
				t.Fatal(err)
			}

			//should be untarred in tempdir
			assert.Equal(t, true, strings.HasPrefix(p, dir))

			//not empty files?
			if fi.Size() == 0 {
				t.Fatal("File should not have been empty")
			}
		}

	}()

	AssertCommand(t, cmd, []string{"--bind", ":9000", "-examples", dir}, `(?s)Stopping`, out)

}
