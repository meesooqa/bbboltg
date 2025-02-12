package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// Conf for channels config yml
type Conf struct {
	Channels       map[string]ConfChannel `yaml:"channels"`
	UpdateInterval time.Duration          `yaml:"update"`
}

type ConfChannel struct {
	From    string             `yaml:"from"`
	To      string             `yaml:"to"`
	Filters ConfChannelFilters `yaml:"filters,omitempty"`
}

type ConfChannelFilters struct {
	Bolt      *ConfChannelFilter `yaml:"bolt,omitempty"`
	Important *ConfChannelFilter `yaml:"important,omitempty"`
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
