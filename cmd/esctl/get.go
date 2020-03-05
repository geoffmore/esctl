package cmd

import (
	"github.com/geoffmore/esctl-go/escmd"
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
		// Boilerplate
		client, err := genClient()
		if err != nil {
			log.Fatal(err)
		}

		err = escmd.GetClusterInfo(client)
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
		// Boilerplate
		client, err := genClient()
		if err != nil {
			log.Fatal(err)
		}
		err = escmd.GetClusterHealth(client, outputFmt)
		if err != nil {
			log.Fatal(err)
		}
	},
}
