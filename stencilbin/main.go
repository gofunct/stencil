// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/gofunct/gofs"
	"github.com/gofunct/stencil"
	"github.com/mgutz/minimist"
	"os"
	"path/filepath"
)

func main() {
	stencilFiles := []string{"bin/main.go"}
	src := ""
	for _, filename := range stencilFiles {
		src = gofs.FindUp(".", filename)
		if src != "" {
			break
		}
	}

	if src == "" {
		stencil.Usage("")
		os.Exit(0)
	}

	wd, err := os.Getwd()
	if err != nil {
		gofs.Error("stencil", "Could not get working directory: %s\n", err.Error())
	}

	// parent of Stencildir/main.go
	absParentDir, err := filepath.Abs(filepath.Dir(filepath.Dir(src)))
	if err != nil {
		gofs.Error("stencil", "Could not get absolute parent of %s: %s\n", src, err.Error())
	}
	if wd != absParentDir {
		relDir, _ := filepath.Rel(wd, src)
		os.Chdir(absParentDir)
		gofs.Info("stencil", "Using %s\n", relDir)
	}

	os.Setenv("STENCILFILE", src)
	argm := minimist.Parse()
	isRebuild = argm.AsBool("rebuild")
	isWatch = argm.AsBool("w", "watch")
	isVerbose = argm.AsBool("v", "verbose")
	hasTasks = len(argm.NonFlags()) > 0
	run(src)
}
