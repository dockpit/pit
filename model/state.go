package model

import (
	"fmt"
	"regexp"
	"time"

	"code.google.com/p/go-uuid/uuid"
	"github.com/samalba/dockerclient"
)

type StateSettings struct {
	ReadyTimeout    Duration                      `json:"ready_timeout"`
	ReadyPattern    *RegexpPattern                `json:"ready_pattern"`
	HostConfig      *dockerclient.HostConfig      `json:"host_config"`
	ContainerConfig *dockerclient.ContainerConfig `json:"container_config"`
}

type State struct {
	ID        string          `json:"id"`
	Name      string          `json:"name"`
	Files     map[string]File `json:"files"`
	Settings  *StateSettings  `json:"settings"`
	ImageName string          `json:"image_name"`
}

func NewStateFromTemplate(name string, t *Template) (*State, error) {
	if name == "" || t == nil {
		return nil, fmt.Errorf("State Validation Error: Name and Type cannot be empty")
	}

	files := map[string]File{}
	for name, f := range t.Files {
		files[name] = *f
	}

	st := &State{
		ID:       uuid.New(),
		Name:     name,
		Files:    files,
		Settings: t.StateSettings,
	}

	//default settings
	if st.Settings == nil {
		d, _ := time.ParseDuration("38s")
		patt := RegexpPattern(*regexp.MustCompile(`.*`))

		st.Settings = &StateSettings{
			ReadyTimeout: Duration(d),
			ReadyPattern: &patt,
		}
	}

	//default host config
	if st.Settings.HostConfig == nil {
		st.Settings.HostConfig = &dockerclient.HostConfig{}
	}

	if st.Settings.ContainerConfig == nil {
		st.Settings.ContainerConfig = &dockerclient.ContainerConfig{}
	}

	return st, nil
}
