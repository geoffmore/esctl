package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a resource",
	Long:  `No Description`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("esctl get")
	},
}
