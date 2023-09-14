/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:   "backup {ProductId(s)}",
	Short: "backup stock data",
	Long:  ``,

	PreRunE: func(cmd *cobra.Command, args []string) error {
		// Pair of string value and its appropriate validator

		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
		isDataToRemote, _ := cmd.Flags().GetBool("remote")
		backupType, _ := cmd.Flags().GetString("type")

		var err error

		ctx := context.TODO()
		if len(args) == 0 {
			err = CommandAdaptor.BackupTrade(ctx, backupType, isDataToRemote)
		} else {
			for _, item := range args {
				err = CommandAdaptor.BackupProduct(ctx, item, backupType, isDataToRemote)
			}
		}

		if err != nil {
			print(err.Error())
		}
	},
}

func init() {

	rootCmd.AddCommand(backupCmd)

	backupCmd.Flags().StringP("type", "t", "full", "Type of backup. possible value: full, diff (required)")
	backupCmd.Flags().BoolP("remote", "r", false, "Uploads backup(s) to remote")
	regCmd.MarkFlagRequired("type")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// backupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// backupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
