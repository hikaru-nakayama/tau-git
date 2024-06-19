/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize tau-git",
	Long:  "This is a commmand to initialize tau-git",

	RunE: func(cmd *cobra.Command, args []string) error {
		curPath, err := os.Getwd()
		if err != nil {
			return errors.New("fail to get current path")
		}
		dir := filepath.Join(curPath, ".tau")
		if err := os.Mkdir(dir, 0750); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	// initCmd.RunE(&cobra.Command{}, make([]string, 0))
	rootCmd.AddCommand(initCmd)
}
