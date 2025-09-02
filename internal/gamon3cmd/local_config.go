package gamon3cmd

import (
	"os"

	"github.com/goccy/go-yaml"
)

type LocalConfig struct {
	Account string `yaml:"account"`
}

func (l *LocalConfig) Load(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(data, l); err != nil {
		return err
	}

	// TODO: Validate?

	return nil
}
