package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of tau-git",
	Long:  "Print the version number of tau-git",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version 0.0 -- HEAD")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
