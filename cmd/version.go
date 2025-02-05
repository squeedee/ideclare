package cmd

import (
	"fmt"
	"runtime/debug"

	"github.com/spf13/cobra"
)

var Version string

func init() {
	if Version == "" {
		i, ok := debug.ReadBuildInfo()
		if !ok {
			return
		}
		Version = i.Main.Version
	}
}

func NewCmdVersion() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if Version == "" {
				return fmt.Errorf("could not determine build information")
			} else {
				_, err := fmt.Fprintln(cmd.OutOrStdout(), Version)
				if err != nil {
					return fmt.Errorf("could not print version: %w", err)
				}
			}
			return nil
		},
	}
}
