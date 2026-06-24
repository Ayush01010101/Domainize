/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/Ayush01010101/Custom-Domain-CLI.git/src/utlities"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "for checking health of the domainize",
	Long:  `for checking health of the domainize`,
	Run: func(cmd *cobra.Command, args []string) {
		config := utlities.Isconfigpresent()
		if config {
			fmt.Print("domainize is healthy")
			return
		}
		fmt.Print("config not present  -- run domainize setup to fix")

	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

}
