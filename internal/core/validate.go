package core

import (
	"fmt"
	"slices"

	"github.com/peter-bread/gamon3/internal/config"
)

func ValidateMainConfig(cfg *config.MainConfig, ghHosts []string) []error {
	var errs []error

	if cfg.Default == "" {
		errs = append(errs, fmt.Errorf("default account is required"))
	} else if !slices.Contains(ghHosts, cfg.Default) {
		errs = append(errs, fmt.Errorf("default account %q is not authenticated with GH CLI", cfg.Default))
	}

	for account := range cfg.Accounts {
		if !slices.Contains(ghHosts, account) {
			errs = append(errs, fmt.Errorf("account %q is not authenicated with GH CLI", account))
		}
	}

	return errs
}

func ValidateLocalConfig(cfg *config.LocalConfig, ghHosts []string) error {
	if cfg.Account == "" {
		return fmt.Errorf("default account is required")
	}

	if !slices.Contains(ghHosts, cfg.Account) {
		return fmt.Errorf("account %q is not authenicated with GH CLI", cfg.Account)
	}

	return nil
}
