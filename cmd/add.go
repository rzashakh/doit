package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Creates a new task",
	Run: func(cmd *cobra.Command, args []string) {
		argument := strings.Join(args, " ")
		fmt.Println(argument)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
