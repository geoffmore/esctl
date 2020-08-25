package escfg

import (
	"fmt"
	"github.com/geoffmore/esctl/pkg/opts"
	"gopkg.in/yaml.v2"
	"os"
)

// Context functions

func ConfigContextGen(cfgOpts *opts.ConfigOptions, cmdOpts *opts.CommandOptions, credOpts *opts.CredentialOptions) error {

	var oldCfg, newCfg, combinedCfg Config

	// Assign a naming scheme for the config file user and context fields
	var baseContext, fullContext, configUsername string

	// Context will be defined from args, not the context flag
	baseContext = cmdOpts.Args[0]

	if credOpts.User != "" {
		configUsername = fmt.Sprintf("%s@%s", credOpts.User, baseContext)
		fullContext = configUsername
	} else {
		configUsername = fmt.Sprint("%s-user", cfgOpts.Context)
		fullContext = fmt.Sprintf("%s@%s", configUsername, baseContext)
	}

	// ... and the full context will be assigned to cfgOpts
	cfgOpts.SetContext(fullContext)

	newCfg, err := NewConfig(baseContext, fullContext, configUsername, cfgOpts, credOpts)
	if err != nil {
		return err
	}

	// Validate config with api call
	// Note, the cluster's '/' endpoint has the possibility of requiring no auth
	// in which case, the test will still pass
	isValidContext, err := newCfg.testContext(cfgOpts.Context, cfgOpts.Debug)
	if (!isValidContext) || (err != nil) {
		return err
	}

	configFileExists := exists(cfgOpts.ConfigFile)
	if !configFileExists {
		// Config file does not exist, write context as config
		combinedCfg = newCfg
	} else {
		// Config file exists, add context to config
		file := os.Expand(cfgOpts.ConfigFile, os.Getenv)
		oldCfg, err = ReadConfig(file)
		if err != nil {
			return err
		}
		combinedCfg, err = oldCfg.merge(newCfg)
		if err != nil {
			return err
		}
	}

	err = combinedCfg.write(cfgOpts.ConfigFile)

	return err
}

func ConfigContextList(cfgOpts *opts.ConfigOptions) error {
	file := os.Expand(cfgOpts.ConfigFile, os.Getenv)
	cfg, err := ReadConfig(file)
	if err != nil {
		return err
	}
	fmt.Println(cfg.getContexts())
	return nil
}

func ConfigContextGetCurrent(cfgOpts *opts.ConfigOptions) error {
	file := os.Expand(cfgOpts.ConfigFile, os.Getenv)
	cfg, err := ReadConfig(file)
	if err != nil {
		return err
	}

	currentContext := cfg.CurrentContext
	if currentContext == "" {
		fmt.Printf("No value for field \"current-context\"\n")
	} else {
		fmt.Printf("Current context is \"%v\"\n", currentContext)
	}

	return nil
}

func ConfigContextShow(cfgOpts *opts.ConfigOptions, cmdOpts *opts.CommandOptions) error {
	var currentContext string

	file := os.Expand(cfgOpts.ConfigFile, os.Getenv)
	cfg, err := ReadConfig(file)
	if err != nil {
		return err
	}

	if len(cmdOpts.Args) == 0 {
		currentContext = cfg.CurrentContext
	} else {
		currentContext = cmdOpts.Args[0]
	}

	context, err := cfg.getContext(currentContext)
	if err != nil {
		return err
	}

	str, err := context.show()
	if err != nil {
		return err
	}
	fmt.Println(str)
	return nil
}

func ConfigContextTest(cfgOpts *opts.ConfigOptions, cmdOpts *opts.CommandOptions) error {
	file := os.Expand(cfgOpts.ConfigFile, os.Getenv)
	cfg, err := ReadConfig(file)
	if err != nil {
		return err
	}

	var currentContext string

	if len(cmdOpts.Args) == 0 {
		currentContext = cfg.CurrentContext
	} else {
		currentContext = cmdOpts.Args[0]
	}

	isContext := cfg.hasContext(currentContext)
	if !isContext {
		return fmt.Errorf("Context doesn't exist")
	}

	isValidContext, err := cfg.testContext(currentContext, cfgOpts.Debug)
	if (!isValidContext) || (err != nil) {
		return fmt.Errorf("Context is not valid")
	}

	fmt.Println("Context is valid")
	return nil
}

// Change the value of Config.CurrentContext
func ConfigContextUse(cfgOpts *opts.ConfigOptions, cmdOpts *opts.CommandOptions) error {

	context := cmdOpts.Args[0]

	file := os.Expand(cfgOpts.ConfigFile, os.Getenv)
	cfg, err := ReadConfig(file)
	if err != nil {
		return err
	}

	if context == cfg.CurrentContext {
		fmt.Println("Provided context name matches current value. No action needed\n")
		return nil
	}

	if cfg.hasContext(context) {
		cfg.CurrentContext = context
		fmt.Printf("Writing changes...")
		err = cfg.write(file)
		if err != nil {
			return err
		}
		fmt.Printf("  Done!")
	}

	return nil
}

func (c Context) show() (string, error) {
	b, err := yaml.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
