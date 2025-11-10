// Package config defines functions that read YAML files into corresponding
// structures.
package config

import (
	"os"

	"github.com/goccy/go-yaml"
)

// MainConfig is a structure that stores data from the main configuration file.
type MainConfig struct {
	Default  string              `yaml:"default"`
	Accounts map[string][]string `yaml:"accounts"`
}

// LoadMainConfig attempts to read data from a YAML file at `path` and returns
// a new `MainConfig` structure.
func LoadMainConfig(path string) (*MainConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg MainConfig
	return &cfg, yaml.Unmarshal(data, &cfg)
}
