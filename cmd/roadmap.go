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

// roadmapCmd represents the roadmap command
var roadmapCmd = &cobra.Command{
	Use:   "roadmap",
	Short: "a list of goals for stencil",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
## Project Roadmap

- [ ] submit beta-ready realease v0.1.1
- [ ] pass variables to funcmap through config file, flags, or env(viper, cobra)
- [ ] embed template assets in virtual os with go-asset-builder and afero
- [ ] create custom func map that supports file system methods
- [ ] create plugins architecture to support extensibility
- [ ] create cli utility for common template executions

### Template Variables (viper, cobra, env)
- [x] author
- [x] email
- [x] description
- [x] dependencies
- [x] serviceNames
- [x] gitUser
- [x] dockerUser
- [x] listenPort
- [x] bucket            
- [x] dbHost            
- [x] dbName            
- [x] dbUser            
- [x] dbPassword        
- [x] runVar            
- [x] runVarWaitTime    
- [x] listen            
- [x] cloudSQLRegion    
- [x] runtimeConfigName
- [x] baseImage 

### Func Map(methods)
- [ ] Gopath
- [ ] ImportDir
- [ ] PackageDir
- [ ] CmdDir
- [ ] RootDir
- [ ] VendorDir
- [ ] ProtoDir
- [ ] CamelCase
- [ ] Proto Methods
- [ ] Proto Services
- [ ] Download(go get ...)
- [ ] GoDocs
- [ ] Project File Tree

### File Generation
- [ ] cmd/root.go(cobra): root cobra command
- [ ] cmd/run.go(cobra): start a stackdriver server
- [ ] deploy/: main.tf, vars.tf, output.tf, deployment.yaml(kube manifest)
- [ ] certs/: makefile and .gitignore to generate certs
- [ ] proto/: protobuf file with gogo & gateway annotations
- [ ] bin/: third party binaries
- [ ] main.go: executes root command
- [ ] Dockerfile
- [ ] Makefile
- [ ] README.md
- [ ] CONTRIBUTING.md
- [ ] MIT LICENSE
- [ ] GoDocs2md
- [ ] reviewdog.yaml
- [ ] config.yaml
- [ ] .gitignore
- [ ] .dockerignore
- [ ] .gcloudignore
- [ ] .tools.go: 3rd party binary imports
- [ ] homebrew formulae

`)
	},
}

func init() {
	rootCmd.AddCommand(roadmapCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// roadmapCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// roadmapCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
