package locator

import (
	"fmt"
	"os"
	"path/filepath"
)

type LocalOS interface {
	Stat(name string) (os.FileInfo, error)
	Getwd() (dir string, err error)
	UserHomeDir() (string, error)
}

// LocalConfigPath searches upwards from the current directory, looking for a local config file.
// This file can be any one of: .gamon.yaml, .gamon.yml, .gamon3.yaml, or .gamon3.yml.
// In most cases, the search will stop after checking the user's home directory. If the directory
// where the search is starting is not a descendant of the home directory, then the search will go
// all the way to the filesystem root.
func LocalConfigPath(os LocalOS) (string, error) {
	start, err := os.Getwd()
	if err != nil {
		return "", err
	}

	stop, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	candidates := []string{".gamon.yaml", ".gamon.yml", ".gamon3.yaml", ".gamon3.yml"}

	return searchUpward(os, start, stop, candidates)
}

func searchUpward(os LocalOS, start string, stop string, candidates []string) (string, error) {
	dir := start

	for {
		for _, name := range candidates {
			candidate := filepath.Join(dir, name)
			if _, err := os.Stat(candidate); err == nil {
				return candidate, nil
			}
		}

		parent := filepath.Dir(dir)
		if dir == stop || parent == dir {
			break
		}
		dir = parent
	}

	return "", fmt.Errorf("could not find any of %v starting from %q", candidates, start)
}
