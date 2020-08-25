package cmd

import (
	"github.com/geoffmore/esctl/pkg/opts"
	"github.com/spf13/cobra"
	"strings"
)

// Initialize options passed to commands
func initCmdOpts(cmd *cobra.Command, c *opts.CommandOptions, args []string) error {
	// Persistent
	c.SetArgs(args)
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

func initCredOpts(cmd *cobra.Command, c *opts.CredentialOptions) error {
	c.SetAddresses(strings.Split(addresses, ","))
	c.SetUser(user)
	c.SetPassword(pass)
	c.SetCloudID(cloudID)
	c.SetAPIKey(apiKey)
	c.SetToken(token)
	if insecure {
		c.SetInsecure()
	}
	// Not yet implemented
	//c.SetCertFile(certFile)
	//c.SetCert(cert)

	return nil
}
