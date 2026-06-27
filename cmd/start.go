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
	Short: "Start serving all linked domains",
	Long: `Start the reverse proxy for every domain you have linked.

It reads your saved configuration and, for each linked domain, forwards
traffic from the domain to its local port. Run "link" first to map a port
to a domain, then run "start" to begin serving them.`,
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
