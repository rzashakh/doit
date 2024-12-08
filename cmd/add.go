package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Creates a new task",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("add called")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
