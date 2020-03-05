package cmd

import (
	"github.com/geoffmore/esctl-go/pkg/api"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(apiCmd)
	//apiCmd.AddCommand(apiTest)
	//apiCmd.AddCommand(apiInt)
}

var apiCmd = &cobra.Command{
	// esctl api
	Use:   "api",
	Short: "Executes commands using <method> <endpoint> syntax",
	Args:  cobra.ExactArgs(2),
	// I need a persistent flag for output format
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate
		client, err := genClient()
		if err != nil {
			log.Fatal(err)
		}
		// I don't like the function name, but it'll have to... _Do_
		err = api.MakeGenericRequest(client, outputFmt, args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}
	},
}

//var apiTest = &cobra.Command{
//	// esctl api test
//	Use:   "test-cmd",
//	Short: "Generated code to demonstrate project structure",
//	Run: func(cmd *cobra.Command, args []string) {
//		api.CmdTest()
//	},
//}
//
//var apiInt = &cobra.Command{
//	// esctl api test
//	Use:   "test-int",
//	Short: "Generated code to demonstrate project structure",
//	Run: func(cmd *cobra.Command, args []string) {
//		// Boilerplate
//		client, err := genClient()
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		err = api.IntegrationTest(client)
//		if err != nil {
//			log.Fatal(err)
//		}
//	},
//}
