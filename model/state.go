package model

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"code.google.com/p/go-uuid/uuid"
	"github.com/samalba/dockerclient"
)

type RegexpPattern regexp.Regexp

func (r RegexpPattern) MarshalJSON() ([]byte, error) {
	exp := regexp.Regexp(r)

	return []byte(strconv.Quote(exp.String())), nil
}

func (r *RegexpPattern) UnmarshalJSON(data []byte) error {
	patt, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}

	exp, err := regexp.Compile(patt)
	if err != nil {
		return err
	}

	*r = RegexpPattern(*exp)
	return nil
}

type Duration time.Duration

func (d *Duration) UnmarshalJSON(data []byte) error {
	s, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}

	nd, err := time.ParseDuration(s)
	if err != nil {
		return err
	}

	*d = Duration(nd)
	return nil
}

func (d Duration) MarshalJSON() ([]byte, error) {
	duration := time.Duration(d)

	return []byte(strconv.Quote(duration.String())), nil
}

type StateSettings struct {
	ReadyTimeout    Duration                      `json:"ready_timeout"`
	ReadyPattern    *RegexpPattern                `json:"ready_pattern"`
	HostConfig      *dockerclient.HostConfig      `json:"host_config"`
	ContainerConfig *dockerclient.ContainerConfig `json:"container_config"`
}

type State struct {
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	Files     map[string]string `json:"files"`
	Settings  *StateSettings    `json:"settings"`
	ImageName string            `json:"image_name"`
}

func NewStateFromTemplate(name string, t *Template) (*State, error) {
	if name == "" || t == nil {
		return nil, fmt.Errorf("State Validation Error: Name and Type cannot be empty")
	}

	files := map[string]string{}
	for name, tf := range t.TemplateFiles {
		files[name] = tf.Content
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
