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

// Package source defines the `source` command.
package source

import (
	"fmt"

	"github.com/peter-bread/gamon3/v2/internal/resolve/runtime"

	"github.com/spf13/cobra"
)

// SourceCmd represents the source command.
var SourceCmd = &cobra.Command{
	Use:   "source",
	Short: "Prints source of current acount",
	Long: `Prints source of current acount.

Prints which GH CLI account should be active and which method was used
to resolve this.

If account resolution is successful, the output has the following format:
  <account> <source kind> <source value>

If account resolution is unsuccessful, an error message will be printed and
the program will exit with a non-zero exit code.

There are three methods used to determine which account should be used:
1. $GAMON3_ACCOUNT environment variable
2. Checking '.gamon.yaml' or '.gamon3.yaml' local config file
3. Main user config file 'config.yaml'
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		resolver := runtime.NewResolver()

		result, err := resolver.Resolve()
		if err != nil {
			return err
		}

		_, err = fmt.Println(result.Account, result.SourceKind, result.SourceValue)

		return err
	},
}
