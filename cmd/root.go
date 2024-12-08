package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "doit",
	Short: "Doit is a simple CLI task manager",
}

func Execute() {
	rootCmd.Execute()
}
