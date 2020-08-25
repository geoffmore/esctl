package opts

import "fmt"

func NewConfigOptions() *ConfigOptions {
	const (
		// Low to high priority
		// How do I handle the case where the provided string is different file
		// name is different than what was expected?
		// ex ~/.elastic/foo instead of ~/.elastic/config
		defaultConfigFile = "$HOME/.elastic/config"
		defaultContext    = ""
		defaultDebug      = false
		// verbose flag should exist in the future
	)

	c := &ConfigOptions{
		ConfigFile: defaultConfigFile,
		Context:    defaultContext,
		Debug:      defaultDebug,
	}

	return c
}

func NewCommandOptions() *CommandOptions {

	const (
		defaultInputFile           = ""
		defaultWatcherInitInactive = false
		defaultOutputFormat        = ""
	)

	c := &CommandOptions{
		InputFile:           defaultInputFile,
		WatcherInitInactive: defaultWatcherInitInactive,
		OutputFormat:        defaultOutputFormat,
	}

	return c
}

func NewCredentialOptions() *CredentialOptions {
	c := &CredentialOptions{
		User:     "",
		Password: "",
		APIKey:   "",
		Token:    "",
	}
	return c
}

func (c *CommandOptions) SetOutputFormat(a string) error {
	// https://www.elastic.co/guide/en/elasticsearch/reference/7.8/common-options.html#_content_type_requirements
	var validFormats = [...]string{
		"json",
		"yaml",
		"cbor",
		"smile",
	}

	var isValidFormat bool
	for _, b := range validFormats {

		if b == a {
			isValidFormat = true
		}
	}

	if isValidFormat {
		c.OutputFormat = a
		return nil
	} else {
		return fmt.Errorf("Invalid output format '%s' provided. Please use one of %v\n", a, validFormats)
	}
}

func (c *CommandOptions) SetArgs(args []string) error {
	c.Args = make([]string, len(args))
	c.Args = args
	return nil
}
func (c *CommandOptions) SetInputFile(a string) error {
	c.InputFile = a
	return nil
}

func (c *CommandOptions) SetWatcherInitInactive(a bool) {
	if a {
		c.WatcherInitInactive = true
	}
}

func (c *CommandOptions) SetVerbose(a bool) {
	if a {
		c.Verbose = true
	}
}

func (c *ConfigOptions) SetDebug(a bool) {
	if a {
		c.Debug = true
	}
}

func (c *ConfigOptions) SetConfigFile(a string) {
	if a != "" {
		c.ConfigFile = a
	}
}

func (c *ConfigOptions) SetContext(a string) {
	if a != "" {
		c.Context = a
	}
}

func (c *CredentialOptions) SetAddresses(a []string) {
	l := len(a)
	if l > 0 {
		c.Addresses = make([]string, l)
		c.Addresses = a
	}
}

func (c *CredentialOptions) SetUser(a string) {
	if a != "" {
		c.User = a

	}
}

func (c *CredentialOptions) SetPassword(a string) {
	if a != "" {
		c.Password = a
	}
}

func (c *CredentialOptions) SetAPIKey(a string) {
	if a != "" {
		c.APIKey = a
	}
}

func (c *CredentialOptions) SetToken(a string) {
	if a != "" {
		c.Token = a
	}
}

func (c *CredentialOptions) SetInsecure() {
	c.Insecure = true
}

func (c *CredentialOptions) SetCloudID(a string) {
	if a != "" {
		c.CloudID = a
	}
}
