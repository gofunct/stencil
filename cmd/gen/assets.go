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

package gen

import (
	"github.com/gofunct/stencil/runtime"
	"github.com/jessevdk/go-assets"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
)

var (
	packageName, varName, strip, output string
)

func init() {
	Root.AddCommand(asset)
	asset.PersistentFlags().StringVarP(&packageName, "package", "p", "", "The package name to generate the assets for")
	asset.PersistentFlags().StringVarP(&varName, "variable", "v", "", "The name of the generated asset tree")
	asset.PersistentFlags().StringVarP(&strip, "strip", "s", "", "Strip the specified prefix from all paths")
	asset.PersistentFlags().StringVarP(&output, "output", "o", "", "File to write output to, or - to write to stdout")
}

// assetsCmd represents the assets command
var asset = &cobra.Command{
	Use:   runtime.Blue("asset"),
	Short: runtime.Blue("An asset builder program to generate embedded assets using go-assets"),
	Run: func(cmd *cobra.Command, args []string) {
		g := &assets.Generator{
			PackageName:  packageName,
			VariableName: varName,
			StripPrefix:  strip,
		}

		for _, f := range args {
			if err := g.Add(f); err != nil {
				ui.Z.Fatal("failed to add file to generator", zap.Error(err))
			}
		}

		if len(output) == 0 || output == "-" {
			g.Write(os.Stdout)
		} else {
			f, err := os.Create(output)

			if err != nil {
				ui.Z.Fatal("failed to create output", zap.Error(err))
			}

			defer f.Close()

			if err := g.Write(f); err != nil {
				ui.Z.Fatal("failed to write file", zap.Error(err))
			}
		}
	},
}
