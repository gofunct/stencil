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
	"github.com/gofunct/stencil/cmd/debug"
	"github.com/gofunct/stencil/cmd/gen"
	"github.com/gofunct/stencil/cmd/sh"
	"go.uber.org/zap"

	"github.com/spf13/cobra"
)

func init() {
	API.AddCommand(
		gen.Root,
		debug.Root,
		sh.Root,
	)
}

var API = &cobra.Command{
	Use:   "stencil",
	Short: "a super amazing golang development utility",
	Long: `
 dP"8   d8                            ,e, 888
C8b Y  d88    ,e e,  888 8e   e88'888  "  888
 Y8b  d88888 d88 88b 888 88b d888  '8 888 888
b Y8D  888   888   , 888 888 Y888   , 888 888
8edP   888    "YeeP" 888 888  "88,e8' 888 888
`,
}

func Execute() {
	if err := API.Execute(); err != nil {
		zap.L().Fatal("failed to execute command", zap.Error(err))
	}
}
