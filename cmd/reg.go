/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// regCmd represents the reg command
var regCmd = &cobra.Command{
	Use:   "reg ",
	Short: "Add a stock metadata to DB",
	Long:  ``,

	PreRunE: func(cmd *cobra.Command, args []string) error {

		isAuto, _ := cmd.Flags().GetBool("auto")
		exchange, _ := cmd.Flags().GetString("exchange")

		if !isAuto && exchange == "" {
			return errors.New("expected \"exchange\" flag")
		}
		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("reg called")
	},
}

func init() {
	rootCmd.AddCommand(regCmd)

	regCmd.Flags().StringP("type", "t", "", "Type of product available: stock, coin")
	regCmd.Flags().String("code", "", "Product code used in one's market")
	regCmd.Flags().String("name", "", "English name of product")
	regCmd.Flags().String("location", "", "Country code defined in ISO 3166-1.")
	regCmd.Flags().String("exchange", "", "Human readable exchange name")

	regCmd.Flags().BoolP("auto", "a", false, "fetch stock metadata automatically")

	regCmd.MarkFlagRequired("type")
	regCmd.MarkFlagRequired("code")
	regCmd.MarkFlagRequired("name")
	regCmd.MarkFlagRequired("location")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// regCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// regCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
