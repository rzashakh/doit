package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete specific task",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("delete called")
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
