package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func setPFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().VisitAll(func(f *pflag.Flag) {
		if !f.Changed && viper.IsSet(f.Name) {
			f.Value.Set(viper.GetString(f.Name))
		}
	})
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		if !f.Changed && viper.IsSet(f.Name) {
			f.Value.Set(viper.GetString(f.Name))
		}
	})
	for _, c := range cmd.Commands() {
		setPFlag(c)
	}
}

func init() {
	viper.SetConfigName("podman_flag")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.config")
	viper.ReadInConfig()
	viper.AutomaticEnv()
	viper.SetEnvPrefix("podman")
	setPFlag(rootCmd)
}
