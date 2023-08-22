package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/mcuadros/go-defaults"
	"github.com/oxygenz6/cfwips/cfg"
	"github.com/oxygenz6/cfwips/lib/storage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "cfwips",
	Short: "Cloudflare Whitelisted IP Scanner",
	Long:  `A utility program to help you find whitelisted cloudflare IPs with your active connection.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig, func() {
		defaults.SetDefaults(cfg.Instance)

		if err := viper.UnmarshalExact(cfg.Instance); err != nil {
			panic(errors.Join(errors.New("failed to unmarshal conf"), err))
		}

		storage.Init(cfg.Instance.DbPath)
	})

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cfwips.yaml)")
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

		// Search config in home directory with name ".cfwips" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cfwips")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
