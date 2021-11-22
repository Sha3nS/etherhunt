package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	flagConfigPath  = "config"
)

// attachFlags
func attachFlags(cmd *cobra.Command, flagNames []string) {
	flags := &pflag.FlagSet{}
	cmdFlags := cmd.Flags()
	for _, flagName := range flagNames {
		if flag := flags.Lookup(flagName); flag != nil {
			cmdFlags.AddFlag(flag)
		}
	}
}

