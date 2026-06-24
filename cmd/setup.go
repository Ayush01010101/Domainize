/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Ayush01010101/Custom-Domain-CLI.git/src/utlities"
	"github.com/pterm/pterm"
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

		if config && !reset {
			pterm.Warning.Println("Config already present. Run 'setup --reset' for a clean installation.")
			return
		}

		pterm.DefaultSection.Println("Setting up domainize")

		if reset {
			pterm.Info.Println("Resetting existing configuration...")
		}

		spinner, _ := pterm.DefaultSpinner.Start("Writing config file")
		utlities.SetupConfig()
		spinner.Success("Config file ready")

		utlities.Installmkcert()

		pterm.Success.Println("Setup complete!")
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
