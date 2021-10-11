package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	flagPrivateKey          = "pk"
	flagContract            = "contract"
	flagMethod              = "method"
	flagArgs				= "args"
	flagAfterBlockHeight	= "height"
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

