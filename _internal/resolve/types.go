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
