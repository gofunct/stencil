package commands

import (
	"github.com/gofunct/stencil/pkg/fs"
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
	Use: "load",
	Run: func(cmd *cobra.Command, args []string) {
		fs.Load(loadSrc, loadDst)
	},
}
