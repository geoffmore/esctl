package cmd

import (
	"github.com/geoffmore/esctl/pkg/opts"
	"github.com/spf13/cobra"
)

// Initialize options passed to commands
func initCmdOpts(cmd *cobra.Command, c *opts.CommandOptions) error {
	// Persistent
	c.SetOutputFormat(outputFmt)
	c.SetVerbose(verbose)
	// Transient
	f, err := cmd.Flags().GetString("input-file")
	if err == nil {
		c.SetInputFile(f)
	}
	initInactive, err := cmd.Flags().GetBool("initInactive")
	if err == nil {
		c.SetWatcherInitInactive(initInactive)
	}

	return nil
}

// Initialize options used for client generation
func initCfgOpts(cmd *cobra.Command, c *opts.ConfigOptions) error {
	// Persistent
	c.SetDebug(debug)
	c.SetConfigFile(cfgFile)
	c.SetContext(context)
	// Transient
	return nil
}
