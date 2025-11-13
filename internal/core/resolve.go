package core

import (
	"fmt"
	"os"
	"slices"

	"github.com/peter-bread/gamon3/internal/config"
	"github.com/peter-bread/gamon3/internal/locator"
)

type SourceKind string

const (
	Env   SourceKind = "env"
	Local SourceKind = "local"
	Main  SourceKind = "main"
)

type ResolveInput struct {
	EnvAccount      string
	EnvSourceValue  string
	LocalConfig     *config.LocalConfig
	LocalConfigPath string
	MainConfig      *config.MainConfig
	MainConfigPath  string
	GhUsers         []string
	CurrentGhUser   string
	Pwd             string
}

type Result struct {
	Current     string
	Account     string
	SourceKind  SourceKind
	SourceValue string
	Warnings    []string
	Err         error
}

func ResolvePure(in ResolveInput) Result {
	if in.EnvAccount != "" {
		if slices.Contains(in.GhUsers, in.EnvAccount) {
			return Result{
				Current:     in.CurrentGhUser,
				Account:     in.EnvAccount,
				SourceKind:  Env,
				SourceValue: "GAMON3_ACCOUNT",
			}
		}
		return Result{
			Err: fmt.Errorf("env account %q is not authenticated", in.EnvAccount),
		}
	}

	if in.LocalConfig != nil && in.LocalConfig.Account != "" {
		if slices.Contains(in.GhUsers, in.LocalConfig.Account) {
			return Result{
				Current:     in.CurrentGhUser,
				Account:     in.LocalConfig.Account,
				SourceKind:  Local,
				SourceValue: in.LocalConfigPath,
			}
		}
		return Result{
			Err: fmt.Errorf("local account %q is not authenticated", in.LocalConfig.Account),
		}
	}

	if in.MainConfig == nil {
		return Result{
			Err: fmt.Errorf("main config missing"),
		}
	}

	account := MatchAccount(in.Pwd, in.MainConfig.Accounts, in.MainConfig.Default)
	if !slices.Contains(in.GhUsers, account) {
		return Result{
			Err: fmt.Errorf("account %q is not authenticated", account),
		}
	}

	return Result{
		Current:     in.CurrentGhUser,
		Account:     account,
		SourceKind:  Main,
		SourceValue: in.MainConfigPath,
	}
}

func Resolve() Result {
	ghHostsPath, _ := locator.GhHostsPath()
	ghHosts, _ := config.LoadGhHosts(ghHostsPath)

	localPath, _ := locator.LocalConfigPath()
	localConfig, err := config.LoadLocalConfig(localPath)
	if err != nil {
		fmt.Println(err)
	}

	mainPath, _ := locator.MainConfigPath()
	mainConfig, _ := config.LoadMainConfig(mainPath)

	envAccount, found := locator.EnvAccount()
	envVar := "GAMON3_ACCOUNT"

	if !found {
		envAccount = ""
		envVar = ""
	}

	return ResolvePure(ResolveInput{
		EnvAccount:      envAccount,
		EnvSourceValue:  envVar,
		LocalConfig:     localConfig,
		LocalConfigPath: localPath,
		MainConfig:      mainConfig,
		MainConfigPath:  mainPath,
		GhUsers:         ghHosts.AllUsers(),
		CurrentGhUser:   ghHosts.CurrentUser(),
		Pwd:             os.Getenv("PWD"),
	})
}
