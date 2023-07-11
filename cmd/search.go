/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search {name}",
	Short: "A brief description of your command",

	PreRunE: func(cmd *cobra.Command, args []string) error {

		if len(args) < 1 {
			return ErrInsufficientArgs
		} else if len(args) > 1 {
			return ErrTooManyArgs
		} else {
			return nil
		}
	},

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("search called")
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
