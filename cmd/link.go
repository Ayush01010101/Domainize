/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// linkCmd represents the link command
var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Link domain to port",
	Long:  `Link your localhost port to a actual domain , for example localhost:300, link 3000 to domain.com`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Your Args ")
	},
}

func init() {
	fmt.Println("Link called")
	rootCmd.AddCommand(linkCmd)

}
