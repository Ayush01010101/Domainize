/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/Ayush01010101/Custom-Domain-CLI.git/src/functions"
	"github.com/Ayush01010101/Custom-Domain-CLI.git/src/utlities"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called")

		config, err := utlities.ReadConfig()
		if err != nil {
			fmt.Println("could not read config, run link first:", err)
			return
		}

		if len(config.Domain) == 0 {
			fmt.Println("no domain configured, run link first")
			return
		}

		for domain, domainConfig := range config.Domain {
			functions.ReverseProxy(strconv.Itoa(domainConfig.Port), domain)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

}
