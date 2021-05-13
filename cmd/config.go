package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"gopkg.in/yaml.v3"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Print the current config",
	Run: func(cmd *cobra.Command, args []string) {
		c := viper.AllSettings()
		rt, err := yaml.Marshal(c)
		if err != nil {
			fmt.Println("read file fail", err)
			log.Fatalf("unable to marshal config to YAML: %v", err)
		} else {
			fmt.Print(string(rt))
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
