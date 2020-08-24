package cmd

import (
	"github.com/geoffmore/esctl/pkg/api"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(apiCmd)
}

var apiCmd = &cobra.Command{
	// esctl api
	Use:   "api",
	Short: "Executes commands using <method> <endpoint> syntax",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate //
		// Command init
		initCmdOpts(cmd, cmdOpts, args)
		// Client init
		initCfgOpts(cmd, cfgOpts)
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
