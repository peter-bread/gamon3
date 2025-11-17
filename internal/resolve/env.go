package resolve

import (
	"fmt"

	"github.com/peter-bread/gamon3/internal/data"
)

func DoEnv(account string, gh GhHosts) (Result, error) {
	if account == "" {
		return Result{}, fmt.Errorf("env account cannot be empty")
	}

	if !IsValidGitHubAccount(account, gh) {
		return Result{}, fmt.Errorf("env account %q is not authenticated", account)
	}

	return Result{
		Current:     gh.CurrentUser(),
		Account:     account,
		SourceKind:  Env,
		SourceValue: data.EnvVarAccount,
	}, nil
}
