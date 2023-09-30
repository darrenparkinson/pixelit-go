package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/darrenparkinson/pixelit-go"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Embed cobra.Command inside cliApp so we can access it and also add our own version field
// https://github.com/spf13/cobra/issues/282
type cliApp struct {
	cobra.Command
	Version string
}

var pxc *pixelit.Client
var cfgFile string

// Embedded fields can't be used as struct literals, so we have to have this weird syntax.
// https://github.com/golang/go/issues/9859
var rootCmd = &cliApp{Command: cobra.Command{
	Use:              "pixelit-cli",
	Short:            "PixelIt CLI is used to interact with the PixelIt API.",
	Version:          "0.0.1",
	PersistentPreRun: checkForConfig,
}}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
// func Execute() {
// 	if err := rootCmd.Execute(); err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// }

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default $HOME/.pixelit.json)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "provide verbose output")
	rootCmd.PersistentFlags().String("host", "", "host name or ip of your PixelIt")
	cobra.OnInitialize(initConfig)
	cobra.AddTemplateFunc("version", func() string { return rootCmd.Version })
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	// From the command line
	if err := viper.BindPFlags(rootCmd.Flags()); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".pixelit" (without extension).
		viper.SetConfigName(".pixelit")
		viper.AddConfigPath(home)
		viper.AddConfigPath("./")
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		if viper.GetBool("verbose") {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
	}

	// From the environment
	viper.SetEnvPrefix("pixelit") // will be uppercased automatically
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv() // read in environment variables that match

}

func checkForConfig(cmd *cobra.Command, args []string) {
	host := viper.GetString("host")
	requiredConfig := true
	if host == "" {
		requiredConfig = false
	}
	if !requiredConfig {
		fmt.Println("Missing required config, host must be specified.")
		os.Exit(1)
	}
	var err error
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	pxc, err = pixelit.NewClient(host, client)
	if err != nil {
		fmt.Printf("error getting gitlab client: %s\n", err)
	}
}
