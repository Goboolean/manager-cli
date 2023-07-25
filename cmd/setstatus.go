/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// setstatusCmd represents the setstatus command
var setstatusCmd = &cobra.Command{
	Use:   "setstatus {status} {stockId}",
	Short: "Change the status of a specific stock",
	Long:  ``,

	Args: func(cmd *cobra.Command, args []string) error {

		if len(args) < 2 {
			return ErrInsufficientArgs
		} else if len(args) > 2 {
			return ErrTooManyArgs
		} else {
			return nil
		}

	},

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("setstatus called")
	},
}

func init() {
	rootCmd.AddCommand(setstatusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setstatusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setstatusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
