package resolve

import "github.com/peter-bread/gamon3/internal/config"

type mockGhHosts struct {
	current string
	users   []string
}

func (m mockGhHosts) CurrentUser() string { return m.current }
func (m mockGhHosts) AllUsers() []string  { return m.users }

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
