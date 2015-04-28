package host

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/bgentry/speakeasy"
	"github.com/codegangsta/cli"
	"github.com/hashicorp/errwrap"
	"github.com/mitchellh/go-homedir"
)

var MaxRetry = 10

var BoxesListURL = "https://dockpit-eu.appspot.com/api/boxes.list"
var UsersAuthURL = "https://dockpit-eu.appspot.com/api/users.authenticate"

type BoxesListResponse struct {
	Ok    bool     `json:"ok"`
	Error string   `json:"error,omitempty"`
	Boxes *BoxList `json:"boxes"`
}

type UsersAuthResponse struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error,omitempty"`
	User  *User  `json:"user"`
}

type User struct {
	Token string `json:"token"`
}

type Box struct {
	ID string `json:"id"`
}

type BoxList struct {
	Available []*Box `json:"available"`
	OnDemand  []*Box `json:"on_demand"`
	Owned     []*Box `json:"owned"`
}

type Dockpit struct {
	local *Local

	hostURL string
	certDir string
}

func NewLocalBox() *Box {
	return &Box{
		ID: "__local",
	}
}

func NewDockpit(local *Local) *Dockpit {
	return &Dockpit{local: local}
}

func (d *Dockpit) HostURL() string {
	return d.hostURL
}

func (d *Dockpit) CertDir() string {
	return d.certDir
}

func (d *Dockpit) GetBoxes() (*BoxList, error) {
	resp, err := http.Get(BoxesListURL)
	if err != nil {
		return nil, err
	}

	var listresp BoxesListResponse
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&listresp)
	if err != nil {
		return nil, err
	}

	if !listresp.Ok {
		return nil, fmt.Errorf(listresp.Error)
	}

	return listresp.Boxes, nil
}

func (d *Dockpit) GetToken(email, password string) (string, error) {
	resp, err := http.Get(fmt.Sprintf(UsersAuthURL+"?email=%s&password=%s", email, password))
	if err != nil {
		return "nil", err
	}

	var authresp UsersAuthResponse
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&authresp)
	if err != nil {
		return "", err
	}

	if !authresp.Ok {
		return "", fmt.Errorf(authresp.Error)
	}

	return authresp.User.Token, nil
}

func (d *Dockpit) LoginScan(ctx *cli.Context) (string, string, error) {
	var email string
	fmt.Printf("Enter email: ")

	_, err := fmt.Scanln(&email)
	switch {
	case err != nil && err.Error() != "unexpected newline":
		return "", "", err
	case email == "":
		return "", "", fmt.Errorf("email is required.")
	}

	// NOTE: gopass doesn't support multi-byte chars on Windows
	password, err := speakeasy.Ask("Enter password: ")
	switch {
	case err == nil:
	case err.Error() == "unexpected newline":
		return email, "", fmt.Errorf("password is required.")
	default:
		return email, "", err
	}

	return email, password, nil
}

func (d *Dockpit) request(method string) (map[string]interface{}, error) {
	v := map[string]interface{}{}
	url := fmt.Sprintf("https://dockpit-eu.appspot.com/api/%s", method)
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

func (d *Dockpit) LaunchBox(token, boxid string) (map[string]interface{}, string, error) {

	//phase 1: create remote
	fmt.Println("Requesting remote for a new box...")
	data, err := d.request(fmt.Sprintf("boxes.launch?token=%s&box=%s", token, boxid))
	if err != nil {
		return nil, "", err
	}

	var ok bool
	var box map[string]interface{}
	if box, ok = data["box"].(map[string]interface{}); !ok {
		return nil, "", fmt.Errorf("Failed to type assert data: %s", data)
	}

	fmt.Printf("Done! Waiting for box '%s' to be running.", box["id"])
	n := 0
	for {
		fmt.Print(".")

		data, err = d.request(fmt.Sprintf("boxes.status?token=%s&box=%s", token, box["id"]))
		if err != nil {
			return nil, "", err
		}

		var b map[string]interface{}
		if b, ok = data["box"].(map[string]interface{}); !ok {
			return nil, "", fmt.Errorf("Failed to type assert data: %s", data)
		}

		if b["status"] == "RUNNING" {
			box = b
			break
		}

		if n > MaxRetry {
			return nil, "", fmt.Errorf("Box didn't seem to start: %s", data)
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
	if boxid, ok = box["id"].(string); !ok {
		return nil, "", fmt.Errorf("Failed to type assert box id: %s", box)
	}

	home, err := homedir.Dir()
	if err != nil {
		return nil, "", errwrap.Wrapf(fmt.Sprintf("Failed to locate home directory: {{err}}"), err)
	}

	dir := filepath.Join(home, ".dockpit", boxid)
	err = os.MkdirAll(dir, 0711)
	if err != nil {
		return nil, "", errwrap.Wrapf(fmt.Sprintf("Failed to create box directory '%s': {{err}}", dir), err)
	}

	//ca
	var boxca string
	if boxca, ok = box["ca"].(string); !ok {
		return nil, "", fmt.Errorf("Failed to type assert boxca: %s", box)
	}

	capath := filepath.Join(dir, "ca.pem")
	ca, err := os.Create(capath)
	if err != nil {
		return nil, "", errwrap.Wrapf(fmt.Sprintf("Failed to create '%s': {{err}}", capath), err)
	}

	_, err = ca.Write([]byte(boxca))
	if err != nil {
		return nil, "", errwrap.Wrapf(fmt.Sprintf("Failed to write '%s': {{err}}", capath), err)
	}

	//cert
	var boxcert string
	if boxcert, ok = box["cert"].(string); !ok {
		return nil, "", fmt.Errorf("Failed to type assert boxcert: %s", box)
	}

	certpath := filepath.Join(dir, "cert.pem")
	cert, err := os.Create(certpath)
	if err != nil {
		return nil, "", errwrap.Wrapf(fmt.Sprintf("Failed to create '%s': {{err}}", certpath), err)
	}

	_, err = cert.Write([]byte(boxcert))
	if err != nil {
		return nil, "", errwrap.Wrapf(fmt.Sprintf("Failed to write '%s': {{err}}", certpath), err)
	}

	//key
	var boxkey string
	if boxkey, ok = box["key"].(string); !ok {
		return nil, "", fmt.Errorf("Failed to type assert boxkey: %s", box)
	}

	keypath := filepath.Join(dir, "key.pem")
	key, err := os.Create(keypath)
	if err != nil {
		return nil, "", errwrap.Wrapf(fmt.Sprintf("Failed to create key.pem file in '%s': {{err}}", dir), err)
	}

	_, err = key.Write([]byte(boxkey))
	if err != nil {
		return nil, "", errwrap.Wrapf(fmt.Sprintf("Failed to write key.pem file in '%s': {{err}}", dir), err)
	}

	//json
	conff, err := os.Create(filepath.Join(dir, "box.json"))
	if err != nil {
		return nil, "", errwrap.Wrapf(fmt.Sprintf("Failed to create box.json file in '%s': {{err}}", dir), err)
	}

	enc := json.NewEncoder(conff)
	err = enc.Encode(box)
	if err != nil {
		return nil, "", errwrap.Wrapf(fmt.Sprintf("Failed to encode box '%s' to json: {{err}}", box), err)
	}

	//report how to connect
	fmt.Printf("Ready to run: `docker -H tcp://%s:2376 --tlscacert=%s --tlskey=%s --tlscert=%s info`", box["ip"], capath, keypath, certpath)
	return box, dir, nil
}

func (d *Dockpit) Init(ctx *cli.Context) error {

	// //if local host is up and running dont bother with dockpit hosting
	// if d.local.IsUsable() {
	// 	d.hostURL = d.local.HostURL()
	// 	d.certDir = d.local.CertDir()
	// 	return nil
	// }

	boxes, err := d.GetBoxes()
	if err != nil {
		return err
	}

	options := boxes.Owned
	if len(options) == 0 {
		options = append(options, boxes.Available...)
	}

	if len(options) == 0 {
		options = append(options, boxes.OnDemand...)
	}

	//add local host as an option
	if d.local.IsUsable() {
		options = append(options, NewLocalBox())
	}

	//ask user to select host
	var sel *Box
	for {
		if sel != nil {
			break
		}

		fmt.Printf("The following Docker host are avaliable to you:\n\n")
		for i, box := range options {
			fmt.Printf("\t [%d] %s\n", i+1, box.ID)
		}

		fmt.Printf("\nSelect your host (1, 2, etc) and press ENTER:\n")

		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Printf("\n\nInvalid: %s!\n", err)
		}

		nr, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("\n\nInvalid (%s), please provide a number and press ENTER\n", input)
			continue
		}

		nr = nr - 1
		if nr < 0 || nr >= len(options) {
			fmt.Printf("\n\nGiven nr (%d) is not a valid selection!\n", nr)
			continue
		}

		sel = options[nr]
	}

	//to token provided, please login
	token := ctx.String("token")
	if token == "" {
		email, password, err := d.LoginScan(ctx)
		if err != nil {
			return err
		}

		token, err = d.GetToken(email, password)
		if err != nil {
			return err
		}

	}

	bbox, dir, err := d.LaunchBox(token, sel.ID)
	if err != nil {
		return err
	}

	d.hostURL = "tcp://" + bbox["ip"].(string) + ":2376"
	d.certDir = dir
	return nil
}
