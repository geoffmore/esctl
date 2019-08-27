package cmd

import (
	"fmt"
	"github.com/geoffmore/esctl-go/esauth"
	"github.com/spf13/cobra"
	"os"
)

// configPath will come later
var configPath string
var es = esauth.EsAuth()

func init() {
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "", "Path to configuration")
}

var rootCmd = &cobra.Command{
	Use:   "esctl",
	Short: "esctl is a utility able to interact with elasticsearch clusters",
	Long:  `No Description`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
