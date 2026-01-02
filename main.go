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
	"runtime"
	"runtime/debug"

	"github.com/peter-bread/gamon3/v2/cmd"
)

// VCS build info. These should be set via debug.ReadBuildInfo.
var (
	version = "dev"
	commit  = "none"
	date    = "unknown" // Specifically commit date, not build date. This is for reproducible builds.
)

func buildSettingsMap(info *debug.BuildInfo) map[string]string {
	m := make(map[string]string)

	for _, s := range info.Settings {
		m[s.Key] = s.Value
	}

	return m
}

func init() {
	var (
		os   string
		arch string
	)

	if info, ok := debug.ReadBuildInfo(); ok {
		settings := buildSettingsMap(info)

		os = settings["GOOS"]
		arch = settings["GOARCH"]

		if info.Main.Version != "(devel)" {
			version = info.Main.Version
			commit = settings["vcs.revision"]
			date = settings["vcs.time"]
		}

	} else {
		os = runtime.GOOS
		arch = runtime.GOARCH
	}

	cmd.SetVersion(version, commit, date, os, arch)
}

func main() {
	cmd.Execute()
}
