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
	"github.com/extemporalgenome/slug"
	"github.com/hashicorp/errwrap"
	"github.com/samalba/dockerclient"

	"github.com/dockpit/pit/model"
)

var DockpitPrefix = "dp"
var ConnectionTimeout = time.Millisecond * 500

type Docker struct {
	Host string

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
	d.client, err = dockerclient.NewDockerClientTimeout(hurl.String(), &tlsc, ConnectionTimeout)
	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Failed to create Docker client for '%s': {{err}}", host), err)
	}

	return d, nil
}

func (d *Docker) Info() (*dockerclient.Info, error) {
	info, err := d.client.Info()
	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Failed to retrieve Docker host info for '%s': {{err}}", d.Host), err)
	}

	return info, nil
}

func (d *Docker) RemoveAll() error {
	cs, err := d.client.ListContainers(true, false, "")
	if err != nil {
		return err
	}

	for _, c := range cs {
		remove := false
		for _, n := range c.Names {
			if strings.HasPrefix(n[1:], DockpitPrefix) {
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

func (d *Docker) Containers() ([]dockerclient.Container, error) {
	all, err := d.client.ListContainers(true, false, "")
	if err != nil {
		return nil, err
	}

	res := []dockerclient.Container{}
	for _, c := range all {
		for _, n := range c.Names {
			if strings.HasPrefix(n[1:], DockpitPrefix) {
				res = append(res, c)
			}
		}
	}

	return res, nil
}

func (d *Docker) Switch(iso *model.Isolation) error {
	err := d.RemoveAll()
	if err != nil {
		return errwrap.Wrapf("Failed to remove all running dockpit containers before switching isolation: {{err}}", err)
	}

	//start all for the current isolation
	depst, err := d.model.GetIsolationDepsAndStates(iso)
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Failed to get deps and states for starting iso '%s': {{err}}", iso.Name), err)
	}

	for dep, state := range depst {
		run, err := model.NewRun(*state)
		if err != nil {
			return errwrap.Wrapf(fmt.Sprintf("Failed to create run from state '%s': {{err}}", state.Name), err)
		}

		cid, err := d.Start(run, fmt.Sprintf("%s.%s.", DockpitPrefix, iso.ID))
		if err != nil {
			return errwrap.Wrapf(fmt.Sprintf("Failed to start state '%s': {{err}}", state.Name), err)
		}

		//@todo probalby wanna store this somewhere
		_ = run //run instance
		_ = dep //dependency
		_ = cid //container id
	}

	return nil
}

func (d *Docker) Start(run *model.Run, prefix string) (string, error) {
	var err error

	//container config
	cconfig := run.State.Settings.ContainerConfig
	cconfig.Image = run.State.ImageName

	run.ContainerID, err = d.client.CreateContainer(cconfig, prefix+run.State.ImageName)
	if err != nil {
		return "", errwrap.Wrapf(fmt.Sprintf("Failed to create state container with image '%s': {{err}}, is the state build?", run.State.ImageName), err)
	}

	//hostconfig
	hconfig := run.State.Settings.HostConfig

	err = d.client.StartContainer(run.ContainerID, hconfig)
	if err != nil {
		return run.ContainerID, errwrap.Wrapf(fmt.Sprintf("Failed to start state container with image '%s': {{err}}", run.State.ImageName), err)
	}

	rc, err := d.client.ContainerLogs(run.ContainerID, &dockerclient.LogOptions{Follow: true, Stdout: true, Stderr: true})
	if err != nil {
		return run.ContainerID, errwrap.Wrapf(fmt.Sprintf("Failed to follow logs of state container '%s': {{err}}", run.ContainerID), err)
	}

	run.ContainerInfo, err = d.client.InspectContainer(run.ContainerID)
	if err != nil {
		return run.ContainerID, errwrap.Wrapf(fmt.Sprintf("Failed to get container info for %s: {{err}}", run.ContainerID), err)
	}

	tr := io.TeeReader(rc, run.Output.Buffer)

	// wait for ready pattern or error
	readypatt := regexp.Regexp(*run.State.Settings.ReadyPattern)
	readyto := time.Duration(run.State.Settings.ReadyTimeout)

	err = iowait.WaitForRegexp(tr, &readypatt, readyto)
	if err != nil {
		return run.ContainerID, errwrap.Wrapf(fmt.Sprintf("Failed to wait for state container %s: {{err}}", run.ContainerID), err)
	}

	run.IsReady = true

	return run.ContainerID, nil
}

func (d *Docker) Remove(s *model.State) error {
	cs, err := d.client.ListContainers(true, false, "")
	if err != nil {
		return err
	}

	//get container that matches the name
	var container dockerclient.Container
	for _, c := range cs {
		for _, n := range c.Names {
			if n[1:] == s.ImageName {
				container = c
			}
		}
	}

	if container.Id == "" {
		return nil
	}

	return d.client.RemoveContainer(container.Id, true)
}

func (d *Docker) Build(b *model.Build) (string, error) {
	dep := b.Dep
	state := b.State

	out := b.Output.Buffer
	in := bytes.NewBuffer(nil)
	tw := tar.NewWriter(in)

	for fname, f := range state.Files {
		//create writer for tar data
		f := bytes.NewReader([]byte(f.Content))

		hdr := &tar.Header{
			Name: fname,
			Size: int64(f.Len()),
		}

		err := tw.WriteHeader(hdr)
		if err != nil {
			return "", errwrap.Wrapf(fmt.Sprintf("Failed to write tar file '%s' header for '%s'::'%s': {{err}}", fname, dep.Name, state.Name), err)
		}

		if _, err = io.Copy(tw, f); err != nil {
			return "", errwrap.Wrapf(fmt.Sprintf("Failed to tar file '%s' for '%s'::'%s': {{err}}", fname, dep.Name, state.Name), err)
		}
	}

	// generate an unique image name based on the provider name and path to the state folder
	iname := fmt.Sprintf("%s-%s-%s", DockpitPrefix, dep.Name, state.Name)
	iname = slug.SlugAscii(iname)

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
