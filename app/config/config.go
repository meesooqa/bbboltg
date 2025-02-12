package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// Conf Channel Filter Types - subkeys of yaml-element "filters"
const (
	ccFilterTypeBolt      = "bolt"
	ccFilterTypeImportant = "important"
)

// Conf for channels config yml
type Conf struct {
	Channels                        map[string]ConfChannel `yaml:"channels"`
	UpdateInterval                  time.Duration          `yaml:"update"`
	availableConfChannelFilterTypes []string
}

type ConfChannel struct {
	Filters map[string]*ConfChannelFilter `yaml:"filters,omitempty"`
	From    string                        `yaml:"from"`
	To      string                        `yaml:"to,omitempty"`
}

type ConfChannelFilter struct {
	Prefix string `yaml:"prefix,omitempty"`
}

// Load config from file
func Load(fname string) (res *Conf, err error) {
	res = &Conf{}
	data, err := os.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}
