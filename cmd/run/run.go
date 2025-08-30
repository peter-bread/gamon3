package run

import (
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
	mappingPath, err := ghswitch.GetMappingPath()
	if err != nil {
		log.Fatalln(err)
	}

	if err := mapping.Load(mappingPath); err != nil {
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

	currentAccount := ghHosts.GetCurrentUser()

	if account != currentAccount {
		exec.Command("gh", "auth", "switch", "--user", account).Run()
	}
}
