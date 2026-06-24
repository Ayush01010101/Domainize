/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/Ayush01010101/Custom-Domain-CLI.git/src/utlities"
	"github.com/spf13/cobra"
)

var reset bool

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "for setup the customcli",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		config := utlities.Isconfigpresent()
		if reset {
			fmt.Print("reset is called")
			utlities.SetupConfig()
			utlities.Installmkcert()
			return

		}
		if config {
			fmt.Print("config is present run --reset to clean installation ")
			return
		}

		utlities.SetupConfig()
		utlities.Installmkcert()
		fmt.Print("setup is done")
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
	setupCmd.Flags().BoolVar(
		&reset,
		"reset",
		false,
		"reset the config file",
	)

}
