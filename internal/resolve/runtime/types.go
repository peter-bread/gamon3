package runtime

// Define implementations for interfaces.

import (
	"os"

	"github.com/peter-bread/gamon3/internal/config"
	"github.com/peter-bread/gamon3/internal/locator"
	"github.com/peter-bread/gamon3/internal/resolve"
)

type Locator struct{}

func (Locator) GhHostsPath() (string, error)     { return locator.GhHostsPath() }
func (Locator) EnvAccount() (string, bool)       { return locator.EnvAccount() }
func (Locator) LocalConfigPath() (string, error) { return locator.LocalConfigPath(OS{}) }
func (Locator) MainConfigPath() (string, error)  { return locator.MainConfigPath(OS{}) }

type OS struct{}

func (OS) Getwd() (string, error)                { return os.Getwd() }
func (OS) Stat(name string) (os.FileInfo, error) { return os.Stat(name) }
func (OS) UserHomeDir() (string, error)          { return os.UserHomeDir() }
func (OS) UserConfigDir() (string, error)        { return os.UserConfigDir() }
func (OS) LookupEnv(key string) (string, bool)   { return os.LookupEnv(key) }

type Loader[T any] struct {
	loadFunc func(string) (T, error)
}

func (r Loader[T]) Load(path string) (T, error) {
	return r.loadFunc(path)
}

type (
	GhHostsLoader     = Loader[resolve.GhHosts]
	LocalConfigLoader = Loader[*config.LocalConfig]
	MainConfigLoader  = Loader[*config.MainConfig]
)

type Resolver struct {
	locator     Locator
	ghLoader    GhHostsLoader
	localLoader LocalConfigLoader
	mainLoader  MainConfigLoader
	os          OS
}
