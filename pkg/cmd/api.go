package cmd

import (
	"github.com/geoffmore/esctl/pkg/api"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(apiCmd)
	apiCmd.Flags().StringVarP(&bodyFile, "body", "f", "", "Path to file to use as input.")
}

var apiCmd = &cobra.Command{
	// esctl api
	Use:   "api",
	Short: "Executes commands using <method> <endpoint> syntax",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate //
		// Command init
		err := initCmdOpts(cmd, cmdOpts, args)
		if err != nil {
			log.Fatal(err)
		}
		// Client init
		err = initCfgOpts(cmd, cfgOpts)
		if err != nil {
			log.Fatal(err)
		}

		client, err := genClientWOpts(cfgOpts)
		if err != nil {
			log.Fatal(err)
		}
		// Everything else
		// The function name isn't great, but it'll have to... _Do_
		err = api.MakeGenericRequest(client, cmdOpts)
		if err != nil {
			log.Fatal(err)
		}
	},
}
