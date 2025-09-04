/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package sucky

import (
	"github.com/Easy1000/duotone/stuff"

	"github.com/spf13/cobra"
)

// butt represents the convert command
var butt = &cobra.Command{
	Use:   "mew [goon file]",
	Short: "I don't even know man",
	Long: `Something something

explanation`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		stuff.ChangingThingsUp(args[0])
	},
}

func init() {
	ewwwwhhhh.AddCommand(butt)
}
