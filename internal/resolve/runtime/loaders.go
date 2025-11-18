package runtime

import (
	"github.com/peter-bread/gamon3/internal/config"
	"github.com/peter-bread/gamon3/internal/resolve"
)

func NewGhHostsLoader(f func(string) (*config.GhHosts, error)) GhHostsLoader {
	return GhHostsLoader{
		loadFunc: func(path string) (resolve.GhHosts, error) {
			return f(path)
		},
	}
}

func NewLocalConfigLoader(f func(string) (*config.LocalConfig, error)) LocalConfigLoader {
	return LocalConfigLoader{loadFunc: f}
}

func NewMainConfigLoader(f func(string) (*config.MainConfig, error)) MainConfigLoader {
	return MainConfigLoader{loadFunc: f}
}
