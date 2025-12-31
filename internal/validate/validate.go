/*
Copyright © 2025 Peter Sheehan

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

// Package validate provides functions to validate Gamon3 config files.
package validate

import (
	"fmt"
	"slices"

	"github.com/peter-bread/gamon3/v2/internal/config"
)

// ValidateMainConfig validates a main config and reports any errors.
func ValidateMainConfig(cfg *config.MainConfig, ghHosts []string) []error {
	var errs []error

	if cfg.Default == "" {
		errs = append(errs, fmt.Errorf("default account is required"))
	} else if !slices.Contains(ghHosts, cfg.Default) {
		errs = append(errs, fmt.Errorf("default account %q is not authenticated with GH CLI", cfg.Default))
	}

	for account := range cfg.Accounts {
		if !slices.Contains(ghHosts, account) {
			errs = append(errs, fmt.Errorf("account %q is not authenticated with GH CLI", account))
		}
	}

	return errs
}

// ValidateLocalConfig validates a local config and reports any errors.
func ValidateLocalConfig(cfg *config.LocalConfig, ghHosts []string) error {
	if cfg.Account == "" {
		return fmt.Errorf("default account is required")
	}

	if !slices.Contains(ghHosts, cfg.Account) {
		return fmt.Errorf("account %q is not authenticated with GH CLI", cfg.Account)
	}

	return nil
}
