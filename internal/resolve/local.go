package resolve

import "fmt"

func resolveLocal(path string, gh GhHosts, loader LocalConfigLoader) (Result, error) {
	cfg, err := loader.Load(path)
	if err != nil {
		return Result{}, err
	}

	account := cfg.Account

	if account == "" {
		return Result{}, fmt.Errorf("local config field 'account' cannot be empty")
	}

	if !isValidGitHubAccount(account, gh) {
		return Result{}, fmt.Errorf("local config account %q is not authenticated", account)
	}

	return Result{
		Current:     gh.CurrentUser(),
		Account:     account,
		SourceKind:  Local,
		SourceValue: path,
	}, nil
}
