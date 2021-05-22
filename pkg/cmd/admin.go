package cmd

import (
	"errors"
	"github.com/geoffmore/esctl/pkg/admin"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

// verbCmd should probably have its own file that can be referenced in places.
// Maybe esutil?
type verbCmd struct {
	Command string
	Aliases []string
}

// TODO - improve boilerplate for all functions in admin package
// TODO - removal optional status of admin package
// TODO - make verbCmd part of esutil and figure out a structure for commands at various levels
var listCmd = verbCmd{Command: "list", Aliases: []string{"ls"}}

//var showCmd = verbCmd{Command: "show"}

func init() {
	rootCmd.AddCommand(adminCmd)
	adminCmd.AddCommand(adminNode)
	adminNode.AddCommand(adminNodeList)
	adminCmd.AddCommand(listNodes)
	adminCmd.AddCommand(listNodesStorage)
	// TODO - add showShards function with newest boilerplate
	// adminCmd.AddCommand(showShards)
	adminCmd.AddCommand(showBigShards)
	adminCmd.AddCommand(showSmallShards)
	adminCmd.AddCommand(showShardUsageByNode)
}

var adminAliases = [...]string{"adm"}

// Maybe this could go under a ResourceCmd struct also under esutil?
var nodeAliases = [...]string{"nodes", "no"}

// Temporary until admin commands are refactored
var help = false

// I should add a --help flag that adds the pointer help field to a request. Not
// sure how to wrap that for the admin package. Maybe I don't need to?
var adminCmd = &cobra.Command{
	// esctl admin
	Use:     "admin",
	Aliases: adminAliases[:],
	Short:   "Commands useful for Elasticsearch operators",
	Long:    "Commands useful for Elasticsearch operators.\nMuch of this is inspired by https://github.com/slmingol/escli",
}

var adminNode = &cobra.Command{
	Use:     "node",
	Aliases: nodeAliases[:],
	Short:   "node stuff",
}

var adminNodeList = &cobra.Command{
	Use:     listCmd.Command,
	Aliases: listCmd.Aliases[:],
	Short:   "list nodes",
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate //
		// Command init

		err := initCmdOpts(cmd, cmdOpts, args)
		if err != nil {
			log.Fatal(err)
		}
		// Client init
		err = initCfgOpts(cmd, cfgOpts)
		if err != nil {
			log.Fatal(err)
		}
		client, err := genClientWOpts(cfgOpts)
		if err != nil {
			log.Fatal(err)
		}
		// Everything else
		err = admin.NodeList(client, cmdOpts)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var listNodes = &cobra.Command{
	// esctl admin list-nodes
	Use:   "list-nodes",
	Short: "No description",
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate
		client, err := genClient(context)
		if err != nil {
			log.Fatal(err)
		}

		err = admin.ListNodes(client, outputFmt, help)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var listNodesStorage = &cobra.Command{
	// esctl admin list-nodes
	Use:   "list-nodes-storage",
	Short: "No description",
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate
		client, err := genClient(context)
		if err != nil {
			log.Fatal(err)
		}

		err = admin.ListNodesStorage(client, outputFmt, help)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var showBigShards = &cobra.Command{
	// esctl admin list-nodes
	Use:   "show-big-shards",
	Short: "No description",
	// first arg int
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires an integer argument")
		}
		i, err := strconv.Atoi(args[0])
		if err != nil {
			return errors.New("provided argument was not a an integer")
		}
		if i < 1 {
			return errors.New("provided integer needs to be greater than 0")
		}
		return nil

	},
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate
		client, err := genClient(context)
		if err != nil {
			log.Fatal(err)
		}

		// Error checking is done in Args, so there is no need to duplicate logic
		// in Run
		i, _ := strconv.Atoi(args[0])
		err = admin.ShowBigShards(client, outputFmt, help, i)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var showSmallShards = &cobra.Command{
	// esctl admin list-nodes
	Use:   "show-small-shards",
	Short: "No description",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires an integer argument")
		}
		i, err := strconv.Atoi(args[0])
		if err != nil {
			return errors.New("provided argument was not a an integer")
		}
		if i < 1 {
			return errors.New("provided integer needs to be greater than 0")
		}
		return nil

	},
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate
		client, err := genClient(context)
		if err != nil {
			log.Fatal(err)
		}

		// Error checking is done in Args, so there is no need to duplicate logic
		// in Run
		i, _ := strconv.Atoi(args[0])
		err = admin.ShowSmallShards(client, outputFmt, help, i)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var showShardUsageByNode = &cobra.Command{
	// esctl admin show-shard-usage-by-node
	Use:   "show-shard-usage-by-node",
	Short: "No description",
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate
		client, err := genClient(context)
		if err != nil {
			log.Fatal(err)
		}

		err = admin.ShowShardUsageByNode(client, outputFmt)
		if err != nil {
			log.Fatal(err)
		}
	},
}
