package commands

import (
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gofunct/common/pkg/zap"
	"github.com/gofunct/stencil"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	Input   string
	Output  string
	Package string
}

var cfg = &Config{}

// deployCmd represents the deploy command
var RootCmd = &cobra.Command{
	Use:   "stencil",
	Short: "A golang scripting utility tool",
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&cfg.Input, "input", "i", "", "generic input flag for all subcommands")
	RootCmd.PersistentFlags().StringVarP(&cfg.Output, "output", "o", ".", "generic output flag for all subcommands")
	RootCmd.PersistentFlags().StringVarP(&cfg.Package, "package", "p", filepath.Base(os.Getenv("PWD")), "generic package flag for all subcommands")
	RootCmd.AddCommand(loadCmd)
	RootCmd.AddCommand(initCmd)
	RootCmd.AddCommand(watchCmd)
	RootCmd.AddCommand(genCmd)
	RootCmd.AddCommand(playCmd)
	RootCmd.AddCommand(debugCmd)
	for _, c := range RootCmd.Commands() {
		if err := viper.BindPFlags(c.Flags()); err != nil {
			fmt.Printf("%s%v", err, err)
		}
		if err := viper.BindPFlags(c.PersistentFlags()); err != nil {
			fmt.Printf("%s%v", err, err)
		}
	}
	viper.AutomaticEnv()
	viper.SetConfigName("stencil")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	if err := viper.ReadInConfig(); err != nil {
		hasCfg := stencil.Prompt("no config found. Do you currently have a config file?")
		if strings.Contains("y", hasCfg) || strings.Contains("Y", hasCfg) || strings.Contains("yes", hasCfg) || strings.Contains("Yes", hasCfg) {
			cfg := stencil.Prompt("Please give the absolute path your config(including its filename and extension)")
			viper.SetConfigFile(cfg)
		} else {
			write := stencil.Prompt("Would you like to write a default config instead?")
			if strings.Contains("y", write) || strings.Contains("Y", write) || strings.Contains("yes", write) || strings.Contains("Yes", write) {
				b, _ := json.Marshal(viper.AllSettings())
				if err := ioutil.WriteFile("stencil.yaml", b, 0755); err != nil {
					fmt.Printf("%s%v", err, err)
				}
			}
		}
	}
	go viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {

		zap.L.Debug("config file changed\n" + in.String())
	})
}

// Run adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Run() {
	zap.LogF("Bootstrap error", RootCmd.Execute())
}
