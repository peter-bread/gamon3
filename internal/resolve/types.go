package resolve

import "github.com/peter-bread/gamon3/internal/config"

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

type GhHosts interface {
	AllUsers() []string
	CurrentUser() string
}

type Locator interface {
	GhHostsPath() (string, error)
	EnvAccount() (string, bool)
	LocalConfigPath() (string, error)
	MainConfigPath() (string, error)
}

type Loader[T any] interface {
	Load(path string) (T, error)
}

type (
	GhHostsLoader     = Loader[GhHosts]
	LocalConfigLoader = Loader[*config.LocalConfig]
	MainConfigLoader  = Loader[*config.MainConfig]
)

type OS interface {
	Getwd() (string, error)
}
