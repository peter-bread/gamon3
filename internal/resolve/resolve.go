package resolve

import (
	"fmt"
	"os"
	"slices"

	"github.com/peter-bread/gamon3/internal/config"
	"github.com/peter-bread/gamon3/internal/data"
	"github.com/peter-bread/gamon3/internal/locator"
	"github.com/peter-bread/gamon3/internal/matcher"
)

type SourceKind string

const (
	Env   SourceKind = "env"
	Local SourceKind = "local"
	Main  SourceKind = "main"
)

type Result struct {
	Current     string
	Account     string
	SourceKind  SourceKind
	SourceValue string
}

func IsValidGitHubAccount(account string, gh *config.GhHosts) bool {
	return slices.Contains(gh.AllUsers(), account)
}

func DoEnv(account string, gh *config.GhHosts) (Result, error) {
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

func DoLocal(path string, gh *config.GhHosts) (Result, error) {
	cfg, err := config.LoadLocalConfig(path)
	if err != nil {
		return Result{}, err
	}

	account := cfg.Account

	if account == "" {
		return Result{}, fmt.Errorf("local config field 'account' cannot be empty")
	}

	if !IsValidGitHubAccount(account, gh) {
		return Result{}, fmt.Errorf("local config account %q is not authenticated", account)
	}

	return Result{
		Current:     gh.CurrentUser(),
		Account:     account,
		SourceKind:  Local,
		SourceValue: path,
	}, nil
}

func DoMain(path string, gh *config.GhHosts) (Result, error) {
	cfg, err := config.LoadMainConfig(path)
	if err != nil {
		return Result{}, err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return Result{}, err
	}

	// This is the account that should be used based on the current directory.
	account, err := matcher.MatchAccount(cwd, cfg.Accounts, cfg.Default)
	if err != nil {
		return Result{}, fmt.Errorf("main config %s", err)
	}

	if !IsValidGitHubAccount(account, gh) {
		return Result{}, fmt.Errorf("main config account %q is not authenticated", account)
	}

	return Result{
		Current:     gh.CurrentUser(),
		Account:     account,
		SourceKind:  Main,
		SourceValue: path,
	}, nil
}

// Resolve decides which method should be used to determine current account and delegates to other functions.
// It returns information about the account, and any errors that occur.
func Resolve() (Result, error) {
	ghHostsPath, err := locator.GhHostsPath()
	if err != nil {
		return Result{}, err
	}

	ghHosts, err := config.LoadGhHosts(ghHostsPath)
	if err != nil {
		return Result{}, err
	}

	if envAccount, found := locator.EnvAccount(); found {
		return DoEnv(envAccount, ghHosts)
	}

	if localPath, err := locator.LocalConfigPath(); err == nil {
		return DoLocal(localPath, ghHosts)
	}

	if mainPath, err := locator.MainConfigPath(); err == nil {
		return DoMain(mainPath, ghHosts)
	}

	return Result{}, fmt.Errorf("no method found to resolve account")
}
