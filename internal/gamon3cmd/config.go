// Package gamon3cmd provides functions to switch gh accounts.
package gamon3cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/goccy/go-yaml"
)

// Config stores config data. This data is read from a YAML file.
type Config struct {
	Accounts map[string][]string `yaml:"accounts"`
	Default  string              `yaml:"default"`
}

// Load reads `Config` data from a YAML file located at `path`. It then
// validates this data by comparing it to a list of `users`. This list
// should be obtained from a `GHHosts` structure.
func (c *Config) Load(path string, allowedUsers []string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(data, c); err != nil {
		return err
	}

	if err := c.Validate(allowedUsers); err != nil {
		return err
	}

	return nil
}

func (c *Config) Validate(allowedUsers []string) error {
	if err := c.ValidateUsers(allowedUsers); err != nil {
		return err
	}

	// TODO: Validate paths?

	return nil
}

func (c *Config) ValidateUsers(allowedUsers []string) error {
	var errs []string

	if c.Default == "" {
		errs = append(errs, "config: 'default' field is required")
	}

	if !slices.Contains(allowedUsers, c.Default) {
		errs = append(errs, "config: '"+c.Default+"' has not been registered with GH CLI")
	}

	if len(c.Accounts) == 0 {
		errs = append(errs, "config: 'accounts' section is either empty or missing")
	}

	for account := range c.Accounts {
		if account == "" {
			errs = append(errs, "config: account names must be non-empty")
		}
		if !slices.Contains(allowedUsers, account) {
			errs = append(errs, "config: '"+account+"' has not been registered with GH CLI")
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("%s", strings.Join(errs, "\n"))
	}

	return nil
}

func GetConfigDir() (string, error) {
	if configDir, found := os.LookupEnv("GAMON3_CONFIG_DIR"); found {
		return filepath.Join(configDir), nil
	}

	if xdgConfigDir, found := os.LookupEnv("XDG_CONFIG_HOME"); found {
		return filepath.Join(xdgConfigDir, "gamon3"), nil
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, ".config", "gamon3"), nil
}

func GetConfigPath() (string, error) {
	configDir, err := GetConfigDir()
	if err != nil {
		log.Fatalln(err)
	}

	if err := os.MkdirAll(configDir, 0755); err != nil {
		log.Fatalln(err)
	}

	candidates := []string{"config.yaml", "config.yml"}

	for _, name := range candidates {
		candidate := filepath.Join(configDir, name)
		if _, err := os.Stat(candidate); err == nil {
			return candidate, nil
		}
	}

	return "", fmt.Errorf("config: Config file does not exist")
}

func normalise(path string) string {
	return filepath.Clean(os.ExpandEnv(path))
}

func (c *Config) GetAccount(dir string) string {
	abs, err := filepath.Abs(dir)
	if err != nil {
		return c.Default
	}

	for account, paths := range c.Accounts {
		for _, path := range paths {
			norm := normalise(path)
			if strings.HasPrefix(abs, norm) {
				return account
			}
		}
	}

	return c.Default
}
