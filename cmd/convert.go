/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"duotone/function"

	"github.com/spf13/cobra"
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert [image file]",
	Short: "Apply duotone effect on an image",
	Long: `Apply duotone effect on an image

The output path is duotone.jpeg
Only accept 1 image`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		function.Convert(args[0])
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)
}
