/*
Copyright © 2025 Peter Sheehan

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package resolve

import "github.com/peter-bread/gamon3/v2/internal/config"

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
