package cmd

import (
	"github.com/geoffmore/esctl-go/escmd"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(watcherCmd)
	watcherCmd.AddCommand(watcherPut)
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
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate
		client, err := genClient()
		if err != nil {
			log.Fatal(err)
		}

		err = escmd.WatcherPut(client)
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
		// Boilerplate
		client, err := genClient()
		if err != nil {
			log.Fatal(err)
		}

		err = escmd.WatcherGet(client)
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
		// Boilerplate
		client, err := genClient()
		if err != nil {
			log.Fatal(err)
		}

		err = escmd.WatcherDelete(client)
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
		// Boilerplate
		client, err := genClient()
		if err != nil {
			log.Fatal(err)
		}

		err = escmd.WatcherExecute(client)
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
		client, err := genClient()
		if err != nil {
			log.Fatal(err)
		}

		err = escmd.WatcherPut(client)
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
		// Boilerplate
		client, err := genClient()
		if err != nil {
			log.Fatal(err)
		}

		err = escmd.WatcherAck(client)
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
		// Boilerplate
		client, err := genClient()
		if err != nil {
			log.Fatal(err)
		}

		err = escmd.WatcherActivate(client)
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
		// Boilerplate
		client, err := genClient()
		if err != nil {
			log.Fatal(err)
		}

		err = escmd.WatcherDeactivate(client)
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
		client, err := genClient()
		if err != nil {
			log.Fatal(err)
		}

		err = escmd.WatcherServiceStop(client)
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
		client, err := genClient()
		if err != nil {
			log.Fatal(err)
		}

		err = escmd.WatcherServiceStart(client)
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
		client, err := genClient()
		if err != nil {
			log.Fatal(err)
		}

		err = escmd.WatcherList(client)
		if err != nil {
			log.Fatal(err)
		}
	},
}
