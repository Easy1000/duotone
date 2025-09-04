/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package sucky

import (
	"os"

	"github.com/spf13/cobra"
)

var ewwwwhhhh = &cobra.Command{
	Use:   "tung",
	Short: "cli app i guess",
	Long: `maybe good

maybe shit`,
}

func Execute() {
	err := ewwwwhhhh.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	ewwwwhhhh.CompletionOptions.HiddenDefaultCmd = true
}
