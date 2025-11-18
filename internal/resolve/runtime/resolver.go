package runtime

import (
	"github.com/peter-bread/gamon3/internal/config"
	"github.com/peter-bread/gamon3/internal/resolve"
)

func NewResolver() *Resolver {
	return &Resolver{
		locator:     Locator{},
		ghLoader:    NewGhHostsLoader(config.LoadGhHosts),
		localLoader: NewLocalConfigLoader(config.LoadLocalConfig),
		mainLoader:  NewMainConfigLoader(config.LoadMainConfig),
		os:          OS{},
	}
}

func (r *Resolver) Resolve() (resolve.Result, error) {
	return resolve.Resolve(
		r.locator,
		r.ghLoader,
		r.localLoader,
		r.mainLoader,
		r.os,
	)
}
