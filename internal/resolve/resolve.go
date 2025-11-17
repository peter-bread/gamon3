package resolve

import (
	"fmt"
)

// Resolve decides which method should be used to determine current account and delegates to other functions.
// It returns information about the account, and any errors that occur.
func Resolve(l Locator, ghLoader GhHostsLoader, localLoader LocalConfigLoader, mainLoader MainConfigLoader, os OS) (Result, error) {
	ghHostsPath, err := l.GhHostsPath()
	if err != nil {
		return Result{}, err
	}

	ghHosts, err := ghLoader.Load(ghHostsPath)
	if err != nil {
		return Result{}, err
	}

	if envAccount, found := l.EnvAccount(); found {
		return resolveEnv(envAccount, ghHosts)
	}

	if localPath, err := l.LocalConfigPath(); err == nil {
		return resolveLocal(localPath, ghHosts, localLoader)
	}

	if mainPath, err := l.MainConfigPath(); err == nil {
		return resolveMain(mainPath, ghHosts, mainLoader, os)
	}

	return Result{}, fmt.Errorf("no config sources found to be able to resolve account")
}
