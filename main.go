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

package main

import (
	"fmt"
	"os"

	"github.com/peter-bread/gamon3/cmd"
	"github.com/peter-bread/gamon3/internal/config"
	"github.com/peter-bread/gamon3/internal/locator"
	"github.com/peter-bread/gamon3/internal/resolve"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

type Locator struct{}

func (Locator) GhHostsPath() (string, error)     { return locator.GhHostsPath() }
func (Locator) EnvAccount() (string, bool)       { return locator.EnvAccount() }
func (Locator) LocalConfigPath() (string, error) { return locator.LocalConfigPath() }
func (Locator) MainConfigPath() (string, error)  { return locator.MainConfigPath() }

type Loader[T any] struct {
	loadFunc func(string) (T, error)
}

func (r Loader[T]) Load(path string) (T, error) {
	return r.loadFunc(path)
}

type OS struct{}

func (OS) Getwd() (string, error) { return os.Getwd() }

func init() {
	ghLoader := Loader[resolve.GhHosts]{loadFunc: func(path string) (resolve.GhHosts, error) {
		return config.LoadGhHosts(path)
	}}

	localLoader := Loader[*config.LocalConfig]{loadFunc: config.LoadLocalConfig}

	mainLoader := Loader[*config.MainConfig]{loadFunc: config.LoadMainConfig}

	fmt.Println(resolve.Resolve(Locator{}, ghLoader, localLoader, mainLoader, OS{}))
	os.Exit(0)
}

func main() {
	cmd.SetVersion(version, commit, date)
	cmd.Execute()
}
