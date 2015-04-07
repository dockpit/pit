package ui

import (
	"net"

	"github.com/dockpit/pit/client"
	"github.com/dockpit/pit/model"

	"github.com/samalba/dockerclient"
)

type State struct {
	DockerHostAddress string // docker host address
	DockerHostStatus  string // docker host status
	DockerHostFixture string // running fixture on docker

	Isolations []*model.Isolation
	DockerInfo *dockerclient.Info
	Containers []dockerclient.Container
	Errors     []error
}

func NewState() *State {
	return &State{
		DockerHostAddress: "Unkown",
		DockerHostStatus:  "Unkown",
		DockerHostFixture: "None",
		Errors:            []error{},
	}
}

type Store struct {
	model  *model.Model
	client *client.Docker

	State *State
}

func NewStore(m *model.Model, c *client.Docker) (*Store, error) {
	s := &Store{m, c, NewState()}

	go func() {
		for range s.model.Events {
			s.Sync()
		}
	}()

	return s, nil
}

func (s *Store) Sync() {
	var err error

	//reset state, we're gonna repopulate it
	s.State = NewState()

	//get host info
	s.State.DockerInfo, err = s.client.Info()
	if err != nil {
		if err, ok := err.(net.Error); ok && err.Timeout() {
			//timeout is it running?
		}

		s.State.Errors = append(s.State.Errors, err)
	}

	//get dockpit containers
	s.State.Containers, err = s.client.Containers()
	if err != nil {
		s.State.Errors = append(s.State.Errors, err)
	}

	//get isolations
	s.State.Isolations, err = s.model.GetAllIsolations()
	if err != nil {
		s.State.Errors = append(s.State.Errors, err)
	}

	//@todo determine current isolation
	//from container names

}
