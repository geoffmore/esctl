package cmd

import (
	"github.com/geoffmore/esctl/pkg/opts"
	"github.com/spf13/cobra"
	"strings"
)

// Initialize options passed to commands
func initCmdOpts(cmd *cobra.Command, c *opts.CommandOptions, args []string) error {
	// Persistent
	_ = c.SetArgs(args)
	_ = c.SetOutputFormat(outputFmt)
	c.SetVerbose(verbose)
	// Transient
	// TODO - remove input-file in favor of bodyFile
	f, err := cmd.Flags().GetString("input-file")
	if err == nil {
		_ = c.SetInputFile(f)
	}
	initInactive, err := cmd.Flags().GetBool("initInactive")
	if err == nil {
		c.SetWatcherInitInactive(initInactive)
	}

	err = c.SetBody(bodyFile)

	return err
}

// Initialize options used for client generation
func initCfgOpts(cmd *cobra.Command, c *opts.ConfigOptions) error {
	// Not sure how cmd should be used here
	_ = cmd

	// Persistent
	c.SetDebug(debug)
	c.SetConfigFile(cfgFile)
	c.SetContext(context)
	// Transient
	return nil
}

func initCredOpts(cmd *cobra.Command, c *opts.CredentialOptions) error {
	// Not sure how cmd should be used here
	_ = cmd

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
