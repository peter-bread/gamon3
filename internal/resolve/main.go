package resolve

import (
	"fmt"
	"path/filepath"

	"github.com/peter-bread/gamon3/internal/matcher"
)

type fp struct{}

func (fp) Abs(path string) (string, error) { return filepath.Abs(path) }

func resolveMain(path string, gh GhHosts, loader MainConfigLoader, os OS) (Result, error) {
	cfg, err := loader.Load(path)
	if err != nil {
		return Result{}, err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return Result{}, err
	}

	// This is the account that should be used based on the current directory.
	account, err := matcher.MatchAccount(fp{}, cwd, cfg.Accounts, cfg.Default)
	if err != nil {
		return Result{}, fmt.Errorf("main config %s", err)
	}

	if !isValidGitHubAccount(account, gh) {
		return Result{}, fmt.Errorf("main config account %q is not authenticated", account)
	}

	return Result{
		Current:     gh.CurrentUser(),
		Account:     account,
		SourceKind:  Main,
		SourceValue: path,
	}, nil
}
