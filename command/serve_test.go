package command_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/bmizerany/assert"

	"github.com/dockpit/dirtar"
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
		t.Fatal(err)
	}

	//write file that the mock should remove
	tpath := filepath.Join(dir, "- users", "'list all users'", "test.json")
	err = os.MkdirAll(filepath.Dir(tpath), 0777)
	if err != nil {
		t.Fatal(err)
	}

	err = ioutil.WriteFile(tpath, []byte{}, 0777)
	if err != nil {
		t.Fatal(err)
	}

	//a routine that stops the server after some time
	go func() {
		<-time.After(time.Millisecond * 100)
		defer func() { p.Signal(os.Interrupt) }()

		//tar current wd for kicks
		data := bytes.NewBuffer(nil)
		err := dirtar.Tar(wd, data)
		if err != nil {
			t.Fatal(err)
		}

		//upload a tar
		resp, err := http.Post("http://localhost:9000/_examples", "application/x-tar", data)
		if err != nil {
			t.Error(err)
		}
		defer resp.Body.Close()

		//assert
		assert.Equal(t, 201, resp.StatusCode, fmt.Sprintf("Expected server to return %d, but got %s", 200, resp.Status))

		//read original
		files, err := ioutil.ReadDir(wd)
		if err != nil {
			t.Fatal(err)
		}

		// check if each file actually was created in the new dir
		// and is not empty
		for _, f := range files {
			fi, err := os.Stat(filepath.Join(dir, f.Name()))
			if err != nil {
				t.Fatal(err)
			}

			//not empty files?
			if fi.Size() == 0 {
				t.Fatal("File should not have been empty")
			}
		}

		//the old file should have been removed
		_, err = os.Stat(tpath)
		assert.Equal(t, true, os.IsNotExist(err), "test file should have been removed")

	}()

	AssertCommand(t, cmd, []string{"--bind", ":9000", "-examples", dir}, `(?s)Stopping`, out)

}
