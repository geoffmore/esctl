package opts

import "io"

type CredentialOptions struct {
	Addresses []string
	User      string
	Password  string
	APIKey    string
	Token     string
	Insecure  bool
	CloudID   string
}
type ConfigOptions struct {
	ConfigFile string
	Context    string
	Debug      bool
}

type CommandOptions struct {
	// Persistent
	Args         []string
	OutputFormat string
	Verbose      bool
	// Non-persistent
	InputFile           string
	Body                io.Reader
	WatcherInitInactive bool

	// pretty and human aren't really configurable today. The current
	// assumption is that a human consumes the output by default, so these are
	// set true in functions
	//defaultPretty              = true
	//defaultHuman               = true
}

// Map of CommandOptions to field names
var CmdsToFieldNames = map[string]string{
	"OutputFormat": "Format",
	"Verbose":      "V",
	//"InputFile": "",
	"WatcherInitInactive": "Active",
}
