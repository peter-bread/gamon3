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
