package cmd

import (
	"fmt"
	"github.com/geoffmore/esctl-go/esauth"
	"github.com/geoffmore/esctl-go/escmd"
	"github.com/spf13/cobra"
)

var es = esauth.EsAuth()

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(getClusterInfo)
	getCmd.AddCommand(getClusterHealth)
}

var getCmd = &cobra.Command{
	// esctl get
	Use:   "get",
	Short: "Get a resource",
	Long:  `No Description`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("esctl get")
	},
}

var getClusterInfo = &cobra.Command{
	// esctl get cluster-info
	Use:   "cluster-info",
	Short: "Get cluster-info",
	Long:  `No Description`,
	Run: func(cmd *cobra.Command, args []string) {
		escmd.GetClusterInfo(es)
	},
}
var getClusterHealth = &cobra.Command{
	// esctl get cluster-health
	Use:   "cluster-health",
	Short: "Get cluster-health",
	Long:  `No Description`,
	Run: func(cmd *cobra.Command, args []string) {
		escmd.GetClusterHealth(es)
	},
}
