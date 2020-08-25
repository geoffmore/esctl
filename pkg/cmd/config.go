package cmd

import (
	"github.com/geoffmore/esctl/pkg/escfg"
	"github.com/spf13/cobra"
	"log"
)

// Variables needed for initCredOpts
var addresses, user, pass, token, apiKey, cloudID string
var insecure bool

// ES_CONFIG variable should be a thing eventually

func init() {
	rootCmd.AddCommand(configCmd)
	//configCmd.AddCommand(generateCfg)
	//configCmd.AddCommand(getContexts)
	//configCmd.AddCommand(currentContext)
	//configCmd.AddCommand(useContext)
	//configCmd.AddCommand(genContext)
	//configCmd.AddCommand(testContext)
	//configCmd.AddCommand(validateCfg)
	//configCmd.AddCommand(showCfg)

	//// Context functions
	configCmd.AddCommand(configGenContext)
	configGenContext.Flags().StringVarP(&addresses, "addresses", "", "", "comma-delimeted list of addresses used to attempt a connection")
	configGenContext.Flags().StringVarP(&cloudID, "cloud-id", "", "", "cloud-id to connect to")
	configGenContext.Flags().StringVarP(&user, "user", "", "", "username to connect with")
	configGenContext.Flags().StringVarP(&pass, "password", "", "", "password to connect with")
	configGenContext.Flags().StringVarP(&token, "token", "", "", "token to connect with")
	configGenContext.Flags().StringVarP(&apiKey, "api-key", "", "", "api-key to connect with")
	configGenContext.Flags().BoolVarP(&insecure, "insecure", "", false, "whether or not to skip tls verification")
	//configCmd.AddCommand(configGetContexts)
	//configCmd.AddCommand(configGetCurrentContext)
	//configCmd.AddCommand(configShowContext)
	//configCmd.AddCommand(configTestContext)
	//configCmd.AddCommand(configUseContext)
	//// Config functions
	configCmd.AddCommand(configGenDefaultConfig)
	//configCmd.AddCommand(configGenConfig)
	//configCmd.AddCommand(configShowConfig)
	//configCmd.AddCommand(configTestConfig)

}

//// Boilerplate //
//// Command init
//initCmdOpts(cmd, cmdOpts, args)
//// Client init
//initCfgOpts(cmd, cfgOpts)
//client, err := genClientWOpts(cfgOpts)
//if err != nil {
//	log.Fatal(err)
//}
//// Everything else

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Interact with a config file",
}

//var configGenerateConfig = &cobra.Command{
//	Use:   "generate",
//	Short: "Generate a config file",
//	Run: func(cmd *cobra.Command, args []string) {
//		// Boilerplate //
//		// Client init
//		initCfgOpts(cmd, cfgOpts)
//		// Everything else
//		err := escfg.ConfigGenerateConfig(cmdOpts)
//		if err != nil {
//			log.Fatal(err)
//		}
//	},
//}
var configGenDefaultConfig = &cobra.Command{
	Use:   "gen-default-config",
	Short: "generate the default config file",
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate //
		// Client init
		initCfgOpts(cmd, cfgOpts)
		// Everything else
		err := escfg.ConfigGenDefaultConfig(cfgOpts)
		if err != nil {
			log.Fatal(err)
		}
	},
}

//
//var configGetContexts = &cobra.Command{
//	Use:   "get-contexts",
//	Short: "Get a list of defined contexts",
//	Long:  `No Description`,
//	Run: func(cmd *cobra.Command, args []string) {
//		// Boilerplate //
//		// Client init
//		initCfgOpts(cmd, cfgOpts)
//		// Everything else
//		err := escfg.GetContexts(cfgOpts)
//		if err != nil {
//			log.Fatal(err)
//		}
//		//fmt.Println(escfg.GetContexts(cfg))
//	},
//}
//
//var currentContext = &cobra.Command{
//	Use:   "current-context",
//	Short: "Get the current context",
//	Long:  `No Description`,
//	Run: func(cmd *cobra.Command, args []string) {
//		// Boilerplate //
//		// Client init
//		initCfgOpts(cmd, cfgOpts)
//		// Everything else
//		err := escfg.CurrentContext(cfgOpts)
//
//		//cfg, err := escfg.ReadConfig(file)
//		if err != nil {
//			log.Fatal(err)
//		}
//		//currentContext := cfg.CurrentContext
//		//if currentContext == "" {
//		//	fmt.Printf("No value for field \"current-context\"\n")
//		//} else {
//		//	fmt.Printf("Current context is \"%v\"\n", currentContext)
//		//}
//	},
//}
//
//var useContext = &cobra.Command{
//	Use:   "use-context",
//	Short: "Use a named context",
//	Args:  cobra.MinimumNArgs(1),
//	//Args:
//	Run: func(cmd *cobra.Command, args []string) {
//		if len(args) != 1 {
//			errMsg := fmt.Sprintf("Invalid length of arguments. Expecting 1. Got %d\n", len(args))
//			log.Fatal("%s\n", errMsg)
//		}
//		// Boilerplate //
//		// Command init
//		initCmdOpts(cmd, cmdOpts, args)
//		// Client init
//		initCfgOpts(cmd, cfgOpts)
//		// Everything else
//		err := escfg.UseContext(cfgOpts)
//		if err != nil {
//			log.Fatal(err)
//		}
//		//cfg, err := escfg.ReadConfig(file)
//		//if err != nil {
//		//	log.Fatal(err)
//		//}
//		//contexts := escfg.GetContexts(cfg)
//
//		//var ctxInCtxs bool = false
//		//var ctxName string = args[0]
//		//for _, ctx := range contexts {
//		//	if ctx == ctxName {
//		//		ctxInCtxs = true
//		//	}
//		//}
//		//if ctxInCtxs {
//		//	err = escfg.UseContext(ctxName, cfg, file)
//		//	if err == nil {
//		//		fmt.Printf("Current context changed to \"%v\"\n", ctxName)
//		//	} else {
//		//		log.Fatal(err)
//		//	}
//		//} else {
//		//	fmt.Printf("Context \"%v\" not found\n", ctxName)
//		//}
//	},
//}
//
////var testContext = &cobra.Command{
////	// esctl config
////	Use:   "test-context",
////	Short: "Verify connectivity to a cluster using a context",
////	Long:  `No Description`,
////	Run: func(cmd *cobra.Command, args []string) {
////		fmt.Println("esctl config test-context <context>")
////	},
////}
//
//var validateCfg = &cobra.Command{
//	Use:   "validate",
//	Short: "Check if a config file is valid",
//	Long:  `No Description`,
//	Run: func(cmd *cobra.Command, args []string) {
//		// Boilerplate //
//		// Command init
//		initCmdOpts(cmd, cmdOpts, args)
//		// Client init
//		initCfgOpts(cmd, cfgOpts)
//		// Everything else
//		err := escfg.ValidateConfig(cfgOpts)
//		if err != nil {
//			log.Fatal(err)
//		}
//		//err := escfg.UseContext(cfgOpts)
//		//b, err := escfg.ToBytes(file)
//		//if err != nil {
//		//	log.Fatal(err)
//		//} else {
//		//	isValid := escfg.IsValidCfg(b)
//		//	if isValid {
//		//		fmt.Println("Config file is valid")
//		//	} else {
//		//		fmt.Println("Config file is not valid")
//		//	}
//		//}
//	},
//}
//
//var showCfg = &cobra.Command{
//	Use:   "read",
//	Short: "Display a config file",
//	Long:  `No Description`,
//	Run: func(cmd *cobra.Command, args []string) {
//		fmt.Println("esctl config read")
//		cfg, err := escfg.ReadConfig(file)
//		if err != nil {
//			log.Fatal(err)
//		}
//		var b []byte
//		b, err = escfg.DisplayConfig(cfg)
//		if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Printf("%+v\n", string(b))
//
//	},
//}
//
////var showContext
//
//var genContext = &cobra.Command{
//	Use:   "gen-context",
//	Short: "Generate a context after verifying cluster connectivity",
//	Run: func(cmd *cobra.Command, args []string) {
//		// Boilerplate //
//		// Command init
//		initCmdOpts(cmd, cmdOpts, args)
//		// Client init
//		initCfgOpts(cmd, cfgOpts)
//		// Everything else
//		err := escfg.GenerateContext(cfgOpts)
//		if err != nil {
//			log.Fatal(err)
//		}
//	},
//}

//// Context functions
//configCmd.AddCommand(configGenContext)
//configCmd.AddCommand(configGetCurrentContext)
//configCmd.AddCommand(configShowContext)
//configCmd.AddCommand(configTestContext)
//configCmd.AddCommand(configUseContext)
//// Config functions
//configCmd.AddCommand(configGenConfig)
//configCmd.AddCommand(configShowConfig)
//configCmd.AddCommand(configTestConfig)

// Boilerplate
//var foo = &cobra.Command{
//	Use:   "foo",
//	Short: "foo",
//	Run: func(cmd *cobra.Command, args []string) {
//		// Boilerplate //
//		// Command init
//		initCmdOpts(cmd, cmdOpts, args)
//		// Client init
//		initCfgOpts(cmd, cfgOpts)
//		// Everything else
//		err := escfg.foo(cmdOpts)
//		if err != nil {
//			log.Fatal(err)
//		}
//	},
//}

var configGenContext = &cobra.Command{
	Use:   "gen-context",
	Short: "Generate a context based on provided connection details",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate //
		// Command flag init (only for args)
		initCmdOpts(cmd, cmdOpts, args)
		// Config flag init
		initCfgOpts(cmd, cfgOpts)
		// Credential init
		initCredOpts(cmd, credOpts)

		// Everything else
		err := escfg.ConfigGenContext(cfgOpts, cmdOpts, credOpts)
		if err != nil {
			log.Fatal(err)
		}
	},
}

//configCmd.AddCommand(configGetCurrentContext)
//var configGetCurrentContext = &cobra.Command{
//	Use:   "current-context",
//	Short: "Get the current context",
//	Run: func(cmd *cobra.Command, args []string) {
//		// Boilerplate //
//		// Command init
//		initCmdOpts(cmd, cmdOpts, args)
//		// Client init
//		initCfgOpts(cmd, cfgOpts)
//		// Everything else
//		err := escfg.foo(cmdOpts)
//		if err != nil {
//			log.Fatal(err)
//		}
//	},
//}
