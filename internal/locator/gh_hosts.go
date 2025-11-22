package locator

import "path/filepath"

type GhOS interface {
	LookupEnv(key string) (string, bool)
	UserConfigDir() (string, error)
}

// GhHostsPath returns the path to the GitHub CLI hosts.yml file if it exists,
// otherwise it returns an error.
func GhHostsPath(os GhOS) (string, error) {
	dir, err := getGhConfigDir(os)
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, "hosts.yml"), nil
}

func getGhConfigDir(os GhOS) (string, error) {
	if dir, found := os.LookupEnv("GH_CONFIG_DIR"); found {
		return dir, nil
	}

	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, "gh"), nil
}
