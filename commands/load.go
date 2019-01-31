package commands

import (
	"github.com/gofunct/stencil"
	"github.com/spf13/cobra"
)

var (
	loadSrc string
	loadDst string
)

func init() {
	loadCmd.Flags().StringVar(&loadSrc, "source", "", "the source of the desired file or directory")
	loadCmd.Flags().StringVar(&loadDst, "dest", ".", "the destination of the desired file or directory")
}

// getCmd represents the get command
var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Intelligently download files or directories from a remote source",
	Run: func(cmd *cobra.Command, args []string) {
		stencil.Load(loadSrc, loadDst)
	},
}
