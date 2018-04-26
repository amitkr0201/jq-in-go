// Copyright © 2018 Amit Kumar
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/amitkr0201/jq-in-go/cmd/jqingo"
	"github.com/spf13/cobra"
)

var compact bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Short: "Implementation of jq in Go.",
	Long: `Implemeting jq in Go.
	
Original jq: https://github.com/stedolan/jq
Also a pathway to learn Go.`,
	Version: "0.0.1",
	Example: "jq-in-go -c '.' ~/input.json",
	Use:     "jq-in-go '<selector>' <input-json-file>",

	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Usage()
			os.Exit(1)
		}
		if err := jqingo.ProcessCommand(compact, args); err != nil {
			fmt.Printf("%v", err.Error())
			os.Exit(1)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// rootCmd.PersistentFlags().BoolVarP(&compact, "compact", "c", false, "Print JSON in compact.")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolVarP(&compact, "compact", "c", false, "Print JSON in compact")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

}
