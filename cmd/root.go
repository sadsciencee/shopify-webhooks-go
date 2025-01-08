package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "shopify-webhooks-go",
	Short: "Initialize and introspect Shopify client",
	Long:  `Tooling to automate schema introspection and client generation`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("To download payloads for an API version, run `github.com/sadsciencee/shopify-webhooks-go payloads`")
		fmt.Println("To generate rough models for an API version, run `github.com/sadsciencee/shopify-webhooks-go models *after* running the payloads command`")
	},
}

func init() {
	rootCmd.AddCommand(payloads)
	rootCmd.AddCommand(models)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
