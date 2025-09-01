package gamon3cmd

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

type Mapping struct {
	Paths   map[string]string `json:"paths"`
	Default string            `json:"default"`
}

func normalise(path string) string {
	// TODO: Use `filepath.Abs`?
	return filepath.Clean(os.ExpandEnv(path))
}

func (m *Mapping) Create(c *Config) {
	m.Default = c.Default
	m.Paths = make(map[string]string)

	for account, paths := range c.Accounts {
		for _, path := range paths {
			m.Paths[normalise(path)] = account
		}
	}
}

func (m *Mapping) Save(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(m)
}

func (m *Mapping) Load(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, m); err != nil {
		return err
	}

	return nil
}

func (m *Mapping) GetAccount(dir string) string {
	abs, err := filepath.Abs(dir)
	if err != nil {
		return m.Default
	}

	if account, ok := m.Paths[abs]; ok {
		return account
	}

	for path, account := range m.Paths {
		if strings.HasPrefix(abs, path) {
			return account
		}
	}

	return m.Default
}

func GetMappingPath() (string, error) {
	if xdgStateDir, found := os.LookupEnv("XDG_STATE_HOME"); found {
		return filepath.Join(xdgStateDir, "gamon3", "mapping.json"), nil
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, ".local", "state", "gamon3", "mapping.json"), nil
}
