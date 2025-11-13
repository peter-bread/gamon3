package locator

import (
	"os"
	"path/filepath"
)

func MainConfigPath() (string, error) {
	dir, err := getMainConfigDir()
	if err != nil {
		return "", err
	}

	// TODO: Also check for config.yml
	return filepath.Join(dir, "config.yaml"), nil
}

func getMainConfigDir() (string, error) {
	if dir, found := os.LookupEnv("GAMON3_CONFIG_DIR"); found {
		return dir, nil
	}

	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, "gamon3"), nil
}
