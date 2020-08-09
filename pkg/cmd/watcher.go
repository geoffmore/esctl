package cmd

import (
	"fmt"
	"github.com/geoffmore/esctl/pkg/esutil"
	"github.com/geoffmore/esctl/pkg/watcher"
	"github.com/spf13/cobra"
	"log"
)

var initInactive bool

func init() {
	rootCmd.AddCommand(watcherCmd)

	watcherCmd.AddCommand(watcherPut)

	watcherPut.Flags().StringP("input-file", "f", "", "path to file. Use '-' to specify stdin")
	watcherPut.Flags().BoolVarP(&initInactive, "inactive", "", false, "controls whether or not a watcher is initialized as inactive (default active)")
	watcherPut.MarkFlagRequired("input-file")

	watcherCmd.AddCommand(watcherGet)
	watcherCmd.AddCommand(watcherDelete)
	watcherCmd.AddCommand(watcherExecute)
	watcherCmd.AddCommand(watcherAck)
	watcherCmd.AddCommand(watcherActivate)
	watcherCmd.AddCommand(watcherDeactivate)
	watcherCmd.AddCommand(watcherGetStats)
	watcherCmd.AddCommand(watcherServiceStop)
	watcherCmd.AddCommand(watcherServiceStart)
	watcherCmd.AddCommand(watcherList)
	watcherCmd.AddCommand(watcherShowActive)
	watcherCmd.AddCommand(watcherShowInactive)
}

var watcherCmd = &cobra.Command{
	// esctl watcher
	Use:   "watcher",
	Short: "Endpoints under /_watcher",
}

// PUT /_watcher/watch/<watch_id>
var watcherPut = &cobra.Command{
	Use:   "put",
	Short: "PUT /_watcher/watch/<watch_id>",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate
		client, err := genClient(context)
		if err != nil {
			log.Fatal(err)
		}

		// Get the file name
		inputFile, err := cmd.Flags().GetString("input-file")
		if err != nil {
			log.Fatal(err)
		}
		// Generate a reader
		r, err := esutil.FilenameToReader(inputFile)
		if err != nil {
			log.Fatal(err)
		}

		err = watcher.WatcherPut(client, args[0], r, initInactive, outputFmt)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// GET /_watcher/watch/<watch_id>
var watcherGet = &cobra.Command{
	Use:   "get",
	Short: "GET /_watcher/watch/<watch_id>",
	Run: func(cmd *cobra.Command, args []string) {
		// Argument check
		if len(args) != 1 {
			errMsg := fmt.Sprintf("Invalid length of arguments. Expecting 1. Got %d\n", len(args))
			log.Fatal("%s\n", errMsg)
		}
		// Boilerplate
		client, err := genClient(context)
		if err != nil {
			log.Fatal(err)
		}

		//err = watcher.WatcherGet(client)
		err = watcher.WatcherGet(client, args[0], outputFmt)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// DELETE /_watcher/watch/<watch_id>
var watcherDelete = &cobra.Command{
	Use:   "delete",
	Short: "DELETE /_watcher/watch/<watch_id>",
	Run: func(cmd *cobra.Command, args []string) {
		// Argument check
		if len(args) != 1 {
			errMsg := fmt.Sprintf("Invalid length of arguments. Expecting 1. Got %d\n", len(args))
			log.Fatal("%s\n", errMsg)
		}
		// Boilerplate
		client, err := genClient(context)
		if err != nil {
			log.Fatal(err)
		}

		//err = watcher.WatcherDelete(client)
		err = watcher.WatcherDelete(client, args[0], outputFmt)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// POST /_watcher/watch/<watch_id>/_execute
var watcherExecute = &cobra.Command{
	Use:   "execute",
	Short: "POST /watcher/watch/<watch_id>/_execute",
	Run: func(cmd *cobra.Command, args []string) {
		// Argument check
		if len(args) != 1 {
			errMsg := fmt.Sprintf("Invalid length of arguments. Expecting 1. Got %d\n", len(args))
			log.Fatal("%s\n", errMsg)
		}
		// Boilerplate
		client, err := genClient(context)
		if err != nil {
			log.Fatal(err)
		}

		//err = watcher.WatcherExecute(client)
		err = watcher.WatcherExecute(client, args[0], outputFmt)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// GET /_watcher/stats
var watcherGetStats = &cobra.Command{
	Use:   "get-stats",
	Short: "GET /_watcher/stats",
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate
		client, err := genClient(context)
		if err != nil {
			log.Fatal(err)
		}

		err = watcher.WatcherGetStats(client, outputFmt)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// PUT /_watcher/watch/<watch-id>/_ack
var watcherAck = &cobra.Command{
	Use:   "ack",
	Short: "PUT /_watcher/watch/<watch-id>/_ack",
	Run: func(cmd *cobra.Command, args []string) {
		// Argument check
		if len(args) != 1 {
			errMsg := fmt.Sprintf("Invalid length of arguments. Expecting 1. Got %d\n", len(args))
			log.Fatal("%s\n", errMsg)
		}
		// Boilerplate
		client, err := genClient(context)
		if err != nil {
			log.Fatal(err)
		}

		//err = watcher.WatcherAck(client)
		err = watcher.WatcherAck(client, args[0], outputFmt)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// PUT _watcher/watch/<watch_id>/_activate
var watcherActivate = &cobra.Command{
	Use:   "activate",
	Short: "PUT _watcher/watch/<watch_id>/_activate",
	Run: func(cmd *cobra.Command, args []string) {
		// Argument check
		if len(args) != 1 {
			errMsg := fmt.Sprintf("Invalid length of arguments. Expecting 1. Got %d\n", len(args))
			log.Fatal("%s\n", errMsg)
		}
		// Boilerplate
		client, err := genClient(context)
		if err != nil {
			log.Fatal(err)
		}

		//err = watcher.WatcherActivate(client)
		err = watcher.WatcherActivate(client, args[0], outputFmt)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// PUT _watcher/watch/<watch_id>/_deactivate
var watcherDeactivate = &cobra.Command{
	Use:   "deactivate",
	Short: "PUT _watcher/watch/<watch_id>/_deactivate",
	Run: func(cmd *cobra.Command, args []string) {
		// Argument check
		if len(args) != 1 {
			errMsg := fmt.Sprintf("Invalid length of arguments. Expecting 1. Got %d\n", len(args))
			log.Fatal("%s\n", errMsg)
		}
		// Boilerplate
		client, err := genClient(context)
		if err != nil {
			log.Fatal(err)
		}

		//err = watcher.WatcherDeactivate(client)
		err = watcher.WatcherDeactivate(client, args[0], outputFmt)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// POST /_watcher/_stop
var watcherServiceStop = &cobra.Command{
	Use:   "service-stop",
	Short: "POST /_watcher/_stop",
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate
		client, err := genClient(context)
		if err != nil {
			log.Fatal(err)
		}

		err = watcher.WatcherServiceStop(client)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// POST /_watcher/_start
var watcherServiceStart = &cobra.Command{
	Use:   "service-start",
	Short: "POST /_watcher/_start",
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate
		client, err := genClient(context)
		if err != nil {
			log.Fatal(err)
		}

		err = watcher.WatcherServiceStart(client)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// GET /.watches/_search?filter_path=hits.hits._id
var watcherList = &cobra.Command{
	Use:   "list",
	Short: "GET /.watches/_search?filter_path=hits.hits._id",
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate
		client, err := genClient(context)
		if err != nil {
			log.Fatal(err)
		}

		err = watcher.WatcherList(client, outputFmt)
		if err != nil {
			log.Fatal(err)
		}
	},
}
var watcherShowActive = &cobra.Command{
	Use:   "show-active",
	Short: "Show active watchers",
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate
		client, err := genClient(context)
		if err != nil {
			log.Fatal(err)
		}

		err = watcher.WatcherShowActive(client, outputFmt)
		if err != nil {
			log.Fatal(err)
		}
	},
}
var watcherShowInactive = &cobra.Command{
	Use:   "show-inactive",
	Short: "Show inactive watchers",
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate
		client, err := genClient(context)
		if err != nil {
			log.Fatal(err)
		}

		err = watcher.WatcherShowInactive(client, outputFmt)
		if err != nil {
			log.Fatal(err)
		}
	},
}
