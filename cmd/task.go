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

package cmd

import (
	"fmt"
	"github.com/gofunct/stencil/cmd/task"
	"github.com/gofunct/stencil/util"
	"github.com/gofunct/stencil/watcher"
	"github.com/mgutz/minimist"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

var isWindows = runtime.GOOS == "windows"
var isRebuild bool
var isWatch bool
var isVerbose bool
var hasTasks bool

// taskCmd represents the task command
var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		// v2 ONLY uses Stencildir/main.go
		stencilFiles := []string{"Stencildir/main.go", "Stencildir/Stencilfile.go", "tasks/Stencilfile.go"}
		src := ""
		for _, filename := range stencilFiles {
			src = util.FindUp(".", filename)
			if src != "" {
				break
			}
		}

		if src == "" {
			task.Usage("")
			os.Exit(0)
		}

		wd, err := os.Getwd()
		if err != nil {
			util.Error("stencil", "Could not get working directory: %s\n", err.Error())
		}

		// parent of Stencildir/main.go
		absParentDir, err := filepath.Abs(filepath.Dir(filepath.Dir(src)))
		if err != nil {
			util.Error("stencil", "Could not get absolute parent of %s: %s\n", src, err.Error())
		}
		if wd != absParentDir {
			relDir, _ := filepath.Rel(wd, src)
			os.Chdir(absParentDir)
			util.Info("stencil", "Using %s\n", relDir)
		}

		os.Setenv("STENCILFILE", src)
		argm := minimist.Parse()
		isRebuild = argm.AsBool("rebuild")
		isWatch = argm.AsBool("w", "watch")
		isVerbose = argm.AsBool("v", "verbose")
		hasTasks = len(argm.NonFlags()) > 0
		run(src)
	},
}

func checkError(err error, format string, args ...interface{}) {
	if err != nil {
		util.Error("ERR", format, args...)
		os.Exit(1)
	}
}

func hasMain(data []byte) bool {
	hasMainRe := regexp.MustCompile(`\nfunc main\(`)
	matches := hasMainRe.Find(data)
	return len(matches) > 0
}

func isPackageMain(data []byte) bool {
	isMainRe := regexp.MustCompile(`(\n|^)?package main\b`)
	matches := isMainRe.Find(data)
	return len(matches) > 0
}

func run(stencilFile string) {
	if isWatch {
		runAndWatch(stencilFile)
	} else {
		cmd, _ := buildCommand(stencilFile, false)
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func buildCommand(stencilFile string, forceBuild bool) (*exec.Cmd, string) {
	exe := buildMain(stencilFile, forceBuild)
	cmd := exec.Command(exe, os.Args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	//cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	// process stencilenv file
	env := stencilenv(stencilFile)
	if env != "" {
		cmd.Env = task.EffectiveEnv(task.ParseStringEnv(env))
	}

	return cmd, exe
}

func stencilenv(stencilFile string) string {
	stencilenvFile := filepath.Join(filepath.Dir(stencilFile), "stencilenv")
	if _, err := os.Stat(stencilenvFile); err == nil {
		b, err := ioutil.ReadFile(stencilenvFile)
		if err != nil {
			util.Error("stencil", "Cannot read %s file", stencilenvFile)
			os.Exit(1)
		}
		return string(b)
	}
	return ""
}

func runAndWatch(stencilFile string) {
	done := make(chan bool, 1)
	run := func(forceBuild bool) (*exec.Cmd, string) {
		cmd, exe := buildCommand(stencilFile, forceBuild)
		cmd.Start()
		go func() {
			err := cmd.Wait()
			done <- true
			if err != nil {
				if isVerbose {
					util.Debug("stencil", "stencil process killed\n")
				}
			}
		}()
		return cmd, exe
	}

	bufferSize := 2048
	watchr, err := watcher.NewWatcher(bufferSize)
	if err != nil {
		util.Panic("project", "%v\n", err)
	}
	stencilDir := filepath.Dir(stencilFile)
	watchr.WatchRecursive(stencilDir)
	watchr.ErrorHandler = func(err error) {
		util.Error("stencil", "Watcher error %v\n", err)
	}

	cmd, exe := run(false)
	// this function will block forever, Ctrl+C to quit app
	// var lastHappenedTime int64
	watchr.Start()
	util.Info("stencil", "watching %s\n", stencilDir)

	<-time.After(task.GetWatchDelay() + (300 * time.Millisecond))

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			killStencil(cmd, false)
			os.Exit(0)
		}
	}()

	// forloop:
	for {
		select {
		case event := <-watchr.Event:
			// looks like go build starts with the output file as the dir, then
			// renames it to output file
			if event.Path == exe || event.Path == path.Join(path.Dir(exe), path.Base(path.Dir(exe))) {
				continue
			}
			util.Debug("watchmain", "%+v\n", event)
			killStencil(cmd, true)
			<-done
			cmd, _ = run(true)
		}
	}

}

// killStencil kills the spawned stencil process.
func killStencil(cmd *exec.Cmd, killProcessGroup bool) {
	cmd.Process.Kill()
	// process group may not be cross platform but on Darwin and Linux, this
	// is the only way to kill child processes
	if killProcessGroup {
		pgid, err := syscall.Getpgid(cmd.Process.Pid)
		if err != nil {
			panic(err)
		}
		syscall.Kill(-pgid, syscall.SIGKILL)
	}
}

func mustBeMain(src string) {
	data, err := ioutil.ReadFile(src)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if !hasMain(data) {
		msg := `%s is not runnable. Rename package OR make it runnable by adding

	func main() {
		stencil.Stencil(tasks)
	}
	`
		fmt.Printf(msg, src)
		os.Exit(1)
	}

	if !isPackageMain(data) {
		msg := `%s is not runnable. It must be package main`
		fmt.Printf(msg, src)
		os.Exit(1)
	}
}

func buildMain(src string, forceBuild bool) string {
	mustBeMain(src)
	dir := filepath.Dir(src)

	exeFile := "stencilbin-" + task.Version
	if isWindows {
		exeFile += ".exe"
	}

	exe := filepath.Join(dir, exeFile)

	build := false
	reasonFormat := ""
	if isRebuild || forceBuild {
		build = true
		reasonFormat = "Rebuilding %s...\n"
	} else {
		build = util.Outdated([]string{dir + "/**/*.go"}, []string{exe})
		reasonFormat = "Stencil tasks changed. Rebuilding %s...\n"
	}

	if build {
		util.Debug("stencil", reasonFormat, exe)
		env := stencilenv(src)
		if env != "" {
			task.Env = env
		}
		_, err := task.Run("go build -a -o "+exeFile, task.M{"$in": dir})
		if err != nil {
			panic(fmt.Sprintf("Error building %s: %s\n", src, err.Error()))
		}
		// for some reason go build does not delete the exe named after the dir
		// which ends up with Stencildir/Stencildir
		if filepath.Base(dir) == "Stencildir" {
			orphanedFile := filepath.Join(dir, filepath.Base(dir))
			if _, err := os.Stat(orphanedFile); err == nil {
				os.Remove(orphanedFile)
			}
		}
	}

	if isRebuild {
		util.Info("stencil", "ok\n")
	}

	return exe
}
