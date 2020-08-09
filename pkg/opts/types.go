package opts

type ConfigOptions struct {
	ConfigFile string
	Context    string
	Debug      bool
	// verbose flag should exist in the future
	// Verbose bool
}
type CommandOptions struct {
	InputFile           string
	WatcherInitInactive bool
	OutputFormat        string
	// pretty and human aren't really configurable today. The current
	// assumption is that a human consumes the output by default, so these are
	// set true in functions
	//defaultPretty              = true
	//defaultHuman               = true
}
