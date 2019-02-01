package commands

import (
	"github.com/gofunct/stencil"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Intelligently download files or directories from a remote source",
	Run: func(cmd *cobra.Command, args []string) {
		stencil.Load(cfg.Input, cfg.Output)
	},
}
