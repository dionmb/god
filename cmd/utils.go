package cmd

import (
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
	"fmt"
)

func Aliased(name string) string {
	version := ""
	rt := strings.Split(name, "@")
	if len(rt) == 2 {
		name, version = rt[0], rt[1]
	}

	alias := viper.GetStringMapString("alias")
	val, b := alias[name]

	if b {
		name = val
 	}

	if version != "" && !strings.Contains(name, "@") {
		return name + "@" + version
	} else {
		return name
	}
}

func CreateConfigFile() {
	home, err := homedir.Dir()
	cobra.CheckErr(err)
	file := home + ".godrc" + ".yaml"
	if err = viper.WriteConfigAs(file); err != nil {
		fmt.Println(err)
	}
}