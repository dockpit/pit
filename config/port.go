package config

import (
	"fmt"
	"strings"
)

type PortConfig struct {
	Host      string
	Container string
}

func ParsePort(pstr string) (*PortConfig, error) {
	parts := strings.SplitN(pstr, ":", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("Invalid port format: '%s'", pstr)
	}

	return &PortConfig{Host: parts[1], Container: parts[0]}, nil
}
