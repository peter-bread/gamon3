package gamon3cmd

import (
	"os"
	"path/filepath"

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

func GetLocalConfigPath() string {
	start, _ := os.Getwd()
	stop, _ := os.UserHomeDir()
	dir := start
	candidates := []string{".gamon.yaml", ".gamon.yml"}
	for {

		for _, name := range candidates {
			candidate := filepath.Join(dir, name)
			if _, err := os.Stat(candidate); err == nil {
				return candidate
			}
		}

		parent := filepath.Dir(dir)

		if dir == stop || parent == dir {
			return ""
		}

		dir = parent
	}
}
