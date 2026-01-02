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

// Package cmd defines root command and is the entry point for this tool.
package cmd

import (
	"fmt"
	"os"

	"github.com/peter-bread/gamon3/v2/cmd/hook"
	"github.com/peter-bread/gamon3/v2/cmd/run"
	"github.com/peter-bread/gamon3/v2/cmd/source"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gamon3",
	Short: "Automatic GH CLI account switching on cd",
	Long: `Automatic GH CLI account switching on cd.

Gamon3 is a tool that enables automatic GH CLI account switching based on a
context, specifically 'pwd'. It does this by hooking into the 'cd' command.
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func SetVersion(version, commit, date, os, arch string) {
	rootCmd.Version = fmt.Sprintf("%s %s-%s\nCommit: %s (%s)", version, os, arch, commit, date)
}

func init() {
	rootCmd.AddCommand(run.RunCmd)
	rootCmd.AddCommand(hook.HookCmd)
	rootCmd.AddCommand(source.SourceCmd)
}
