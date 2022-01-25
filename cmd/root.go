package cmd

import (
	"fmt"
	"log"
	"nnsay/ngcli/lib"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ngcli",
	Short: "A CLI tool for cloud service",
	Long: `NGCLI providers the ability to operate and manage Neural Galaxy cloud resource with the commond line style. The resources cloud be:
	- login
	- upload
	- download

But CLI could not do anything for example: view the brain 2D/3D data. The tool is good at non-visualization feature.

CLI support the config file, the config options looks like:
	endpoint: app-api.cloud.cn
	username: user1@cloud.com
	password: user1password
	applicationType: 1

Tip: 
	- ngcli's output format is json, so the jq is good tool for parse the output, eg: ngcli project list | jq
`,

	Run: func(cmd *cobra.Command, args []string) {
		endpoint := viper.GetString("endpoint")
		if endpoint == "" {
			log.Panic("endpoint is requried")
		}
		username := viper.GetString("username")
		if username == "" {
			log.Panic("username is requried")
		}
		password := viper.GetString("password")
		if password == "" {
			log.Panic("password is requried")
		}

		loginCmd.Run(cmd, args)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// config file
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ngcli.yaml)")

	// endpoint
	rootCmd.PersistentFlags().StringP("endpoint", "e", "", "required, the cloud server endpoint, eg: app-api.cloud.cn")
	viper.BindPFlag("endpoint", rootCmd.PersistentFlags().Lookup("endpoint"))
	viper.SetDefault("endpoint", "app-api.cloud.cn")
	rootCmd.MarkFlagRequired("endpoint")
	// username
	rootCmd.PersistentFlags().StringP("username", "u", "", "required, eg: user@cloud.cn")
	rootCmd.MarkFlagRequired("username")
	viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("username"))
	rootCmd.PersistentFlags().StringP("password", "p", "", "required, eg: password for user account")
	rootCmd.MarkFlagRequired("password")
	viper.BindPFlag("password", rootCmd.PersistentFlags().Lookup("password"))

	// applicationType
	rootCmd.PersistentFlags().IntP("applicationType", "a", 0, "required, the production code, eg: data service(1), presurge(2), default 1")
	viper.BindPFlag("applicationType", rootCmd.PersistentFlags().Lookup("applicationType"))
	viper.SetDefault("applicationType", 1)

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".ngcli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".ngcli")
		cfgFile = home + "/.ngcli.yaml"
	}
	lib.MakeSureConfigFile(cfgFile)

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
