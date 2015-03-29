package client

import (
	"archive/tar"
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/dockpit/iowait"
	"github.com/hashicorp/errwrap"
	"github.com/samalba/dockerclient"

	"github.com/dockpit/pit/model"
)

type Docker struct {
	Host string
	Info *dockerclient.Info

	client *dockerclient.DockerClient
	model  *model.Model
}

func NewDocker(m *model.Model, host, cert string) (*Docker, error) {
	d := &Docker{
		Host:  host,
		model: m,
	}

	//parse docker host addr as url
	hurl, err := url.Parse(host)
	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Failed to parse Docker host '%s' as url: {{err}}", host), err)
	}

	//use tlsc?
	var tlsc tls.Config
	if cert != "" {
		c, err := tls.LoadX509KeyPair(filepath.Join(cert, "cert.pem"), filepath.Join(cert, "key.pem"))
		if err != nil {
			return nil, errwrap.Wrapf(fmt.Sprintf("Failed to load Docker cert.pem and key.pem at '%s': {{err}}", cert), err)
		}

		tlsc.Certificates = append(tlsc.Certificates, c)
		tlsc.InsecureSkipVerify = true //@todo switch to secure with docker ca.pem
	}

	//create docker client
	d.client, err = dockerclient.NewDockerClient(hurl.String(), &tlsc)
	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Failed to create Docker client for '%s': {{err}}", host), err)
	}

	d.Info, err = d.client.Info()
	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Failed to retrieve Docker host info for '%s': {{err}}", host), err)
	}

	return d, nil
}

func (d *Docker) RemoveAll() error {
	cs, err := d.client.ListContainers(true, false, "")
	if err != nil {
		return err
	}

	for _, c := range cs {
		remove := false
		for _, n := range c.Names {
			if strings.HasPrefix(n[1:], "dockpit_") {
				remove = true
			}
		}

		if remove {
			err := d.client.RemoveContainer(c.Id, true)
			if err != nil {
				return errwrap.Wrapf(fmt.Sprintf("Failed to remove existing container %s %s: {{err}}", c.Id, c.Names), err)
			}
		}
	}

	return nil
}

func (d *Docker) Switch(iso *model.Isolation) error {
	err := d.RemoveAll()
	if err != nil {
		return err
	}

	//start all for the current isolation
	depst, err := d.model.GetIsolationDepsAndStates(iso)
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Failed to get deps and states for starting iso '%s': {{err}}", iso.Name), err)
	}

	for dep, state := range depst {

		id, err := d.client.CreateContainer(&dockerclient.ContainerConfig{Image: state.ImageName}, state.ImageName)
		if err != nil {
			return errwrap.Wrapf(fmt.Sprintf("Failed to create state container with image '%s': {{err}}, are your states build?", state.ImageName), err)
		}

		_ = dep

		bindings := map[string][]dockerclient.PortBinding{}
		// @todo add port bindings config
		// for _, p := range spconf.Ports() {
		// 	bindings[p.Container+"/tcp"] = []dockerclient.PortBinding{
		// 		dockerclient.PortBinding{"0.0.0.0", p.Host},
		// 	}
		// }

		err = d.client.StartContainer(id, &dockerclient.HostConfig{PortBindings: bindings})
		if err != nil {
			return errwrap.Wrapf(fmt.Sprintf("Failed to start state container with image '%s': {{err}}, are your states build?", state.ImageName), err)
		}

		rc, err := d.client.ContainerLogs(id, &dockerclient.LogOptions{Follow: true, Stdout: true, Stderr: true})
		if err != nil {
			return errwrap.Wrapf(fmt.Sprintf("Failed to follow logs of state container '%s': {{err}}, are your states build?", id), err)
		}
		defer rc.Close()

		//@todo fetch timeout from data
		//@todo fetch regexp from data
		err = iowait.WaitForRegexp(rc, regexp.MustCompile(`.*`), time.Second*5)
		if err != nil {
			return fmt.Errorf("Failed to wait for state container %s: %s", id, err)
		}
	}

	return nil
}

func (d *Docker) Build(b *model.Build) (string, error) {
	dep := b.Dep
	state := b.State
	out := b.Output.Buffer

	//create writer for tar data
	f := bytes.NewReader([]byte(state.Files["Dockerfile"]))
	in := bytes.NewBuffer(nil)
	tw := tar.NewWriter(in)
	hdr := &tar.Header{
		Name: "Dockerfile",
		Size: int64(f.Len()),
	}

	err := tw.WriteHeader(hdr)
	if err != nil {
		return "", errwrap.Wrapf(fmt.Sprintf("Failed to write tar file header for '%s'::'%s': {{err}}", dep.Name, state.Name), err)
	}

	if _, err = io.Copy(tw, f); err != nil {
		return "", errwrap.Wrapf(fmt.Sprintf("Failed to tar Dockerfile for '%s'::'%s': {{err}}", dep.Name, state.Name), err)
	}

	// generate an unique image name based on the provider name and path to the state folder
	iname := fmt.Sprintf("dockpit_%s_%s", dep.Name, state.Name)

	// fall back to streaming ourselves for building the image, samalba has yet to
	// implement image building: https://github.com/samalba/dockerclient/issues/62
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/build?t=%s", d.client.URL, iname), in)
	if err != nil {
		return "", errwrap.Wrapf(fmt.Sprintf("Failed to create Docker build request for '%s'::'%s': {{err}}", dep.Name, state.Name), err)
	}

	req.Header.Set("Content-Type", "application/tar")

	b.IsRunning = true
	resp, err := d.client.HTTPClient.Do(req)
	if err != nil {
		return "", errwrap.Wrapf(fmt.Sprintf("Docker build request failed for '%s'::'%s': {{err}}", dep.Name, state.Name), err)
	}

	b.IsRunning = false

	defer resp.Body.Close()
	defer io.Copy(out, resp.Body)
	if resp.StatusCode > 200 {
		return "", fmt.Errorf("Unexpected response while building Docker image for '%s'::'%s': %s", dep.Name, state.Name, resp.Status)
	}

	return iname, nil

}
