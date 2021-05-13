package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Reverse bool

var aliasCmd = &cobra.Command{
	Use:   "alias [name] [package]",
	Args: cobra.ExactValidArgs(2),
	Short: "Add an alias to the package",
	Run: func(cmd *cobra.Command, args []string) {
		newName, oldName := args[0], args[1]
		if Reverse {
			newName, oldName = oldName, newName
		}
		alias := viper.GetStringMapString("alias")
		alias[newName] = oldName
		viper.Set("alias", alias)
		if err := viper.WriteConfig(); err != nil {
			CreateConfigFile()
		}
	},
}

func init() {
	rootCmd.AddCommand(aliasCmd)
	aliasCmd.Flags().BoolVarP(&Reverse, "reverse", "r", false, "Reverse the args")
}
