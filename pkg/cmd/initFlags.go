package cmd

import (
	"github.com/geoffmore/esctl/pkg/opts"
	"github.com/spf13/cobra"
)

// A struct method is needed for these init functions to access
// the Flags under each command

func initPersistentCmdOpts(cmd *cobra.Command, c *opts.CommandOptions) error {
	c.SetOutputFormat(outputFmt)
	c.SetVerbose(verbose)

	return nil
}

func initCmdOpts(cmd *cobra.Command, c *opts.CommandOptions) error {
	//func initCmdOpts(c *opts.CommandOptions) error {
	f, err := cmd.Flags().GetString("input-file")
	if err == nil {
		c.SetInputFile(f)
	}
	return err
}

func initPersistentCfgOpts(cmd *cobra.Command, c *opts.ConfigOptions) error {
	return nil
}

func initCfgOpts(cmd *cobra.Command, c *opts.ConfigOptions) error {
	c.SetDebug(debug)
	return nil
}
