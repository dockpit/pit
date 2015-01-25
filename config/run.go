package config

import (
	"regexp"
	"time"
)

type RunConfig struct {
	ReadyExp     *regexp.Regexp
	Cmd          []string
	ReadyTimeout time.Duration
	Dir          string
}

func ParseRunConfig(data *RunData) (*RunConfig, error) {
	var err error
	conf := &RunConfig{}

	if data.ReadyPattern == "" {
		data.ReadyPattern = ".*"
	}

	if data.ReadyTimeout == "" {
		data.ReadyTimeout = "2s"
	}

	conf.ReadyTimeout, err = time.ParseDuration(data.ReadyTimeout)
	if err != nil {
		return nil, err
	}

	conf.ReadyExp, err = regexp.Compile(data.ReadyPattern)
	if err != nil {
		return nil, err
	}

	conf.Cmd = data.Command
	conf.Dir = data.Dir

	return conf, nil
}
