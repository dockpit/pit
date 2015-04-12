package ui

import (
	"strings"

	"github.com/dockpit/pit/client"
	"github.com/dockpit/pit/model"

	"github.com/samalba/dockerclient"
)

type State struct {
	DockerHostAddress      string // docker host address
	DockerHostStatus       string // docker host status
	CurrentIsolationStatus string // current isolation
	CurrentIsolationName   string //name of the current isolation

	Deps       []*model.Dep
	Isolations []*model.Isolation
	DockerInfo *dockerclient.Info
	Errors     []error
}

func NewState() *State {
	return &State{
		DockerHostAddress:      "Unkown",
		DockerHostStatus:       "Unkown",
		CurrentIsolationStatus: "",
		CurrentIsolationName:   "<none>",
		Errors:                 []error{},
	}
}

type Store struct {
	model  *model.Model
	client *client.Docker
	events chan model.Event

	State *State
	Syncs chan struct{}
}

func NewStore(m *model.Model, c *client.Docker) (*Store, error) {
	s := &Store{
		model:  m,
		client: c,
		events: make(chan model.Event),

		State: NewState(),
		Syncs: make(chan struct{}),
	}

	m.Subscribe(s.events)
	go func() {
		for range s.events {
			s.Sync()
		}
	}()

	return s, nil
}

func (s *Store) SwitchTo(iso *model.Isolation) {

	//nil means we got the remove all command
	if iso == nil {
		err := s.client.RemoveAll()
		if err != nil {
			s.State.Errors = append(s.State.Errors, err)
		}

		s.Sync()
		return
	}

	s.State.CurrentIsolationName = iso.Name
	s.State.CurrentIsolationStatus = "Starting..."
	s.Syncs <- struct{}{}

	err := s.client.Switch(iso)
	if err != nil {
		s.State.CurrentIsolationStatus = "Error"
		s.State.Errors = append(s.State.Errors, err)
		return
	}

	s.Sync()
}

func (s *Store) Sync() {
	var err error
	s.State.Errors = []error{}

	//get isolations
	s.State.Isolations, err = s.model.GetAllIsolations()
	if err != nil {
		s.State.Errors = append(s.State.Errors, err)
	}

	//get deps
	s.State.Deps, err = s.model.GetAllDeps()
	if err != nil {
		s.State.Errors = append(s.State.Errors, err)
	}

	//get host info
	s.State.DockerInfo, err = s.client.Info()
	s.State.DockerHostAddress = s.client.Host
	if err != nil {
		//@todo retry on error?
		s.State.DockerHostStatus = "Error"
		s.State.Errors = append(s.State.Errors, err)
	} else {
		s.State.DockerHostStatus = "OK"

		//filter all isolations to the instances that are running
		running, err := s.client.Running(s.State.Isolations, s.State.Deps)
		if err != nil {
			s.State.Errors = append(s.State.Errors, err)
		}

		if len(running) > 0 {
			names := []string{}
			for iso, _ := range running {
				names = append(names, iso.Name)
			}

			s.State.CurrentIsolationName = strings.Join(names, ",")

			//@todo determine status though something more sophisticated
			s.State.CurrentIsolationStatus = "OK"
		} else {
			s.State.CurrentIsolationName = "<none>"
			s.State.CurrentIsolationStatus = ""
		}

	}

	s.Syncs <- struct{}{}
}
