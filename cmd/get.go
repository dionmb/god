package cmd

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [name]",
	Args: cobra.ExactValidArgs(1),
	Short: "Use an alias to download the package",
	Run: func(_ *cobra.Command, args []string) {
		name := args[0]
		cmd := exec.Command("go", "get", Aliased(name))
		fmt.Println(cmd.String())
		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			fmt.Print(stderr.String())
			return
		}
		fmt.Print(stdout.String())
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
