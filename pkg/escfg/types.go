package escfg

// https://godoc.org/github.com/elastic/go-elasticsearch#Config
// Inspiration from the above. Ideally I only define addresses and username and
// optionally cloudid, then use that to generate a api key as part of
// ~/.elastic/config

const (
	// Default configuration file name
	DefaultElasticConfigName string = "config"
	// Default configuration file path
	DefaultElasticFolderPath string = "$HOME/.elastic"
	// Default full path to configuration file
	DefaultElasticConfig string = "$HOME/.elastic/config"
	// Default configuration folder mode
	DefaultElasticFolderMode uint32 = 755
	// Default configuration file mode
	DefaultElasticConfigMode uint32 = 0600
)

// Default Config struct used in GenDefaultConfig.
// Connects to a local elasticsearch cluster
var DefaultConfig = Config{
	Clusters: []Cluster{
		Cluster{
			Name: "local",
			ElasticAddresses: []string{
				"http://localhost:9200",
				"http://127.0.0.1:9200",
			},
			AllowSelfSigned: "yes",
		},
	},
	Contexts: []Contexts{
		Contexts{
			Name: "local",
			Context: Context{
				Cluster: "local",
				User:    "elastic",
			},
		},
		Contexts{
			Name: "elastic@local",
			Context: Context{
				Cluster: "local",
				User:    "elastic@local",
			},
		},
	},
	Users: []Users{
		Users{
			Name: "elastic@local",
			User: User{
				Name:     "elastic",
				Password: "",
				ApiKey:   "",
				Token: Token{
					Value:      "",
					Expiration: "",
				},
			},
		},
		Users{
			Name: "elastic",
			User: User{
				Name:     "elastic",
				Password: "changeme",
				ApiKey:   "",
				Token: Token{
					Value:      "",
					Expiration: "",
				},
			},
		},
	},
	CurrentContext: "local",
}

// File configuration
type Config struct {
	Clusters       []Cluster  `yaml:"clusters"`
	Contexts       []Contexts `yaml:"contexts"`
	Users          []Users    `yaml:"users"`
	CurrentContext string     `yaml:"current-context,omitempty"`
}

// Connection details for a cluster
type Cluster struct {
	Name             string   `yaml:"name"`
	ElasticAddresses []string `yaml:"elastic-addresses,omitempty"`
	//KibanaAddresses  []string `yaml:"kibana-addresses",omitempty`
	CloudID         string `yaml:"cloud-id,omitempty"`
	AllowSelfSigned string `yaml:"allowSelfSigned,omitempty"`
}

// Check if all fields in the Cluster struct have their nil value. Necessary
// because arrays nested under structs can't be compared
func (c Cluster) IsNil() bool {
	if c.Name != "" {
		return false
	}
	if len(c.ElasticAddresses) != 0 {
		return false
	}
	// I don't care about KibanaAddresses and Addresses yet
	//if len(c.KibanaAddresses) != 0 {
	//	return false
	//}
	if c.CloudID != "" {
		return false
	}
	return true
}

// Named Context
type Contexts struct {
	Name    string  `yaml:"name,omitempty"`
	Context Context `yaml:"context,omitempty"`
}

// A combination of a named Cluster and named User. Used for lookups of the
// respective objects
type Context struct {
	Cluster string `yaml:"cluster,omitempty"`
	User    string `yaml:"user,omitempty"`
}

// Named User
type Users struct {
	Name string `yaml:"name,omitempty"`
	User User   `yaml:"user,omitempty"`
}

// Authentication credentials
type User struct {
	Name        string    `yaml:"name,omitempty"`
	NameCmd     ConfigCmd `yaml:"name-cmd,omitempty"`
	Password    string    `yaml:"password,omitempty"`
	PasswordCmd ConfigCmd `yaml:"password-cmd,omitempty"`
	ApiKey      string    `yaml:"api-key,omitempty"`
	Token       Token     `yaml:"token,omitempty"`
}

type ConfigCmd struct {
	// A ConfigCmd should at least have a Command
	Command string   `yaml:"command"`
	Env     []string `yaml:"env,omitempty"`
	Args    []string `yaml:"args,omitempty"`
}

// Check if ConfigCmd is empty
func (c ConfigCmd) IsEmpty() bool {
	return (c.Command == "") && (len(c.Env) == 0) && (len(c.Args) == 0)
}

// Token used for authentication
type Token struct {
	Value      string `yaml:"value,omitempty"`
	Expiration string `yaml:"expiration,omitempty"`
}
