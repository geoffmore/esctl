package cmd

import (
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
		// Boilerplate //
		// Command init
		initCmdOpts(cmd, cmdOpts)
		// Client init
		initCfgOpts(cmd, cfgOpts)
		client, err := genClientWOpts(cfgOpts)
		if err != nil {
			log.Fatal(err)
		}
		// Everything else
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

		err = watcher.WatcherPut(client, args[0], r, cmdOpts)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// GET /_watcher/watch/<watch_id>
var watcherGet = &cobra.Command{
	Use:   "get",
	Short: "GET /_watcher/watch/<watch_id>",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate //
		// Command init
		initCmdOpts(cmd, cmdOpts)
		// Client init
		initCfgOpts(cmd, cfgOpts)
		client, err := genClientWOpts(cfgOpts)
		if err != nil {
			log.Fatal(err)
		}
		// Everything else
		err = watcher.WatcherGet(client, args[0], cmdOpts)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// DELETE /_watcher/watch/<watch_id>
var watcherDelete = &cobra.Command{
	Use:   "delete",
	Short: "DELETE /_watcher/watch/<watch_id>",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate //
		// Command init
		initCmdOpts(cmd, cmdOpts)
		// Client init
		initCfgOpts(cmd, cfgOpts)
		client, err := genClientWOpts(cfgOpts)
		if err != nil {
			log.Fatal(err)
		}
		// Everything else
		err = watcher.WatcherDelete(client, args[0], cmdOpts)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// POST /_watcher/watch/<watch_id>/_execute
var watcherExecute = &cobra.Command{
	Use:   "execute",
	Short: "POST /watcher/watch/<watch_id>/_execute",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate //
		// Command init
		initCmdOpts(cmd, cmdOpts)
		// Client init
		initCfgOpts(cmd, cfgOpts)
		client, err := genClientWOpts(cfgOpts)
		if err != nil {
			log.Fatal(err)
		}
		// Everything else
		err = watcher.WatcherExecute(client, args[0], cmdOpts)
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
		// Boilerplate //
		// Command init
		initCmdOpts(cmd, cmdOpts)
		// Client init
		initCfgOpts(cmd, cfgOpts)
		client, err := genClientWOpts(cfgOpts)
		if err != nil {
			log.Fatal(err)
		}
		// Everything else
		err = watcher.WatcherGetStats(client, cmdOpts)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// PUT /_watcher/watch/<watch-id>/_ack
var watcherAck = &cobra.Command{
	Use:   "ack",
	Short: "PUT /_watcher/watch/<watch-id>/_ack",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate //
		// Command init
		initCmdOpts(cmd, cmdOpts)
		// Client init
		initCfgOpts(cmd, cfgOpts)
		client, err := genClientWOpts(cfgOpts)
		if err != nil {
			log.Fatal(err)
		}
		// Everything else
		err = watcher.WatcherAck(client, args[0], cmdOpts)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// PUT _watcher/watch/<watch_id>/_activate
var watcherActivate = &cobra.Command{
	Use:   "activate",
	Short: "PUT _watcher/watch/<watch_id>/_activate",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate //
		// Command init
		initCmdOpts(cmd, cmdOpts)
		// Client init
		initCfgOpts(cmd, cfgOpts)
		client, err := genClientWOpts(cfgOpts)
		if err != nil {
			log.Fatal(err)
		}
		// Everything else
		err = watcher.WatcherActivate(client, args[0], cmdOpts)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// PUT _watcher/watch/<watch_id>/_deactivate
var watcherDeactivate = &cobra.Command{
	Use:   "deactivate",
	Short: "PUT _watcher/watch/<watch_id>/_deactivate",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate //
		// Command init
		initCmdOpts(cmd, cmdOpts)
		// Client init
		initCfgOpts(cmd, cfgOpts)
		client, err := genClientWOpts(cfgOpts)
		if err != nil {
			log.Fatal(err)
		}
		// Everything else
		err = watcher.WatcherDeactivate(client, args[0], cmdOpts)
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
		// Boilerplate //
		// Command init
		initCmdOpts(cmd, cmdOpts)
		// Client init
		initCfgOpts(cmd, cfgOpts)
		client, err := genClientWOpts(cfgOpts)
		if err != nil {
			log.Fatal(err)
		}
		// Everything else
		err = watcher.WatcherServiceStop(client, cmdOpts)
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
		// Boilerplate //
		// Command init
		initCmdOpts(cmd, cmdOpts)
		// Client init
		initCfgOpts(cmd, cfgOpts)
		client, err := genClientWOpts(cfgOpts)
		if err != nil {
			log.Fatal(err)
		}
		// Everything else
		err = watcher.WatcherServiceStart(client, cmdOpts)
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
		// Boilerplate //
		// Command init
		initCmdOpts(cmd, cmdOpts)
		// Client init
		initCfgOpts(cmd, cfgOpts)
		client, err := genClientWOpts(cfgOpts)
		if err != nil {
			log.Fatal(err)
		}
		// Everything else
		err = watcher.WatcherList(client, cmdOpts)
		if err != nil {
			log.Fatal(err)
		}
	},
}
var watcherShowActive = &cobra.Command{
	Use:   "show-active",
	Short: "Show active watchers",
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate //
		// Command init
		initCmdOpts(cmd, cmdOpts)
		// Client init
		initCfgOpts(cmd, cfgOpts)
		client, err := genClientWOpts(cfgOpts)
		if err != nil {
			log.Fatal(err)
		}
		// Everything else
		err = watcher.WatcherShowActive(client, cmdOpts)
		if err != nil {
			log.Fatal(err)
		}
	},
}
var watcherShowInactive = &cobra.Command{
	Use:   "show-inactive",
	Short: "Show inactive watchers",
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate //
		// Command init
		initCmdOpts(cmd, cmdOpts)
		// Client init
		initCfgOpts(cmd, cfgOpts)
		client, err := genClientWOpts(cfgOpts)
		if err != nil {
			log.Fatal(err)
		}
		// Everything else
		err = watcher.WatcherShowInactive(client, cmdOpts)
		if err != nil {
			log.Fatal(err)
		}
	},
}
