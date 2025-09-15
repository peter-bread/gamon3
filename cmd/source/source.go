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
	"os"

	"github.com/peter-bread/gamon3/internal/gamon3cmd"

	"github.com/spf13/cobra"
)

// SourceCmd represents the run command.
var SourceCmd = &cobra.Command{
	Use:   "source",
	Short: "Prints source of current acount",
	Long: `Prints source of current acount.

Prints which GH CLI account should be active and which method was used
to resolve this.

There are three methods used to determine which account should be used:
1. $GAMON3_ACCOUNT environment variable
2. Checking '.gamon.yaml' or '.gamon.yml' local config file
3. Main user config file 'config.yaml'
`,
	Run: func(cmd *cobra.Command, args []string) {
		_, _, source, err := gamon3cmd.Resolve()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(source)
	},
}
