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

	"github.com/spf13/cobra"
)

var (
	typE string
)
func init() {
	infoCmd.PersistentFlags().StringVarP(&typE, "type", "t", "issues", "type of info needed")
}
// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		if typE =="issues" {
			fmt.Println(`
Colemans-MacBook-Pro:temp coleman$ stencil init -a test
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x80 pc=0x14806a6]

goroutine 1 [running]:
github.com/gofunct/stencil/pkg/runtime.(*Runtime).ErrExit(0x0, 0x15a6a55, 0x17, 0x163cae0, 0xc000136ce0)
        /Users/coleman/go/src/github.com/gofunct/stencil/pkg/runtime/runtime.go:62 +0x26
github.com/gofunct/stencil/pkg/runtime.(*Runtime).TrimSrcPath(0x0, 0xc000026084, 0x35, 0x0, 0x0, 0x0, 0xc00013ef70)
        /Users/coleman/go/src/github.com/gofunct/stencil/pkg/runtime/runtime.go:446 +0xba
github.com/gofunct/stencil/pkg/runtime.NewAppFromPath(0xc000026084, 0x35, 0x0)
        /Users/coleman/go/src/github.com/gofunct/stencil/pkg/runtime/app.go:84 +0x191
github.com/gofunct/stencil/cmd.glob..func3(0x19d3a20, 0xc00013caa0, 0x0, 0x2)
        /Users/coleman/go/src/github.com/gofunct/stencil/cmd/init.go:36 +0x293
github.com/spf13/cobra.(*Command).execute(0x19d3a20, 0xc00013ca80, 0x2, 0x2, 0x19d3a20, 0xc00013ca80)
        /Users/coleman/go/pkg/mod/github.com/spf13/cobra@v0.0.3/command.go:766 +0x2cc
github.com/spf13/cobra.(*Command).ExecuteC(0x19d3c80, 0xc00018df88, 0x10074e0, 0xc000092058)
        /Users/coleman/go/pkg/mod/github.com/spf13/cobra@v0.0.3/command.go:852 +0x2fd
github.com/spf13/cobra.(*Command).Execute(0x19d3c80, 0x0, 0x0)
        /Users/coleman/go/pkg/mod/github.com/spf13/cobra@v0.0.3/command.go:800 +0x2b
github.com/gofunct/stencil/cmd.Execute()
        /Users/coleman/go/src/github.com/gofunct/stencil/cmd/root.go:27 +0x2d
main.main()
        /Users/coleman/go/src/github.com/gofunct/stencil/main.go:26 +0x20



`)
		}
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)

}
