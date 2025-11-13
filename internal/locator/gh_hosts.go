package locator

import (
	"os"
	"path/filepath"
)

func GhHostsPath() (string, error) {
	dir, err := getGhConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, "hosts.yml"), nil
}

func getGhConfigDir() (string, error) {
	if dir, found := os.LookupEnv("GH_CONFIG_DIR"); found {
		return dir, nil
	}

	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, "gh"), nil
}
