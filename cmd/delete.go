package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete specific task",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("delete called")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
