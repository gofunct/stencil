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
	"github.com/gofunct/stencil/runtime"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// debugCmd represents the debug command
var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "A brief description of your command",
	PreRun: func(cmd *cobra.Command, args []string) {
		ui.Z.Debug("\n" + runtime.Green("DEBUG COBRA:"))
		cmd.DebugFlags()
	},
	Run: func(cmd *cobra.Command, args []string) {
		ui.Z.Debug("\n" + runtime.Green("DEBUG VIPER:"))
		ui.V.Debug()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		ui.Z.Debug("\n" + runtime.Green("VIPER SETTINGS:"))
		ui.Z.Debug("ui.V.ALLSettings()", zap.Any("settings", ui.V.AllSettings()))
	},
}

func init() {
	root.AddCommand(debugCmd)
}
