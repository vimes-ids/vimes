package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var config Config

var rootCmd = &cobra.Command{
	Use:   "vimes",
	Short: "vimes is an intrusion detection system",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "/etc/vimes.yaml", "config file")
}

func initConfig() {
	if cfgFile == "" {
		cfgFile = "/etc/vimes.yaml"
	}

	viper.SetConfigFile(cfgFile)

	if err := viper.ReadInConfig(); err != nil {
		var cfgErr viper.ConfigFileNotFoundError
		if errors.As(err, &cfgErr) {
			fmt.Fprintln(os.Stderr, "Unable to read config file /etc/vimes.yaml")
			os.Exit(1)
		}
	}

	err := viper.Unmarshal(&config)

	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}
