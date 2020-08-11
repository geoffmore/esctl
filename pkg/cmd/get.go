package cmd

import (
	"github.com/geoffmore/esctl/pkg/escmd"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(getClusterInfo)
	getCmd.AddCommand(getClusterHealth)
}

var getCmd = &cobra.Command{
	// esctl get
	Use:   "get",
	Short: "Get a resource",
}

var getClusterInfo = &cobra.Command{
	Use:   "cluster-info",
	Short: "Get cluster-info",
	Long:  `No Description`,
	Run: func(cmd *cobra.Command, args []string) {

		// Boilerplate //
		// Flag init
		initPersistentCmdOpts(cmd, cmdOpts)
		initCmdOpts(cmd, cmdOpts)
		// Client init
		initPersistentCfgOpts(cmd, cfgOpts)
		initCfgOpts(cmd, cfgOpts)
		client, err := genClient(context)
		if err != nil {
			log.Fatal(err)
		}
		// Everything else
		err = escmd.GetClusterInfo(client, cmdOpts)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var getClusterHealth = &cobra.Command{
	Use:   "cluster-health",
	Short: "Get cluster-health",
	Long:  `No Description`,
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate //
		// Flag init
		initPersistentCmdOpts(cmd, cmdOpts)
		initCmdOpts(cmd, cmdOpts)
		// Client init
		initPersistentCfgOpts(cmd, cfgOpts)
		initCfgOpts(cmd, cfgOpts)
		client, err := genClient(context)
		if err != nil {
			log.Fatal(err)
		}
		// Everything else
		err = escmd.GetClusterHealth(client, cmdOpts)
		if err != nil {
			log.Fatal(err)
		}
	},
}
