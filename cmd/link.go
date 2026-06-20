/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"regexp"
	"strconv"

	hostfile "github.com/Ayush01010101/Custom-Domain-CLI.git/src/functions"
	"github.com/spf13/cobra"
)

var linkNumberRegex = regexp.MustCompile(`^\d{1,4}$`)
var linkDomainRegex = regexp.MustCompile(`(?i)^(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\.)+[a-z]{2,}$`)

// linkCmd represents the link command
var linkCmd = &cobra.Command{
	Use:   "link <number> <domain>",
	Short: "Link domain to port",
	Long:  `Link your localhost port to a actual domain , for example localhost:300, link 3000 to domain.com`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return fmt.Errorf("link requires exactly 2 arguments: <number> <domain>")
		}

		if !linkNumberRegex.MatchString(args[0]) {
			return fmt.Errorf("first argument must contain only numbers and cannot be more than 4 digits")
		}

		if !linkDomainRegex.MatchString(args[1]) {
			return fmt.Errorf("second argument must be a valid domain, for example myname.com")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Your Args ", args)
		port, _ := strconv.Atoi(args[0])
		hostfile.UpdateHostsFile(port, args[1])
	},
}

func init() {
	rootCmd.AddCommand(linkCmd)
}
