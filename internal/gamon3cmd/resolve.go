package gamon3cmd

import (
	"fmt"
	"os"
	"slices"
)

func Resolve() (currentAccount string, account string, source string, err error) {
	var ghHosts GHHosts

	ghHostsPath, err := GetGHHostsPath()
	if err != nil {
		// TODO: Should I still return the original error?
		return "", "", "", fmt.Errorf("%s", "[GAMON3 ERROR] Failed to get path for GH CLI 'hosts.yml'")
	}

	if err := ghHosts.Load(ghHostsPath); err != nil {
		// TODO: Should I still return the original error?
		return "", "", "", fmt.Errorf("%s", "[GAMON3 ERROR] Failed to load GH CLI 'hosts.yml'")
	}

	currentAccount = ghHosts.GetCurrentUser()
	users := ghHosts.GetAllUsers()

	// Check $GAMON3_ACCOUNT.
	// IMPORTANT: $GAMON3_ACCOUNT *must* be exported.

	if account, found := os.LookupEnv("GAMON3_ACCOUNT"); found {
		if slices.Contains(users, account) {
			return currentAccount, account, "environment variable", nil
		} else {
			fmt.Printf("[GAMON3 WARN] '%s' is not a valid account\n", account)
			fmt.Println("[GAMON3 INFO] Falling back to local config file")
		}
	}

	// Walk up filetree looking for a local '.gamon3.yaml' file.
	// It should also stop walking at the $HOME directory, at which point it
	// falls back to 'config.yaml'.

	var localConfig LocalConfig

	if localConfigPath, err := GetLocalConfigPath(); err == nil {
		if err := localConfig.Load(localConfigPath, users); err != nil {
			fmt.Printf("[GAMON3 WARN] '%s' is not a valid local config file\n", localConfigPath)
			fmt.Println("[GAMON3 INFO] Falling back to main config file")
		} else {
			return currentAccount, localConfig.Account, "local config", nil
		}
	}

	// Check main 'config.yaml' file.

	var config Config

	configPath, err := GetConfigPath()
	if err != nil {
		return currentAccount, "", "", nil
	}

	if err := config.Load(configPath, users); err != nil {
		errorMsg := fmt.Sprintf("[GAMON3 ERROR] '%s' is not a valid local config file\n", configPath)
		return currentAccount, "", "", fmt.Errorf("%s", errorMsg)
	}

	pwd := os.Getenv("PWD")
	account = config.GetAccount(pwd)
	return currentAccount, account, "main config", nil
}
