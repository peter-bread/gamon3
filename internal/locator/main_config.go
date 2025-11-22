// Package locator provides functions and interfaces to locate configuration
// sources and GitHub CLI config files.
package locator

import (
	"fmt"
	"os"
	"path/filepath"
)

type MainOS interface {
	Stat(name string) (os.FileInfo, error)
	UserConfigDir() (string, error)
	LookupEnv(key string) (string, bool)
}

// MainConfigPath returns the path to a main config file for Gamon3. First it establishes
// which directory it should look inside. The first of the following to exist is used:
//   - $GAMON3_CONFIG_DIR
//   - $XDG_CONFIG_HOME/gamon3
//   - $HOME/.config/gamon3
//
// If a directory is found, it will look for a file called either config.yaml or config.yml.
// If one of these files exists, then the path to it is returned.
//
// If a file does not exist, the function will return an error. This also means that if
// $GAMON3_CONFIG_DIR is set to /foo/bar and does not contain a config.yaml, but an actual
// config file exists in $HOME/.config/gamon3, it will not be found.
func MainConfigPath(os MainOS) (string, error) {
	dir, err := getMainConfigDir(os)
	if err != nil {
		return "", err
	}

	candidates := []string{"config.yaml", "config.yml"}

	for _, name := range candidates {
		candidate := filepath.Join(dir, name)
		if _, err := os.Stat(candidate); err == nil {
			return candidate, nil
		}
	}

	return "", fmt.Errorf("could not find a main config file")
}

func getMainConfigDir(os MainOS) (string, error) {
	if dir, found := os.LookupEnv("GAMON3_CONFIG_DIR"); found {
		return dir, nil
	}

	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, "gamon3"), nil
}
