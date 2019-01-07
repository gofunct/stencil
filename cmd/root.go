package cmd

import (
	"github.com/gofunct/stencil/pkg/config"
	"github.com/gofunct/stencil/pkg/logging"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Command struct {
	Cobra  *cobra.Command
	Config *config.Config
}

var c = new(config.Config)
var abs string
var debug bool

func init() {
	c.V = viper.New()
	rootCmd.Cobra.PersistentFlags().StringVarP(&abs, "path", "p", "", "")
	rootCmd.Cobra.PersistentFlags().BoolVar(&debug, "debug-flags", false, "debug flags")
	if debug {
		rootCmd.Cobra.DebugFlags()
	}
	logging.AddLoggingFlags(rootCmd.Cobra)
}

var rootCmd = &Command{
	Config: c,
	Cobra: &cobra.Command{
		Use:     "stencil",
		Aliases: []string{"sten"},
		Short:   "auto generate applications",
	},
}

func Execute() {
	if err := rootCmd.Cobra.Execute(); err != nil {
		logging.Er("failed to execute root command", err)
	}
}
