package cmd

import (
	"github.com/geoffmore/esctl-go/escmd"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(catCmd)
	catCmd.AddCommand(catIndices)
	catCmd.AddCommand(catAliases)
	catCmd.AddCommand(catTemplates)
}

var catCmd = &cobra.Command{
	// esctl get
	Use:   "cat",
	Short: "Cat a resource",
}

var catIndices = &cobra.Command{
	Use:   "indices",
	Short: "Cat indices",
	Long:  `No Description`,
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate
		client, err := genClient()
		if err != nil {
			log.Fatal(err)
		}

		err = escmd.CatIndices(client)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var catAliases = &cobra.Command{
	Use:   "aliases",
	Short: "Cat aliases",
	Long:  `No Description`,
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate
		client, err := genClient()
		if err != nil {
			log.Fatal(err)
		}

		err = escmd.CatAliases(client)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var catTemplates = &cobra.Command{
	Use:   "templates",
	Short: "Cat templates",
	Long:  `No Description`,
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate
		client, err := genClient()
		if err != nil {
			log.Fatal(err)
		}

		err = escmd.CatTemplates(client)
		if err != nil {
			log.Fatal(err)
		}
	},
}
