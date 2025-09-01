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
	"path/filepath"
	"peter-bread/gamon3/internal/ghswitch"

	"github.com/spf13/cobra"
)

// mapCmd represents the map command
var mapCmd = &cobra.Command{
	Use:   "map",
	Short: "Creates a path-to-account mapping from config file",
	Long: `Creates a path-to-account mapping from config file.

Reads config file, normalises paths, and creates a mapping.

This command should always be run after you update your config.`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			config  ghswitch.Config
			mapping ghswitch.Mapping
		)

		configPath, err := ghswitch.GetConfigPath()
		if err != nil {
			log.Fatalln(err)
		}

		if err := config.Load(configPath); err != nil {
			fmt.Println(err)
			// TODO: Delete old mapping file?
			os.Exit(1)
		}

		mappingPath, err := ghswitch.GetMappingPath()
		if err != nil {
			log.Fatalln(err)
		}

		if err := os.MkdirAll(filepath.Dir(mappingPath), 0755); err != nil {
			log.Fatalln(err)
		}

		mapping.Create(&config)
		if err := mapping.Save(mappingPath); err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(mapCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mapCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mapCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
