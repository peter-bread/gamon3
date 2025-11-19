package executor

import "os/exec"

func Switch(account string) error {
	return exec.Command("gh", "auth", "switch", "--user", account).Run()
}
