/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Goboolean/manager-cli/cmd/validator"
	"github.com/notEpsilon/go-pair"
	"github.com/spf13/cobra"
)

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "backup stock data",
	Long:  ``,

	PreRunE: func(cmd *cobra.Command, args []string) error {

		// Pair of string value and its appropriate validator
		var ValueWithValidator [2]*pair.Pair[string, validator.Validator]

		ValueWithValidator[0] = pair.New[string, validator.Validator](
			cmd.Flag("input").Value.String(),
			validator.NewStockValidator())

		ValueWithValidator[1] = pair.New[string, validator.Validator](
			cmd.Flag("before").Value.String(),
			validator.NewDateValidator())

		for _, p := range ValueWithValidator {
			err := p.Second.ValidateString(p.First)
			if p.First != "" && p.Second.ValidateString(p.First) != nil {
				return err
			}
		}

		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {

		//Test code to verify if the flag variable is successfully inputted.
		fmt.Println("flag input: " + cmd.Flag("input").Value.String())
		fmt.Println("flag output: " + cmd.Flag("output").Value.String())
		fmt.Println("flag before: " + cmd.Flag("before").Value.String())

		//todo: call domain code.
		fmt.Println("backup started")
	},
}

func init() {

	rootCmd.AddCommand(backupCmd)

	backupCmd.Flags().StringP("input", "i", "", "Target stock by [StockID]-[location] form")
	backupCmd.Flags().StringP("output", "o", "", "Name of backup file")
	backupCmd.Flags().String("before", "", "Back up data created before yyyy/mm/dd")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// backupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// backupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
