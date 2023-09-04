/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "get-status {ProductId}",
	Short: "Show stock status",
	Long: `
	{stockID} is the unique code of each stock.
	For example, country code of korea is "ko" and the united state is "us".
	{Location} must be lower case.`,

	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return ErrInsufficientArgs
		} else if len(args) > 1 {
			return ErrTooManyArgs
		} else {
			return nil
		}
	},

	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.TODO()
		res, err := CommandAdaptor.GetStatus(ctx, args[0])
		println(res)

		if err != nil {
			println(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
