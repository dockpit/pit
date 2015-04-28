package host

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/bgentry/speakeasy"
	"github.com/codegangsta/cli"
)

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

	//@todo, login / register?, but be logged in
	fmt.Println(sel.ID, token)
	<-time.After(time.Second * 30)

	return nil
}
