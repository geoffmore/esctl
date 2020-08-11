package opts

type ConfigOptions struct {
	ConfigFile string
	Context    string
	Debug      bool
}

type CommandOptions struct {
	InputFile           string
	WatcherInitInactive bool
	OutputFormat        string
	Verbose             bool

	// pretty and human aren't really configurable today. The current
	// assumption is that a human consumes the output by default, so these are
	// set true in functions
	//defaultPretty              = true
	//defaultHuman               = true
}
