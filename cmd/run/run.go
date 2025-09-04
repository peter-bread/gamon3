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

package run

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"slices"

	"github.com/peter-bread/gamon3/internal/gamon3cmd"

	"github.com/spf13/cobra"
)

// RunCmd represents the run command
var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "Switches GH CLI account if necessary",
	Long: `Switches GH CLI account if necessary.

Determines which GH CLI account should be active and compares
this to the currently active account. If they differ, it will switch to
the correct account.

There are three methods used to determine which account should be used:
1. $GAMON3_ACCOUNT environment variable
2. Checking '.gamon.yaml' or '.gamon.yml' local config file
3. JSON mapping generated from 'config.yaml'
`,
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
		// IMPORTANT: $GAMON3_ACCOUNT *must* be exported.

		if account, found := os.LookupEnv("GAMON3_ACCOUNT"); found {
			if slices.Contains(users, account) {
				switchIfNeeded(account, currentAccount)
				return
			} else {
				fmt.Println(account + " is not a valid account")
				fmt.Println("Falling back to main config file...")
			}
		}

		// Walk up filetree looking for a local '.gamon3.yaml' file.
		// It should also stop walking at the $HOME directory, at which point it
		// falls back to 'config.yaml'.

		var localConfig gamon3cmd.LocalConfig

		if localConfigPath, err := gamon3cmd.GetLocalConfigPath(); err == nil {
			if err := localConfig.Load(localConfigPath, users); err != nil {
				fmt.Println("Invalid local config file: " + localConfigPath)
				fmt.Println(err)
				fmt.Println("Falling back to main config file...")
			} else {
				switchIfNeeded(localConfig.Account, currentAccount)
				return
			}
		}

		// Check main 'config.yaml' file.

		var config gamon3cmd.Config

		configPath, err := gamon3cmd.GetConfigPath()
		if err != nil {
			fmt.Println(err)
		}

		if err := config.Load(configPath, users); err != nil {
			fmt.Println("Invalid config file: " + configPath)
			fmt.Println(err)
			fmt.Println("Exiting...")
		} else {
			pwd := os.Getenv("PWD")
			account := config.GetAccount(pwd)
			switchIfNeeded(account, currentAccount)
		}
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
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
