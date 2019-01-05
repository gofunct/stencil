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
	"github.com/gofunct/stencil/project"
	"os"

	"github.com/gofunct/common/ui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	currentDir, _ = os.Getwd()
	cfgFile       string
	config        = viper.New()
	notify        = ui.NewUI()
	ask           bool
	write         bool
	proj          = new(project.Project)
)

func init() {
	proj.BindToRoot(rootCmd, config)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "stencil",
	Short: config.GetString("scope"),
	Long: `
  _____ _                  _ _ 
 / ____| |                (_) |
| (___ | |_ ___ _ __   ___ _| |
 \___ \| __/ _ \ '_ \ / __| | |
 ____) | ||  __/ | | | (__| | |
|_____/ \__\___|_| |_|\___|_|_|
                               
By: Coleman Word
Email: coleman.word@gofunct.com
`,
	Version: config.GetString("version"),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
