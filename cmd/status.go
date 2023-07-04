/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"

	"github.com/Goboolean/manager-cli/cmd/validator"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status {stockID}-{Location}",
	Short: "Show stock status",
	Long: `
	{stockID} is the unique code of each stock.
	{Location} is a country code defined in ISO 3166-1.
	For example, country code of korea is "ko" and the united state is "us".
	{Location} must be lower case.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("insufficient args")
		} else if len(args) > 1 {
			return errors.New("too many args")
		} else {
			var v validator.Validator
			v = validator.NewStockValidator()
			return v.ValidateString(args[0])
		}
	},

	Run: func(cmd *cobra.Command, args []string) {

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
