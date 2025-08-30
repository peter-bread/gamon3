package run

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"peter-bread/gamon3/internal/ghswitch"
)

// Run reads from a JSON mapping file and determines the account based on
// current working directory.
func Run() {
	var readMap ghswitch.Mapping
	readPath := "examples/mapping.json"
	if err := readMap.Load(readPath); err != nil {
		log.Fatalln(err)
	}

	pwd := os.Getenv("PWD")
	account := readMap.GetAccount(pwd)

	var ghHosts ghswitch.GHHosts
	ghHostsPath, err := ghswitch.GetGHConfigPath()
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
