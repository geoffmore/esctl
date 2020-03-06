package cmd

import (
	"github.com/geoffmore/esctl/pkg/escmd"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(catCmd)
	catCmd.AddCommand(catIndices)
	catCmd.AddCommand(catAliases)
	catCmd.AddCommand(catTemplates)
	catCmd.AddCommand(catAllocation)
	catCmd.AddCommand(catShards)
	catCmd.AddCommand(catMaster)
	catCmd.AddCommand(catNodes)
	catCmd.AddCommand(catTasks)
	catCmd.AddCommand(catSegments)
	catCmd.AddCommand(catCount)
	catCmd.AddCommand(catRecovery)
	catCmd.AddCommand(catHealth)
	catCmd.AddCommand(catPendingTasks)
	catCmd.AddCommand(catThreadPool)
	catCmd.AddCommand(catPlugins)
	catCmd.AddCommand(catFielddata)
	catCmd.AddCommand(catNodeattrs)
	catCmd.AddCommand(catRepositories)
}

var catCmd = &cobra.Command{
	// esctl get
	Use:   "cat",
	Short: "Endpoints under /_cat",
}

// GET /_cat/indices
var catIndices = &cobra.Command{
	Use:   "indices",
	Short: "GET /_cat/indices",
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate
		client, err := genClient2(context)
		if err != nil {
			log.Fatal(err)
		}

		err = escmd.CatIndices(client, outputFmt)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// GET /_cat/aliases
var catAliases = &cobra.Command{
	Use:   "aliases",
	Short: "GET /_cat/aliases",
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate
		client, err := genClient2(context)
		if err != nil {
			log.Fatal(err)
		}

		err = escmd.CatAliases(client, outputFmt)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// GET /_cat/templates
var catTemplates = &cobra.Command{
	Use:   "templates",
	Short: "GET /_cat/templates",
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate
		client, err := genClient2(context)
		if err != nil {
			log.Fatal(err)
		}

		err = escmd.CatTemplates(client, outputFmt)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// GET /_cat/allocation
var catAllocation = &cobra.Command{
	Use:   "allocation",
	Short: "GET /_cat/allocation",
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate
		client, err := genClient2(context)
		if err != nil {
			log.Fatal(err)
		}

		err = escmd.CatAllocation(client, outputFmt)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// GET /_cat/shards
var catShards = &cobra.Command{
	Use:   "shards",
	Short: "GET /_cat/shards",
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate
		client, err := genClient2(context)
		if err != nil {
			log.Fatal(err)
		}

		err = escmd.CatShards(client, outputFmt)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// GET /_cat/master
var catMaster = &cobra.Command{
	Use:   "master",
	Short: "GET /_cat/master",
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate
		client, err := genClient2(context)
		if err != nil {
			log.Fatal(err)
		}

		err = escmd.CatMaster(client, outputFmt)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// GET /_cat/nodes
var catNodes = &cobra.Command{
	Use:   "nodes",
	Short: "GET /_cat/nodes",
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate
		client, err := genClient2(context)
		if err != nil {
			log.Fatal(err)
		}

		err = escmd.CatNodes(client, outputFmt)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// GET /_cat/tasks
var catTasks = &cobra.Command{
	Use:   "tasks",
	Short: "GET /_cat/tasks",
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate
		client, err := genClient2(context)
		if err != nil {
			log.Fatal(err)
		}

		err = escmd.CatTasks(client, outputFmt)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// GET /_cat/segments
var catSegments = &cobra.Command{
	Use:   "segments",
	Short: "GET /_cat/segments",
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate
		client, err := genClient2(context)
		if err != nil {
			log.Fatal(err)
		}

		err = escmd.CatSegments(client, outputFmt)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// GET /_cat/count
var catCount = &cobra.Command{
	Use:   "count",
	Short: "GET /_cat/count",
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate
		client, err := genClient2(context)
		if err != nil {
			log.Fatal(err)
		}

		err = escmd.CatCount(client, outputFmt)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// GET /_cat/recovery
var catRecovery = &cobra.Command{
	Use:   "recovery",
	Short: "GET /_cat/recovery",
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate
		client, err := genClient2(context)
		if err != nil {
			log.Fatal(err)
		}

		err = escmd.CatRecovery(client, outputFmt)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// GET /_cat/health
var catHealth = &cobra.Command{
	Use:   "health",
	Short: "GET /_cat/health",
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate
		client, err := genClient2(context)
		if err != nil {
			log.Fatal(err)
		}

		err = escmd.CatHealth(client, outputFmt)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// GET /_cat/pending_tasks
var catPendingTasks = &cobra.Command{
	Use:   "pending-tasks",
	Short: "GET /_cat/pending_tasks",
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate
		client, err := genClient2(context)
		if err != nil {
			log.Fatal(err)
		}

		err = escmd.CatPendingTasks(client, outputFmt)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// GET /_cat/thread_pool
var catThreadPool = &cobra.Command{
	Use:   "thread-pool",
	Short: "GET /_cat/thread_pool",
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate
		client, err := genClient2(context)
		if err != nil {
			log.Fatal(err)
		}

		err = escmd.CatThreadPool(client, outputFmt)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// GET /_cat/plugins
var catPlugins = &cobra.Command{
	Use:   "plugins",
	Short: "GET /_cat/plugins",
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate
		client, err := genClient2(context)
		if err != nil {
			log.Fatal(err)
		}

		err = escmd.CatPlugins(client, outputFmt)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// GET /_cat/fielddata
var catFielddata = &cobra.Command{
	Use:   "fielddata",
	Short: "GET /_cat/fielddata",
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate
		client, err := genClient2(context)
		if err != nil {
			log.Fatal(err)
		}

		err = escmd.CatFielddata(client, outputFmt)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// GET /_cat/nodeattrs
var catNodeattrs = &cobra.Command{
	Use:   "nodeattrs",
	Short: "GET /_cat/nodeattrs",
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate
		client, err := genClient2(context)
		if err != nil {
			log.Fatal(err)
		}

		err = escmd.CatNodeattrs(client, outputFmt)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// GET /_cat/repositories
var catRepositories = &cobra.Command{
	Use:   "repositories",
	Short: "GET /_cat/repositories",
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate
		client, err := genClient2(context)
		if err != nil {
			log.Fatal(err)
		}

		err = escmd.CatRepositories(client, outputFmt)
		if err != nil {
			log.Fatal(err)
		}
	},
}
