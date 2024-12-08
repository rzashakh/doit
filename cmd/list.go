package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all your unfinished tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("add called")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
