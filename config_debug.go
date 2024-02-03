//go:build debug
// +build debug

package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	flagBool   bool
	flagString string
)

func init() {
	virzzCmd := &cobra.Command{
		Use: "virzz",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("%v %v\n", flagBool, cmd.PersistentFlags().Lookup("flag-bool"))
			fmt.Printf("%v %v\n", flagString, cmd.PersistentFlags().Lookup("flag-string"))
			return nil
		},
	}
	virzzCmd.PersistentFlags().BoolVar(&flagBool, "flag-bool", true, "test bool flag")
	virzzCmd.Flags().StringVar(&flagString, "flag-string", "default_string", "test string flag")
	rootCmd.AddCommand(virzzCmd)
}
