// +build optional ext api

package cmd

import (
	"github.com/geoffmore/esctl-go/pkg/api"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(apiCmd)
	apiCmd.AddCommand(apiTest)
	apiCmd.AddCommand(apiInt)
}

var apiCmd = &cobra.Command{
	// esctl api
	Use:   "api",
	Short: "No description",
}

var apiTest = &cobra.Command{
	// esctl api test
	Use:   "test-cmd",
	Short: "Generated code to demonstrate project structure",
	Run: func(cmd *cobra.Command, args []string) {
		api.CmdTest()
	},
}

var apiInt = &cobra.Command{
	// esctl api test
	Use:   "test-int",
	Short: "Generated code to demonstrate project structure",
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate
		client, err := genClient()
		if err != nil {
			log.Fatal(err)
		}

		err = api.IntegrationTest(client)
		if err != nil {
			log.Fatal(err)
		}
	},
}
