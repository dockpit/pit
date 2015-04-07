package ui

import (
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
	Containers []dockerclient.Container
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
	s.State.CurrentIsolationName = iso.Name
	s.State.CurrentIsolationStatus = "Starting..."
	s.Syncs <- struct{}{}

	err := s.client.Switch(iso)
	if err != nil {
		s.State.CurrentIsolationStatus = "Error"
		s.State.Errors = append(s.State.Errors, err)
	} else {
		s.State.CurrentIsolationStatus = "OK"
	}

	s.Syncs <- struct{}{}
}

func (s *Store) Sync() {
	var err error

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

		//get dockpit containers
		s.State.Containers, err = s.client.Containers()
		if err != nil {
			s.State.Errors = append(s.State.Errors, err)
		}

		//@todo determine current isolation
		//from container names

	}

	s.Syncs <- struct{}{}
}
