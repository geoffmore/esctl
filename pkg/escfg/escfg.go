package escfg

import (
	"crypto/tls"
	"errors"
	"fmt"
	elastic7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/estransport"
	"golang.org/x/crypto/ssh/terminal"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"os"
)

// Read a file into bytes
func ToBytes(file string) (b []byte, err error) {
	b, err = ioutil.ReadFile(file)
	if err != nil {
		err = fmt.Errorf("File \"%s\" not found\n", file)
	}
	return b, err
}

// Unmarshal bytes into a Config
func ReadConfig(file string) (cfg Config, err error) {
	var dat []byte

	dat, err = ToBytes(file)
	//dat, err = ioutil.ReadFile(file)
	if err != nil {
		//fmt.Errorf("File \"%s\" not found\n", file)
		return cfg, err
	}

	err = yaml.Unmarshal(dat, &cfg)
	if err != nil {
		fmt.Errorf("Invalid config format for file \"%s\"\n", file)
		return cfg, err
	}

	return cfg, err
}

// Marshal a Config into bytes
func DisplayConfig(cfg Config) ([]byte, error) {
	var b []byte
	var err error
	b, err = yaml.Marshal(cfg)
	if err != nil {
		err = errors.New("configMarshalError")
		return b, err
	}
	return b, nil
}

// Write a Config into a file
func writeConfig(cfg Config, file string) error {

	var err error
	var b []byte

	b, err = DisplayConfig(cfg)
	if err != nil {
		err = fmt.Errorf("unable to marshal config")
		return err
	}
	err = ioutil.WriteFile(file, b, os.FileMode(DefaultElasticConfigMode))
	if err != nil {
		return err
	}
	return nil

}

// Check if a file exists
func exists(name string) bool {
	// https://stackoverflow.com/questions/12518876/how-to-check-if-a-file-exists-in-go
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// Generate a configuration file for connecting to a local elasticsearch cluster
// if a file is not found by path
func GenDefaultConfig(file string) (err error) {

	if !exists(file) {
		fmt.Printf("Default config file not found. Creating...\n")
		err = writeConfig(DefaultConfig, file)
	} else {
		fmt.Printf("Config file found, refusing to overwrite...\n")
	}
	return err
}

// Return a list of contexts
func GetContexts(cfg Config) []string {
	var contexts []string
	for _, context := range cfg.Contexts {
		contexts = append(contexts, context.Name)
	}
	return contexts
}

// Change the value of Config.CurrentContext
func UseContext(n string, cfg Config, file string) error {
	var err error

	contexts := GetContexts(cfg)
	// Test to ensure input is not equal to current-context
	if cfg.CurrentContext == n {
		err = fmt.Errorf("Provided context matches current value")
		return err
	}

	// Ensure provided context exists
	var ctxInCtxs bool = false
	for _, ctx := range contexts {
		if ctx == n {
			ctxInCtxs = true
		}
	}

	if !ctxInCtxs {
		err = fmt.Errorf("Context \"%s\" not found", n)
		return err
	}
	// Change the value of current-context to provided context
	cfg.CurrentContext = n
	var b []byte
	b, err = DisplayConfig(cfg)
	if err != nil {
		err = fmt.Errorf("unable to marshal config")
		return err
	}
	err = ioutil.WriteFile(file, b, os.FileMode(DefaultElasticConfigMode))
	if err != nil {
		return err
	}

	return nil
}

// Test a named context
func TestContext(n string) {}

// Test if a config file is valid by attempting to unmarshal into a Config
// struct
func IsValidCfg(b []byte) bool {
	var cfg Config
	err := yaml.Unmarshal(b, &cfg)
	if err != nil {
		return false
	}
	return true

}

// Prompt for password to authenticate a request
func askPass() (str string, err error) {
	fmt.Printf("Enter your password to connect as user: ")
	b, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	str = string(b)
	fmt.Printf("\n\n")
	return str, err
}

func GenESConfig(cfg Config, ctx string, debug bool) (es7cfg elastic7.Config, err error) {
	// Order of operations
	// --- Cluster ---
	// Ordered by completeness of information
	//   CloudID
	//   ElasticAddresses
	//	  KibanaAddresses []string `yaml:"kibana-addresses",omitempty`
	// --- User ---
	// Ordered by interactiveness, then credential longevity
	//	  Token
	//   ApiKey
	//	  Password + Name
	//   Name

	// Use current-context to get relevant structs
	// Need User, Cluster
	var currentUser User
	var currentCluster Cluster
	var currentContext Context

	var currentUserName string
	var currentClusterName string
	var currentContextName string

	// Get the current context
	// Edit: try provided ctx variable
	if ctx != "" {
		currentContextName = ctx
	} else {
		currentContextName = cfg.CurrentContext
	}

	if currentContextName == "" {
		err = fmt.Errorf("Current context not defined")
		return es7cfg, err
	}
	// Use the current context to lookup names
	for _, context := range cfg.Contexts {
		if currentContextName == context.Name {
			currentContext = context.Context
		}
	}
	// How can I catch the context not being found in cfg.Contexts?
	// https://stackoverflow.com/questions/28447297/how-to-check-for-an-empty-struct
	if currentContext == (Context{}) {
		err = fmt.Errorf("Context %s not found", currentContextName)
		return es7cfg, err
	}

	// Get a User struct to work with
	currentUserName = currentContext.User
	if currentUserName == "" {
		err = fmt.Errorf("Current user not defined")
		return es7cfg, err
	}
	for _, user := range cfg.Users {
		if currentUserName == user.Name {
			currentUser = user.User
		}
	}
	if currentUser == (User{}) {
		err = fmt.Errorf("User %s not found", currentUserName)
		return es7cfg, err
	}

	// Get a cluster struct to work with
	currentClusterName = currentContext.Cluster
	if currentClusterName == "" {
		err = fmt.Errorf("Current cluster not defined")
		return es7cfg, err
	}
	for _, cluster := range cfg.Clusters {
		if currentClusterName == cluster.Name {
			currentCluster = cluster
		}
	}
	if currentCluster.IsNil() {
		err = fmt.Errorf("Cluster %s not found", currentClusterName)
		return es7cfg, err
	}

	// Create connection information
	if currentCluster.CloudID != "" {
		es7cfg.CloudID = currentCluster.CloudID
	} else if len(currentCluster.ElasticAddresses) != 0 {
		es7cfg.Addresses = currentCluster.ElasticAddresses
	} else {
		err = fmt.Errorf("Neither CloudID nor ElasticAddresses field populated. Unable to generate es7cfg.")
		return es7cfg, err
	}

	// Create user information
	if currentUser.ApiKey != "" {
		es7cfg.APIKey = currentUser.ApiKey
	} else if currentUser.Name != "" && currentUser.Password != "" {
		es7cfg.Username = currentUser.Name
		es7cfg.Password = currentUser.Password
	} else if currentUser.Name != "" {
		es7cfg.Username = currentUser.Name
		pass, err := askPass()
		if err != nil {
			return es7cfg, err
		}
		es7cfg.Password = pass
	} else {
		err = fmt.Errorf("None of token, apikey, name + password, or name provided.\n Unable to generate config")
		return es7cfg, err
	}

	//https://stackoverflow.com/questions/37557763
	if currentCluster.AllowSelfSigned == "yes" {
		transport := http.DefaultTransport
		tlsClientConfig := &tls.Config{InsecureSkipVerify: true}
		transport.(*http.Transport).TLSClientConfig = tlsClientConfig
		es7cfg.Transport = transport
	}

	// Debug connection stuff. Should be wrapped in a feature flag
	// There are a lot of debugging options. This will likely need to be extended
	// in the future.
	// https://godoc.org/github.com/elastic/go-elasticsearch/estransport#ColorLogger
	if debug {
		es7cfg.Logger = &estransport.ColorLogger{
			Output:            os.Stdout,
			EnableRequestBody: true,
			// Response body is not needed since that is already returned via
			// esutil
			EnableResponseBody: false,
		}
		es7cfg.EnableMetrics = true
		es7cfg.EnableDebugLogger = true
	}

	return es7cfg, err
}
