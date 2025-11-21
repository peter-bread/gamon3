package executor

import "fmt"

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
