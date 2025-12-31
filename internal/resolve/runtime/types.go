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

package runtime

// Define implementations for interfaces.

import (
	"os"

	"github.com/peter-bread/gamon3/v2/internal/config"
	"github.com/peter-bread/gamon3/v2/internal/locator"
	"github.com/peter-bread/gamon3/v2/internal/resolve"
)

type Locator struct{}

func (Locator) GhHostsPath() (string, error)     { return locator.GhHostsPath(OS{}) }
func (Locator) EnvAccount() (string, bool)       { return locator.EnvAccount(OS{}) }
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
