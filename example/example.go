package main

import (
	"fmt"
	"github.com/gofunct/stencil/runtime"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var cfg runtime.Config

func Execute(cmd *cobra.Command) {
	if err := cmd.Execute(); err != nil {
		zap.L().Fatal("failed to exec command", zap.Error(err))
	}
}

func main() {
	u := runtime.NewUI()
	var rootCmd = &cobra.Command{
		Use:   "stencil",
		Short: "testing",
		Run: func(cmd *cobra.Command, args []string) {
			s := u.ExecuteTemplate(tmplTest)
			fmt.Println(s)
		},
	}
	{
		u.Unmarshal(&cfg)
		u.BindCobra(rootCmd, tmplTest)
	}
	{
		Execute(rootCmd)
	}
}
