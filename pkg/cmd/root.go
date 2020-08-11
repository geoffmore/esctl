package cmd

import (
	"fmt"
	elastic7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/geoffmore/esctl/pkg/esauth"
	"github.com/geoffmore/esctl/pkg/escfg"
	"github.com/geoffmore/esctl/pkg/opts"
	"github.com/geoffmore/esctl/pkg/version"
	"github.com/spf13/cobra"
	"os"
)

// file defined within the package scope for reusability
// Config file path
var file string = os.Expand(escfg.DefaultElasticConfig, os.Getenv)

// Initialize defaults
var cfgOpts = opts.NewConfigOptions()
var cmdOpts = opts.NewCommandOptions()

var (
	// Used for flags
	outputFmt string
	context   string
	verbose   bool
	debug     bool
	cfgFile   string
)

var rootCmd = &cobra.Command{
	Use:     "esctl",
	Short:   "esctl is a utility able to interact with elasticsearch clusters",
	Version: version.Version,
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&outputFmt, "output", "o", "", "choice of output format")
	rootCmd.PersistentFlags().StringVarP(&context, "context", "c", "", "choice of context to use for a command")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "", false, "debug connection")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "increase the verbosity of certain api endpoint repsonses")
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "", "", "choice of context to use for a command")

	// Note: Non-persistent flags do not appear to be inheritable
	// See https://github.com/spf13/cobra/issues/747 for context on
	// flag inheritance
}

// Main function of cobra
func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Generates an elasticsearch client from a named file from start to finish
func genClient(ctx string) (client *elastic7.Client, err error) {

	//file := os.Expand(escfg.DefaultElasticConfig, os.Getenv)
	fileConfig, err := escfg.ReadConfig(file)
	if err != nil {
		return client, err
	}
	esConfig, err := escfg.GenESConfig(fileConfig, ctx, debug)
	if err != nil {
		return client, err
	}
	client, err = esauth.EsAuth(esConfig)
	if err != nil {
		return client, err
	}
	return client, err
}
