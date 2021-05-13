/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var Reverse bool

// aliasCmd represents the alias command
var aliasCmd = &cobra.Command{
	Use:   "alias new_name old_name",
	Args: cobra.ExactValidArgs(2),
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		newName, oldName := args[0], args[1]
		if Reverse {
			newName, oldName = oldName, newName
		}
		config := LoadConfig()

		config.Alias[newName] = oldName

		rt, err := SaveConfig(config)

		if err != nil {
			fmt.Println("save config", err)
		} else {
			fmt.Print(rt)
		}
	},
}

func init() {
	rootCmd.AddCommand(aliasCmd)
	aliasCmd.Flags().BoolVarP(&Reverse, "reverse", "r", false, "Reverse the args")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// aliasCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// aliasCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
