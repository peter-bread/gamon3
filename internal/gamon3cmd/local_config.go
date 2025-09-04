package gamon3cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"github.com/goccy/go-yaml"
)

type LocalConfig struct {
	Account string `yaml:"account"`
}

func (l *LocalConfig) Load(path string, allowedUsers []string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(data, l); err != nil {
		return err
	}

	if err := l.Validate(allowedUsers); err != nil {
		return err
	}

	return nil
}

func (l *LocalConfig) Validate(allowedUsers []string) error {
	if err := l.ValidateUsers(allowedUsers); err != nil {
		return err
	}

	return nil
}

func (l *LocalConfig) ValidateUsers(allowedUsers []string) error {
	if l.Account == "" {
		return fmt.Errorf("%s", "local config: 'account' field is required")
	}

	if !slices.Contains(allowedUsers, l.Account) {
		return fmt.Errorf("%s: '%s' %s", "local config", l.Account, "has not been registered with GH CLI")
	}

	return nil
}

func GetLocalConfigPath() (string, error) {
	start, _ := os.Getwd()
	stop, _ := os.UserHomeDir()
	dir := start
	candidates := []string{".gamon.yaml", ".gamon.yml"}
	for {

		for _, name := range candidates {
			candidate := filepath.Join(dir, name)
			if _, err := os.Stat(candidate); err == nil {
				return candidate, nil
			}
		}

		parent := filepath.Dir(dir)

		if dir == stop || parent == dir {
			return "", fmt.Errorf("%s", "Could not find a local config file")
		}

		dir = parent
	}
}
