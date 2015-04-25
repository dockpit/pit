package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/codegangsta/cli"
	"github.com/hashicorp/errwrap"
	"github.com/mitchellh/go-homedir"
)

var MaxRetry = 10

type Init struct {
	*command
}

func NewInit() *Init {
	return &Init{newCommand()}
}

func (c *Init) Name() string {
	return "init"
}

func (c *Init) Description() string {
	return fmt.Sprintf("Initializes Dockpit into the current direcotry")
}

func (c *Init) Usage() string {
	return "Initialize Dockpit in the current Directory"
}

func (c *Init) Flags() []cli.Flag {
	return []cli.Flag{
		DBFlag,
		DockerHostFlag,
		DockerCertFlag,
	}
}

func (c *Init) Action() func(ctx *cli.Context) {
	return c.command.Action(c.Run)
}

func (i *Init) request(method string) (map[string]interface{}, error) {
	v := map[string]interface{}{}
	url := fmt.Sprintf("https://dockpit-v1.appspot.com/api/%s", method)
	resp, err := http.Get(url)
	if err != nil {
		return v, errwrap.Wrapf(fmt.Sprintf("Failed to send request to '%s': {{err}}", url), err)
	}

	buff := bytes.NewBuffer(nil)
	tee := io.TeeReader(resp.Body, buff)
	defer resp.Body.Close()

	dec := json.NewDecoder(tee)
	err = dec.Decode(&v)
	if err != nil {
		return v, errwrap.Wrapf(fmt.Sprintf("Failed to decode response body '%s': {{err}}", buff.String()), err)
	}

	if v["ok"] != true {
		return v, fmt.Errorf("Remote Error for requesting %s: %s", url, v["error"])
	}

	return v, nil
}

func (c *Init) Run(ctx *cli.Context) error {

	//phase 1: create remote
	fmt.Println("Requesting remote for a new box...")
	data, err := c.request("boxes.launch")
	if err != nil {
		return err
	}

	var ok bool
	var box map[string]interface{}
	if box, ok = data["box"].(map[string]interface{}); !ok {
		return fmt.Errorf("Failed to type assert data: %s", data)
	}

	fmt.Printf("Done! Waiting for box '%s' to be running.", box["id"])
	n := 0
	for {
		fmt.Print(".")

		data, err = c.request(fmt.Sprintf("boxes.status?box=%s", box["id"]))
		if err != nil {
			return err
		}

		var b map[string]interface{}
		if b, ok = data["box"].(map[string]interface{}); !ok {
			return fmt.Errorf("Failed to type assert data: %s", data)
		}

		if b["status"] == "RUNNING" {
			box = b
			break
		}

		if n > MaxRetry {
			return fmt.Errorf("Box didn't seem to start: %s", data)
		}

		<-time.After(time.Second * 5)
		n++
	}

	fmt.Printf("\nDone, Running at %s! Waiting for Docker to start.", box["ip"])
	url := fmt.Sprintf("https://%s:2376/v1.17/images/json", box["ip"])
	for {
		fmt.Print(".")

		_, err := http.Get(url)
		if strings.Contains(err.Error(), "certificate signed by unknown authority") {
			break
		}

		<-time.After(time.Second * 5)
	}

	fmt.Printf("\nDone, Docker available at tcp://%s:2376\n", box["ip"])

	//phase 2: write data to home dir
	var boxid string
	if boxid, ok = box["id"].(string); !ok {
		return fmt.Errorf("Failed to type assert box id: %s", box)
	}

	home, err := homedir.Dir()
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Failed to locate home directory: {{err}}"), err)
	}

	dir := filepath.Join(home, ".dockpit", boxid)
	err = os.MkdirAll(dir, 0711)
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Failed to create box directory '%s': {{err}}", dir), err)
	}

	//ca
	var boxca string
	if boxca, ok = box["ca"].(string); !ok {
		return fmt.Errorf("Failed to type assert boxca: %s", box)
	}

	capath := filepath.Join(dir, "ca.pem")
	ca, err := os.Create(capath)
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Failed to create '%s': {{err}}", capath), err)
	}

	_, err = ca.Write([]byte(boxca))
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Failed to write '%s': {{err}}", capath), err)
	}

	//cert
	var boxcert string
	if boxcert, ok = box["cert"].(string); !ok {
		return fmt.Errorf("Failed to type assert boxcert: %s", box)
	}

	certpath := filepath.Join(dir, "cert.pem")
	cert, err := os.Create(certpath)
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Failed to create '%s': {{err}}", certpath), err)
	}

	_, err = cert.Write([]byte(boxcert))
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Failed to write '%s': {{err}}", certpath), err)
	}

	//key
	var boxkey string
	if boxkey, ok = box["key"].(string); !ok {
		return fmt.Errorf("Failed to type assert boxkey: %s", box)
	}

	keypath := filepath.Join(dir, "key.pem")
	key, err := os.Create(keypath)
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Failed to create key.pem file in '%s': {{err}}", dir), err)
	}

	_, err = key.Write([]byte(boxkey))
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Failed to write key.pem file in '%s': {{err}}", dir), err)
	}

	//json
	conff, err := os.Create(filepath.Join(dir, "box.json"))
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Failed to create box.json file in '%s': {{err}}", dir), err)
	}

	enc := json.NewEncoder(conff)
	err = enc.Encode(box)
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Failed to encode box '%s' to json: {{err}}", box), err)
	}

	//report how to connect
	fmt.Printf("Ready to run: `docker -H tcp://%s:2376 --tlscacert=%s --tlskey=%s --tlscert=%s info`", box["ip"], capath, keypath, certpath)

	return nil
}
