package commands

import (
	"github.com/gofunct/common/pkg/logger/zap"
	"github.com/spf13/cobra"
)

// deployCmd represents the deploy command
var RootCmd = &cobra.Command{
	Use:   "stencil",
	Short: "A golang scripting utility tool",
}

func init() {
	RootCmd.AddCommand(loadCmd)
	RootCmd.AddCommand(initCmd)
	RootCmd.AddCommand(watchCmd)
}

// Run adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Run() {
	zap.LogF("Bootstrap error", RootCmd.Execute())
}
