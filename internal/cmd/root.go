package cmd

import (
	"context"
	"os"

	"charm.land/log/v2"
	"github.com/secfault-org/hacktober/internal/cmd/serve"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFile string
	rootCmd    = &cobra.Command{
		Use:   "hacktober",
		Short: "Hacktober is a ctf like game, to lean about container security",
		Long:  "Hacktober provides a ssh application to play some ctf like games, to learn about pwn, web security",
	}
)

func Execute() error {
	ctx := context.Background()
	return rootCmd.ExecuteContext(ctx)
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "config file (default is $HOME/.config/hacktober.toml)")

	rootCmd.AddCommand(serve.Command)
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home + "/.config")
		viper.SetConfigType("toml")
		viper.SetConfigName("hacktober")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		log.Info("Using config file", "file", viper.ConfigFileUsed())
	}
}
