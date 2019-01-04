// Copyright © 2019 Coleman Word <coleman.word@gofunct.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/gofunct/common/ui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	config  = viper.New()
	options = opts{}
	notify  = ui.NewUI()
	ask     bool
	write   bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:        "stencil",
	Aliases:    options.Aliases,
	SuggestFor: nil,
	Short:      options.Scope,
	Example:    options.Example,
	Version:    options.Version,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type opts struct {
	//CamelCase(upper) -> Kebab-case(lower)
	Logo         string        `mapstructure:"logo"`
	Aliases      []string      `mapstructure:"aliases"`
	Import       string        `mapstructure:"import"`
	Example      string        `mapstructure:"example"`
	Author       string        `mapstructure:"author"`
	Email        string        `mapstructure:"email"`
	Scope        string        `mapstructure:"scope"`
	Tools        []string      `mapstructure:"tools"`
	Version      string        `mapstructure:"version"`
	LogLevel     string        `mapstructure:"log-level"`
	Services     []string      `mapstructure:"services"`
	Github       string        `mapstructure:"github"`
	Dockerhub    string        `mapstructure:"dockerhub"`
	Lis          string        `mapstructure:"lis"`
	Bucket       string        `mapstructure:"bucket"`
	DbHost       string        `mapstructure:"db-host"`
	DbPassword   string        `mapstructure:"db-pw"`
	DbUser       string        `mapstructure:"db-user"`
	DbName       string        `mapstructure:"db-name"`
	Runvar       string        `mapstructure:"runvar"`
	RunVarWait   time.Duration `mapstructure:"runvar-wait"`
	CloudRegion  string        `mapstructure:"cloud-region"`
	RunvarConfig string        `mapstructure:"runvar-config"`
	Image        string        `mapstructure:"image"`
}

func init() {
	pFlags := rootCmd.PersistentFlags()
	config.SetConfigFile("config.yaml")
	config.AddConfigPath(".")
	config.SetEnvPrefix("stencil")
	config.AutomaticEnv()

	{
		pFlags.StringVar(&cfgFile, "config", "", configAsk)
		pFlags.BoolVarP(&ask, "ask","a",  false, "fill out a survey to set configs")
		pFlags.BoolVarP(&write, "write", "w",false, "fill out a survey to set configs")
		pFlags.StringVar(&options.Example, "example", "", exampleAsk)
		pFlags.StringSliceVar(&options.Aliases, "aliases", []string{""}, aliasesAsk)
		pFlags.StringVar(&options.Logo, "logo", "https://github.com/gofunct/common/blob/master/logo/dark_logo_transparent_background.png?raw=true", logoAsk)
		pFlags.StringVar(&options.Author, "author", "Coleman Word", authorAsk)
		pFlags.StringVar(&options.Email, "email", "coleman.word@gofunct.com", emailAsk)
		pFlags.StringVar(&options.Scope, "scope", "not available", scopeAsk)
		pFlags.StringVar(&options.Version, "version", "v0.1.1", versionAsk)
		pFlags.StringVar(&options.LogLevel, "log-level", "verbose", logLevelAsk)
		pFlags.StringSliceVar(&options.Tools, "tools", []string{"gex", "reviewdog", "errcheck", "gogoslick", "grpc-swagger", "grpc-gateway", "protocgen-go", "golint", "megacheck", "wire", "unparam", "gox", "gotests"}, toolsAsk)
		pFlags.StringSliceVar(&options.Services, "services", []string{""}, servicesAsk)
		pFlags.StringVar(&options.Github, "github", "gofunct", githubAsk)
		pFlags.StringVar(&options.Dockerhub, "dockerhub", "", dockerhubAsk)
		pFlags.StringVar(&options.Lis, "lis", ":8080", lisAsk)
		pFlags.StringVar(&options.Bucket, "bucket", "gofunct-storage", bucketAsk)
		pFlags.StringVar(&options.DbHost, "db-host", "", dbHostAsk)
		pFlags.StringVar(&options.DbPassword, "db-pw", "admin", dbPasswordAsk)
		pFlags.StringVar(&options.DbPassword, "db-user", "gofunct", dbUserAsk)
		pFlags.StringVar(&options.DbPassword, "db-name", "gofunct-db", dbNameAsk)
		pFlags.StringVar(&options.Runvar, "runvar", "", runvarAsk)
		pFlags.DurationVar(&options.RunVarWait, "runvar-wait", 5*time.Second, runvarWaitAsk)
		pFlags.StringVar(&options.CloudRegion, "cloud-region", "", cloudRegionAsk)
		pFlags.StringVar(&options.RunvarConfig, "runvar-config", "", runvarConfigAsk)
		pFlags.StringVar(&options.Image, "image", "alpine:3.8", imageAsk)
	}

	config.BindPFlags(pFlags)
	config.Unmarshal(&options)
	}
