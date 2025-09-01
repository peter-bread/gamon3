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

package cmd

import (
	"log"
	"os"
	"os/exec"
	"peter-bread/gamon3/internal/gamon3cmd"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Switches GH CLI account if necessary",
	Long: `Switches GH CLI account if necessary.

Determines which GH CLI account should be active and compares
this to the currently active account. If they differ, it will switch to
the correct account.`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			mapping gamon3cmd.Mapping
			ghHosts gamon3cmd.GHHosts
		)

		mappingPath, err := gamon3cmd.GetMappingPath()
		if err != nil {
			log.Fatalln(err)
		}

		if err := mapping.Load(mappingPath); err != nil {
			log.Fatalln(err)
		}

		pwd := os.Getenv("PWD")
		account := mapping.GetAccount(pwd)

		ghHostsPath, err := gamon3cmd.GetGHHostsPath()
		if err != nil {
			log.Fatalln(err)
		}

		if err := ghHosts.Load(ghHostsPath); err != nil {
			log.Fatalln(err)
		}

		currentAccount := ghHosts.GetCurrentUser()

		if account != currentAccount {
			// TODO: Handle error.
			exec.Command("gh", "auth", "switch", "--user", account).Run()
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
