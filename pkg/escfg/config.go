package escfg

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/geoffmore/esctl/pkg/esauth"
	"github.com/geoffmore/esctl/pkg/esutil"
	"github.com/geoffmore/esctl/pkg/opts"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

// Config functions

// Generate a configuration file for connecting to a local elasticsearch cluster
// if a file is not found by path
func ConfigGenDefaultConfig(cfgOpts *opts.ConfigOptions) (err error) {

	file := cfgOpts.ConfigFile

	if !exists(file) {
		fmt.Printf("Default config file not found. Creating...\n")
		err = DefaultConfig.write(file)
	} else {
		fmt.Printf("Config file found, refusing to overwrite...\n")
	}
	return err
}

// Marshal a Config into bytes
func ConfigShow(cfgOpts *opts.ConfigOptions) error {
	file := os.Expand(cfgOpts.ConfigFile, os.Getenv)
	cfg, err := ReadConfig(file)
	if err != nil {
		return err
	}

	b, err := cfg.show()
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func ConfigTest(cfgOpts *opts.ConfigOptions) error {
	var cfg Config

	file := os.Expand(cfgOpts.ConfigFile, os.Getenv)

	b, err := ToBytes(file)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(b, &cfg)
	if err != nil {
		return fmt.Errorf("Config file is invalid")
	}

	fmt.Println("Config file is valid")
	return nil
}

// Config methods

func (c Config) hasCluster(name string) bool {
	var contains bool
	for _, cluster := range c.Clusters {
		if cluster.Name == name {
			contains = true
		}
	}
	return contains
}

func (c Config) hasContext(name string) bool {
	var contains bool
	for _, context := range c.Contexts {
		if context.Name == name {
			contains = true
		}
	}
	return contains
}

func (c Config) hasUser(name string) bool {
	var contains bool
	for _, user := range c.Users {
		if user.Name == name {
			contains = true
		}
	}
	return contains
}

func (c Config) getContexts() []string {
	var contexts = make([]string, len(c.Contexts))

	for i, context := range c.Contexts {
		contexts[i] = context.Name
	}
	return contexts
}

func (c Config) getContext(name string) (Context, error) {
	var context Context
	if !c.hasContext(name) {
		return context, fmt.Errorf("Context '%s' not found\n", name)
	}

	for _, currentContext := range c.Contexts {
		if name == currentContext.Name {
			context = currentContext.Context
		}
	}
	return context, nil
}

func (c Config) show() ([]byte, error) {
	b, err := yaml.Marshal(c)
	if err != nil {
		err = fmt.Errorf("configMarshalError")
	}
	return b, err
}

// Write a Config into a file
func (c Config) write(file string) error {
	var err error
	var b []byte

	b, err = c.show()
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

// Does there really need to be two return values?

func (c Config) testContext(context string, debug bool) (bool, error) {
	// Initialize client. GenClient expects the file to exist, so we must work
	// around it
	esConfig, err := GenESConfig(c, context, debug)
	if err != nil {
		return false, err
	}
	client, err := esauth.EsAuth(esConfig)
	if err != nil {
		return false, err
	}

	req := esapi.InfoRequest{}
	res, err := esutil.GetResponse(req, client)
	if err != nil {
		return false, err
	}

	// res.Status() returns a string "200 OK" if successful. It is safer to use
	// the int value
	if statusCode := res.StatusCode; statusCode != 200 {
		return false, fmt.Errorf("Unable to authenticate to cluster. Returned status code is '%d'\n", statusCode)
	} else {
		return true, nil
	}
}

func (cfg1 Config) merge(cfg2 Config) (Config, error) {
	var cfg3 Config = cfg1
	var conflicts string = fmt.Sprintf("Unable to completely merge objects. Key collisions found. Collisions: \n")

	var hasConflict bool
	var clusterConflicts, contextConflicts, userConflicts []string
	//cfg2 takes priority, but does not overwrite objects with the same name
	// CurrentContext
	cfg3.CurrentContext = cfg2.CurrentContext
	// Clusters
	for _, cluster := range cfg2.Clusters {
		if !cfg3.hasCluster(cluster.Name) {
			cfg3.Clusters = append(cfg3.Clusters, cluster)
		} else {
			hasConflict = true
			clusterConflicts = append(clusterConflicts, cluster.Name)
		}
	}
	// Contexts
	for _, context := range cfg2.Contexts {
		if !cfg3.hasContext(context.Name) {
			cfg3.Contexts = append(cfg3.Contexts, context)
		} else {
			hasConflict = true
			contextConflicts = append(contextConflicts, context.Name)
		}
	}
	// Users
	for _, user := range cfg2.Users {
		if !cfg3.hasUser(user.Name) {
			cfg3.Users = append(cfg3.Users, user)
		} else {
			hasConflict = true
			userConflicts = append(userConflicts, user.Name)
		}
	}

	if hasConflict {
		if len(clusterConflicts) > 0 {
			conflicts = conflicts + fmt.Sprintf("Clusters: %v\n", clusterConflicts)
		}
		if len(contextConflicts) > 0 {
			conflicts = conflicts + fmt.Sprintf("Contexts: %v\n", contextConflicts)
		}
		if len(userConflicts) > 0 {
			conflicts = conflicts + fmt.Sprintf("Users: %v\n", userConflicts)
		}
		return Config{}, fmt.Errorf("%s", conflicts)
	}

	return cfg3, nil
}
