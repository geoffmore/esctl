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
		err = escmd.CatIndices(client, cmdOpts)
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
		err = escmd.CatAliases(client, cmdOpts)
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
		err = escmd.CatTemplates(client, cmdOpts)
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
		err = escmd.CatAllocation(client, cmdOpts)
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
		err = escmd.CatShards(client, cmdOpts)
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
		err = escmd.CatMaster(client, cmdOpts)
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
		err = escmd.CatNodes(client, cmdOpts)
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
		err = escmd.CatTasks(client, cmdOpts)
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
		err = escmd.CatSegments(client, cmdOpts)
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
		err = escmd.CatCount(client, cmdOpts)
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
		err = escmd.CatRecovery(client, cmdOpts)
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
		err = escmd.CatHealth(client, cmdOpts)
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
		err = escmd.CatPendingTasks(client, cmdOpts)
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
		err = escmd.CatThreadPool(client, cmdOpts)
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
		err = escmd.CatPlugins(client, cmdOpts)
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
		err = escmd.CatFielddata(client, cmdOpts)
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
		err = escmd.CatNodeattrs(client, cmdOpts)
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
		err = escmd.CatRepositories(client, cmdOpts)
		if err != nil {
			log.Fatal(err)
		}
	},
}
