package opts

func NewConfigOption() *ConfigOption {
	const (
		defaultConfigFile = "$HOME/.elastic/config"
		defaultContext    = ""
		defaultDebug      = false
		// verbose flag should exist in the future
	)

	c := &ConfigOption{
		ConfigFile: defaultConfigFile,
		Context:    defaultContext,
		Debug:      defaultDebug,
	}

	return c
}

func NewCommandOption() *CommandOption {

	const (
		defaultInputFile           = ""
		defaultWatcherInitInactive = false
		defaultOutputFormat        = ""
	)

	c := &CommandOption{
		InputFile:           defaultInputFile,
		WatcherInitInactive: defaultWatcherInitInactive,
		OutputFormat:        "",
	}

	return c
}
