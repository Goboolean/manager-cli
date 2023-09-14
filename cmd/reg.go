/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"errors"

	"github.com/Goboolean/manager-cli/internal/adaptor/command"
	"github.com/spf13/cobra"
)

// regCmd represents the reg command
var regCmd = &cobra.Command{
	Use:   "reg {product-id}",
	Short: "Store product metadata to metadata repo",
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
		ctx := context.TODO()

		productType, _ := cmd.Flags().GetString("type")
		productCode, _ := cmd.Flags().GetString("code")
		productName, _ := cmd.Flags().GetString("name")
		productLocation, _ := cmd.Flags().GetString("location")
		productExchange, _ := cmd.Flags().GetString("exchange")

		err := CommandAdaptor.Register(ctx, command.RegisterParms{
			Id:       args[0],
			Type:     productType,
			Code:     productCode,
			Name:     productName,
			Location: productLocation,
			Exchange: productExchange,
		})

		if err != nil {
			print(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(regCmd)

	regCmd.Flags().StringP("type", "t", "", "Type of product available: stock, coin (required)")
	regCmd.Flags().String("code", "", "Product code used in one's market (required)")
	regCmd.Flags().String("name", "", "Name of product in english (required)")
	regCmd.Flags().String("location", "", "Country code defined in ISO 3166-1.")
	regCmd.Flags().String("exchange", "", "Human readable exchange name")

	regCmd.MarkFlagRequired("type")
	regCmd.MarkFlagRequired("code")
	regCmd.MarkFlagRequired("name")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// regCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// regCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
