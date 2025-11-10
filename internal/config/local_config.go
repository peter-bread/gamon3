package config

import (
	"os"

	"github.com/goccy/go-yaml"
)

// LocalConfig is a structure that stores data from a local configuration file.
type LocalConfig struct {
	Account string `yaml:"account"`
}

// LoadLocalConfig attempts to read data from a YAML file at `path` and returns
// a new `LocalConfig` structure.
func LoadLocalConfig(path string) (*LocalConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg LocalConfig
	return &cfg, yaml.Unmarshal(data, &cfg)
}
