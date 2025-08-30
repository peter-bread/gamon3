// Package ghswitch provides functions to switch gh accounts.
package ghswitch

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/goccy/go-yaml"
)

// Config stores config data. This data is read from a YAML file.
type Config struct {
	Accounts map[string][]string `yaml:"accounts"`
	Default  string              `yaml:"default"`
}

// MapAccounts runs the function `f` on all acounts in a `Config`.
// This function modifies `Config` in-place.
func (c *Config) MapAccounts(f func(string) string) {
	for account, paths := range c.Accounts {
		for i, p := range paths {
			paths[i] = f(p)
		}
		c.Accounts[account] = paths
	}
}

// Load reads `Config` data from a YAML file located at `path`.
func (c *Config) Load(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(data, c); err != nil {
		return err
	}

	return nil
}

func GetConfigPath() (string, error) {
	// TODO: Handle .yml and .yaml extensions.
	if configDir, found := os.LookupEnv("GAMON3_CONFIG_DIR"); found {
		return filepath.Join(configDir, "config.yml"), nil
	}

	if xdgConfigDir, found := os.LookupEnv("XDG_CONFIG_HOME"); found {
		return filepath.Join(xdgConfigDir, "gamon3", "config.yml"), nil
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, ".config", "gamon3", "config.yml"), nil
}

func (c *Config) printYAML() {
	data, err := yaml.Marshal(c)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(data))
}

func (c *Config) printRaw() {
	fmt.Println(c)
}
