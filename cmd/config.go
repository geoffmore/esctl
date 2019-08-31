package cmd

import (
	"fmt"
	"github.com/geoffmore/esctl-go/pkg/escfg"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(cfgCmd)
	cfgCmd.AddCommand(generateCfg)
	cfgCmd.AddCommand(getContexts)
	cfgCmd.AddCommand(currentContext)
	cfgCmd.AddCommand(useContext)
	//cfgCmd.AddCommand(testContext)
	cfgCmd.AddCommand(validateCfg)
	cfgCmd.AddCommand(showCfg)
}

var cfgCmd = &cobra.Command{
	Use:   "config",
	Short: "Interact with a config file",
}

var generateCfg = &cobra.Command{
	Use:   "generate",
	Short: "Generate a config file",
	Long:  `No Description`,
	Run: func(cmd *cobra.Command, args []string) {
		err := escfg.GenDefaultConfig(file)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var getContexts = &cobra.Command{
	Use:   "get-contexts",
	Short: "Get a list of defined contexts",
	Long:  `No Description`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := escfg.ReadConfig(file)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(escfg.GetContexts(cfg))
	},
}

var currentContext = &cobra.Command{
	Use:   "current-context",
	Short: "Get the current context",
	Long:  `No Description`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := escfg.ReadConfig(file)
		if err != nil {
			log.Fatal(err)
		}
		currentContext := cfg.CurrentContext
		if currentContext == "" {
			fmt.Printf("No value for field \"current-context\"\n")
		} else {
			fmt.Printf("Current context is \"%v\"\n", currentContext)
		}
	},
}

var useContext = &cobra.Command{
	Use:     "use-context",
	Short:   "Use a named context",
	Long:    `Use a named context`,
	Example: `esctl config use-context <context>`,
	//Args:
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			errMsg := fmt.Sprintf("Invalid length of arguments. Expecting 1. Got %d\n", len(args))
			log.Fatal("%s\n", errMsg)
		}
		cfg, err := escfg.ReadConfig(file)
		if err != nil {
			log.Fatal(err)
		}
		contexts := escfg.GetContexts(cfg)

		var ctxInCtxs bool = false
		var ctxName string = args[0]
		for _, ctx := range contexts {
			if ctx == ctxName {
				ctxInCtxs = true
			}
		}
		if ctxInCtxs {
			err = escfg.UseContext(ctxName, cfg, file)
			if err == nil {
				fmt.Printf("Current context changed to \"%v\"\n", ctxName)
			} else {
				log.Fatal(err)
			}
		} else {
			fmt.Printf("Context \"%v\" not found\n", ctxName)
		}
	},
}

//var testContext = &cobra.Command{
//	// esctl config
//	Use:   "test-context",
//	Short: "Verify connectivity to a cluster using a context",
//	Long:  `No Description`,
//	Run: func(cmd *cobra.Command, args []string) {
//		fmt.Println("esctl config test-context <context>")
//	},
//}

var validateCfg = &cobra.Command{
	Use:   "validate",
	Short: "Check if a config file is valid",
	Long:  `No Description`,
	Run: func(cmd *cobra.Command, args []string) {
		b, err := escfg.ToBytes(file)
		if err != nil {
			log.Fatal(err)
		} else {

			isValid := escfg.IsValidCfg(b)
			if isValid {
				fmt.Println("Config file is valid")
			} else {
				fmt.Println("Config file is not valid")
			}
		}
	},
}

var showCfg = &cobra.Command{
	Use:   "read",
	Short: "Display a config file",
	Long:  `No Description`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("esctl config read")
		cfg, err := escfg.ReadConfig(file)
		if err != nil {
			log.Fatal(err)
		}
		var b []byte
		b, err = escfg.DisplayConfig(cfg)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v\n", string(b))

	},
}
