package cmd

import (
	"github.com/geoffmore/esctl/pkg/escfg"
	"github.com/spf13/cobra"
	"log"
)

// Variables needed for initCredOpts
var addresses, user, pass, token, apiKey, cloudID string
var insecure bool

// ES_CONFIG variable should be a thing eventually

func init() {
	rootCmd.AddCommand(configCmd)
	// Aligns with other "esctl <resource> <function>" syntax. Refactoring to
	// form "esctl <function> <resource> is a larger conversation
	// This feels a bit backward to the "esctl <resource> <function>" method of
	// arranging functions, but k

	// Context functions
	configCmd.AddCommand(configContextCmd)
	// This really needs to be a FlagSet
	configContextCmd.AddCommand(configContextGen)
	configContextGen.Flags().StringVarP(&addresses, "addresses", "", "", "comma-delimeted list of addresses used to attempt a connection")
	configContextGen.Flags().StringVarP(&cloudID, "cloud-id", "", "", "cloud-id to connect to")
	configContextGen.Flags().StringVarP(&user, "user", "", "", "username to connect with")
	configContextGen.Flags().StringVarP(&pass, "password", "", "", "password to connect with")
	configContextGen.Flags().StringVarP(&token, "token", "", "", "token to connect with")
	configContextGen.Flags().StringVarP(&apiKey, "api-key", "", "", "api-key to connect with")
	configContextGen.Flags().BoolVarP(&insecure, "insecure", "", false, "whether or not to skip tls verification")

	configContextCmd.AddCommand(configContextList)
	configContextCmd.AddCommand(configContextGetCurrent)
	configContextCmd.AddCommand(configContextShow)
	configContextCmd.AddCommand(configContextTest)
	configContextCmd.AddCommand(configContextUse)
	// Config functions
	configCmd.AddCommand(configGenDefault)
	configCmd.AddCommand(configShow)
	configCmd.AddCommand(configTest)

}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Interact with a config file",
}

// Context functions
var configContextCmd = &cobra.Command{
	Use:   "context",
	Short: "Work with config file contexts",
}

var configContextGen = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"gen"},
	Short:   "Generate a context based on provided connection details",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate //
		// Command flag init (only for args)
		initCmdOpts(cmd, cmdOpts, args)
		// Config flag init
		initCfgOpts(cmd, cfgOpts)
		// Credential init
		initCredOpts(cmd, credOpts)

		// Everything else
		err := escfg.ConfigContextGen(cfgOpts, cmdOpts, credOpts)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var configContextList = &cobra.Command{
	Use:   "list",
	Short: "List available contexts",
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate //
		initCfgOpts(cmd, cfgOpts)

		// Everything else
		err := escfg.ConfigContextList(cfgOpts)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var configContextGetCurrent = &cobra.Command{
	Use:   "get-current",
	Short: "Get the current context",
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate //
		// Command init
		initCmdOpts(cmd, cmdOpts, args)
		// Client init
		initCfgOpts(cmd, cfgOpts)

		// Everything else
		err := escfg.ConfigContextGetCurrent(cfgOpts)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var configContextShow = &cobra.Command{
	Use:   "show",
	Short: "Show the config for the desired context. Defaults to current context",
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate //
		// Command init
		initCmdOpts(cmd, cmdOpts, args)
		// Client init
		initCfgOpts(cmd, cfgOpts)

		// Everything else
		err := escfg.ConfigContextShow(cfgOpts, cmdOpts)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var configContextTest = &cobra.Command{
	Use:   "test",
	Short: "Test the desired context for connectivity. Defaults to current context",
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate //
		// Command init
		initCmdOpts(cmd, cmdOpts, args)
		// Client init
		initCfgOpts(cmd, cfgOpts)

		// Everything else
		err := escfg.ConfigContextTest(cfgOpts, cmdOpts)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var configContextUse = &cobra.Command{
	Use:   "use",
	Short: "Use the desired context. Defaults to current context",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate //
		// Command init
		initCmdOpts(cmd, cmdOpts, args)
		// Client init
		initCfgOpts(cmd, cfgOpts)

		// Everything else
		err := escfg.ConfigContextTest(cfgOpts, cmdOpts)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// Config functions

var configGenDefault = &cobra.Command{
	Use:   "gen-default",
	Short: "generate the default config file",
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate //
		// Client init
		initCfgOpts(cmd, cfgOpts)
		// Everything else
		err := escfg.ConfigGenDefaultConfig(cfgOpts)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var configShow = &cobra.Command{
	Use:   "show",
	Short: "Show the desired config",
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate //
		initCfgOpts(cmd, cfgOpts)

		// Everything else
		err := escfg.ConfigShow(cfgOpts)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var configTest = &cobra.Command{
	Use:   "test",
	Short: "Test the desired config for syntax errors",
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate //
		initCfgOpts(cmd, cfgOpts)

		// Everything else
		err := escfg.ConfigTest(cfgOpts)
		if err != nil {
			log.Fatal(err)
		}
	},
}
