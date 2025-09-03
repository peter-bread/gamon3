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
	"fmt"
	"log"
	"os"
	"os/exec"
	"slices"

	"github.com/peter-bread/gamon3/internal/gamon3cmd"

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
		var ghHosts gamon3cmd.GHHosts

		ghHostsPath, err := gamon3cmd.GetGHHostsPath()
		if err != nil {
			log.Fatalln(err)
		}

		if err := ghHosts.Load(ghHostsPath); err != nil {
			log.Fatalln(err)
		}

		currentAccount := ghHosts.GetCurrentUser()
		users := ghHosts.GetAllUsers()

		// Check $GAMON3_ACCOUNT.
		// If not set, continue.
		// If set, check it is a valid account (using 'hosts.yml').
		// If it is valid, switch to it and ignore 'mapping.json'.
		// If it is not valid, error and fallback to 'mapping.json'.
		//
		// This will allow for even finer grained control.
		// It is useful for people who do not split projects
		// by GitHub account. It does depend on something like direnv
		// though.
		//
		// IMPORTANT: $GAMON3_ACCOUNT *must* be exported.

		if account, found := os.LookupEnv("GAMON3_ACCOUNT"); found {
			if slices.Contains(users, account) {
				switchIfNeeded(account, currentAccount)
				return
			} else {
				fmt.Println(account + " is not a valid account")
				fmt.Println("Falling back to main config file")
			}
		}

		// Walk up filetree looking for a '.gamon3.yaml' file which
		// could contain an account to use. I will probably put this behind
		// an optional flag, perhaps '--walk' (name TBD). This would make it
		// and opt-in feature. The reason for this is that it adds overhead and
		// this command should run as fast as possible. In the future it is
		// possible that this becomes the default behaviour and users need to
		// opt-out, maybe by passing the `--no-walk` flag. It should also stop
		// walking at the $HOME directory, at which point it falls back to
		// 'mapping.json' and likely then falls back to 'default'.
		//
		// This will also allow for finer grained control without depending on
		// direnv. The downside is that it will have some effect on performance,
		// however it may be negligible.

		var localConfig gamon3cmd.LocalConfig

		localConfigPath := gamon3cmd.GetLocalConfigPath()

		if localConfigPath != "" {

			localConfig.Load(localConfigPath)

			account := localConfig.Account

			if slices.Contains(users, account) {
				switchIfNeeded(account, currentAccount)
				return
			} else {
				fmt.Println("Invalid local config file: " + localConfigPath)
				fmt.Println("Falling back to main config file")
			}
		}

		// Check 'mapping.json', i.e. main user config.
		var mapping gamon3cmd.Mapping

		mappingPath, err := gamon3cmd.GetMappingPath()
		if err != nil {
			log.Fatalln(err)
		}

		if err := mapping.Load(mappingPath); err != nil {
			log.Fatalln(err)
		}

		// TODO: Use `os.Getwd`?
		pwd := os.Getenv("PWD")
		account := mapping.GetAccount(pwd)

		switchIfNeeded(account, currentAccount)
	},
}

// TODO: Move to internal?
func switchIfNeeded(account, currentAccount string) {
	if account != currentAccount {
		// TODO: Handle error.
		exec.Command("gh", "auth", "switch", "--user", account).Run()
	}
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
