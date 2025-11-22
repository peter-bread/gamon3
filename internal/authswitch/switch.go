// Package authswitch provides functions to interact with the GitHub CLI to switch which
// account is active.
package authswitch

import "fmt"

// Switch calls `gh auth switch --user <account>`.
func Switch(runner Runner, account string) error {
	stderr, err := runner.Run("gh", "auth", "switch", "--user", account)
	if err != nil {
		return fmt.Errorf(
			"failed to run command:\n  gh auth switch --user %s\ngh stderr:\n  %s",
			account, stderr,
		)
	}
	return nil
}

// SwitchIfNeeded calls Switch only if the current account differs from the account that should
// be active.
func SwitchIfNeeded(runner Runner, account string, current string) error {
	if account != current {
		return Switch(runner, account)
	}
	return nil
}
