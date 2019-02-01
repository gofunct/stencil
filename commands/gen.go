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

package commands

import (
	"fmt"
	"github.com/Masterminds/sprig"
	"github.com/gofunct/common/pkg/zap"
	"github.com/shiyanhui/hero"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

func init() {
	genCmd.AddCommand(htmlCmd, goGoCmd, grpcCmd)
}

// protocCmd represents the protoc command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "easily generate protobuf stubs",
}

var htmlCmd = &cobra.Command{
	Use:   "html",
	Short: "Generate html files",
	Run: func(cmd *cobra.Command, args []string) {
		WalkHtml(cfg.Input, cfg.Output, cfg.Package)
	},
}

// protocCmd represents the protoc command
var goGoCmd = &cobra.Command{
	Use:   "gogo",
	Short: "Compile gogo protobuf stubs",
	Run: func(cmd *cobra.Command, args []string) {
		zap.LogF("generate gogo protobuf stubs", WalkGoGoProto(cfg.Input))
	},
}

// protocCmd represents the protoc command
var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Compile grpc protobuf stubs",
	Run: func(cmd *cobra.Command, args []string) {
		zap.LogF("generate grpc protobuf stubs", WalkGrpc(cfg.Input))
	},
}

// protocCmd represents the protoc command
var tmplCmd = &cobra.Command{
	Use:   "template",
	Short: "Compile templates",
	Run: func(cmd *cobra.Command, args []string) {
		zap.LogF("generate grpc protobuf stubs", WalkTextTmpl(cfg.Input))
	},
}

func WalkHtml(source, dest, pkg string) {
	hero.Generate(source, dest, pkg)
}

func WalkGrpc(path string) error {
	return filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() && info.Name() == "vendor" {
			return filepath.SkipDir
		}
		// find all protobuf files
		if filepath.Ext(path) == ".proto" {
			// args
			args := []string{
				"-I=.",
				fmt.Sprintf("-I=%s", filepath.Join(os.Getenv("PWD"), "third_party")),
				fmt.Sprintf("--go_out=plugins=grpc:%s", "."),
				path,
			}

			cmd := exec.Command("protoc", args...)
			cmd.Env = os.Environ()
			o, _ := cmd.Output()
			fmt.Println(o)
		}
		return nil
	})
}

func WalkGoGoProto(path string) error {
	return filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		// skip vendor directory
		if info.IsDir() && info.Name() == "vendor" {
			return filepath.SkipDir
		}
		// find all protobuf files
		if filepath.Ext(path) == ".proto" {
			// args
			args := []string{
				"-I=.",
				fmt.Sprintf("-I=%s", filepath.Join(os.Getenv("GOPATH"), "src")),
				fmt.Sprintf("-I=%s", filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "gogo", "protobuf", "protobuf")),
				fmt.Sprintf("--proto_path=%s", filepath.Join(os.Getenv("GOPATH"), "src", "github.com")),
				"--gogofaster_out=Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types:.",
				path,
			}
			cmd := exec.Command("protoc", args...)
			cmd.Env = os.Environ()
			o, _ := cmd.Output()
			fmt.Println(o)
		}
		return nil
	})
}

func WalkTextTmpl(path string) error {
	return filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// skip vendor directory
		if info.IsDir() && info.Name() == "vendor" {
			return filepath.SkipDir
		}

		if strings.Contains(path, ".tmpl"); !strings.Contains(path, ".html") {
			fmt.Printf("%s\n", path)
			tpl := template.Must(
				template.New(filepath.Base(path)).Funcs(sprig.TxtFuncMap()).ParseFiles(path),
			)

			outFile := strings.TrimSuffix(path, ".tmpl")

			f, err := os.Create(outFile)
			defer f.Close()

			if err != nil {
				return err
			}

			return tpl.Execute(f, nil)
		}

		return nil
	})
}
