package locator

import (
	"fmt"
	"os"
	"path/filepath"
)

func MainConfigPath() (string, error) {
	dir, err := getMainConfigDir()
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
