package run

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"peter-bread/gamon3/internal/ghswitch"
)

var (
	mapping ghswitch.Mapping
	ghHosts ghswitch.GHHosts
)

// RunCmd reads from a JSON mapping file and determines the account based on
// current working directory.
func RunCmd() {
	// TODO: Read from proper location (XDG_STATE).
	readPath := "examples/mapping.json"
	if err := mapping.Load(readPath); err != nil {
		log.Fatalln(err)
	}

	pwd := os.Getenv("PWD")
	account := mapping.GetAccount(pwd)

	ghHostsPath, err := ghswitch.GetGHHostsPath()
	if err != nil {
		log.Fatalln(err)
	}

	if err := ghHosts.Load(ghHostsPath); err != nil {
		log.Fatalln(err)
	}

	currentAccount := ghHosts.GitHubCom.User
	fmt.Println("Current: ", currentAccount)

	if account != currentAccount {
		// cmd := exec.Command("gh", "auth", "switch", "--user", account).Run()
		cmd := exec.Command("echo", "gh", "auth", "switch", "--user", account)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}
}
