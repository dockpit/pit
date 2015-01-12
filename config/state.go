package config

import (
	"regexp"
	"time"
)

type StateProviderConfig struct {
	name         string
	ports        []*PortConfig
	readyExp     *regexp.Regexp
	cmd          []string
	readyTimeout time.Duration
	defaultState string
}

func (s *StateProviderConfig) Ports() []*PortConfig        { return s.ports }
func (s *StateProviderConfig) ReadyExp() *regexp.Regexp    { return s.readyExp }
func (s *StateProviderConfig) ReadyTimeout() time.Duration { return s.readyTimeout }
func (s *StateProviderConfig) Cmd() []string               { return s.cmd }
func (s *StateProviderConfig) Name() string                { return s.name }
func (s *StateProviderConfig) DefaultState() string        { return s.defaultState }
