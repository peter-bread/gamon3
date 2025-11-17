package resolve_test

import (
	"github.com/peter-bread/gamon3/internal/config"
	"github.com/peter-bread/gamon3/internal/resolve"
)

type mockLocator struct {
	envAccount     string
	envFound       bool
	localPath      string
	localErr       error
	mainPath       string
	mainErr        error
	ghHostsPath    string
	ghHostsPathErr error
}

func (m mockLocator) GhHostsPath() (string, error)     { return m.ghHostsPath, m.ghHostsPathErr }
func (m mockLocator) EnvAccount() (string, bool)       { return m.envAccount, m.envFound }
func (m mockLocator) LocalConfigPath() (string, error) { return m.localPath, m.localErr }
func (m mockLocator) MainConfigPath() (string, error)  { return m.mainPath, m.mainErr }

type mockGhHosts struct {
	current string
	users   []string
}

func (m mockGhHosts) CurrentUser() string { return m.current }
func (m mockGhHosts) AllUsers() []string  { return m.users }

type mockGhHostsLoader struct {
	gh  resolve.GhHosts
	err error
}

func (m mockGhHostsLoader) Load(path string) (resolve.GhHosts, error) {
	return m.gh, m.err
}

type mockLocalLoader struct {
	cfg *config.LocalConfig
	err error
}

func (m mockLocalLoader) Load(path string) (*config.LocalConfig, error) {
	return m.cfg, m.err
}

type mockMainLoader struct {
	cfg *config.MainConfig
	err error
}

func (m mockMainLoader) Load(path string) (*config.MainConfig, error) {
	return m.cfg, m.err
}

type mockOS struct {
	cwd string
	err error
}

func (m mockOS) Getwd() (string, error) {
	return m.cwd, m.err
}
