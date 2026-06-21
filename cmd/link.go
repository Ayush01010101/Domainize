/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	proxy "github.com/Ayush01010101/Custom-Domain-CLI.git/src/functions/ReverseProxy"
	hostfile "github.com/Ayush01010101/Custom-Domain-CLI.git/src/functions/UpdateHostsFile"
	"github.com/spf13/cobra"
	"regexp"
	"strconv"
)

var linkNumberRegex = regexp.MustCompile(`^\d{1,4}$`)
var linkDomainRegex = regexp.MustCompile(`(?i)^(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\.)+[a-z]{2,}$`)

// linkCmd represents the link command
var linkCmd = &cobra.Command{
	Use:   "link <number> <domain>",
	Short: "Forward a local port to a domain",
	Long:  `Forward traffic from a local port to a domain, for example link 8080 example.com forwards localhost:8080 to http://example.com.`,
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
		proxy.ReverseProxy(args[0])
	},
}

func init() {
	rootCmd.AddCommand(linkCmd)
}
