/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Goboolean/manager-cli/inject"
	"github.com/Goboolean/manager-cli/internal/adaptor/command"
	"github.com/Goboolean/manager-cli/util/env"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "manager-cli {command}",
	Short: "Manage goboolean service. ",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}

}

var CommandAdaptor *command.CommandAdaptor

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	env.Init()
	var err error
	CommandAdaptor, err = inject.InitCommandAdaptor()

	if err != nil {
		panic(err)
	}

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.manager-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}
