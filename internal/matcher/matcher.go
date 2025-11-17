package matcher

import (
	"fmt"
	"strings"
)

type FP interface {
	Abs(path string) (string, error)
}

// MatchAccount returns the account that applies to the given directory.
func MatchAccount(fp FP, dir string, accounts map[string][]string, defaultAccount string) (string, error) {
	abs, err := fp.Abs(dir)
	if err != nil {
		return "", err
	}

	for account, paths := range accounts {
		for _, path := range paths {
			norm := normalise(path)
			if strings.HasPrefix(abs, norm) {
				if account == "" {
					return "", fmt.Errorf("account key cannot be empty for path %q", norm)
				}
				return account, nil
			}
		}
	}

	if defaultAccount == "" {
		return "", fmt.Errorf("field 'default' cannot be empty")
	}

	return defaultAccount, nil
}
