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
	"fmt"
	"github.com/gofunct/stencil"
	"github.com/gofunct/stencil/pkg/fs"
	"github.com/gofunct/stencil/pkg/watcher"
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

	"github.com/gofunct/stencil/pkg/print"
)

var isWindows = runtime.GOOS == "windows"
var isRebuild bool
var isWatch bool
var isVerbose bool
var hasTasks bool

func checkError(err error, format string, args ...interface{}) {
	if err != nil {
		print.Error("ERR", format, args...)
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
		cmd.Env = stencil.EffectiveEnv(stencil.ParseStringEnv(env))
	}

	return cmd, exe
}

func stencilenv(stencilFile string) string {
	stencilenvFile := filepath.Join(filepath.Dir(stencilFile), "stencilenv")
	if _, err := os.Stat(stencilenvFile); err == nil {
		b, err := ioutil.ReadFile(stencilenvFile)
		if err != nil {
			print.Error("stencil", "Cannot read %s file", stencilenvFile)
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
					print.Debug("stencil", "stencil process killed\n")
				}
			}
		}()
		return cmd, exe
	}

	bufferSize := 2048
	watchr, err := watcher.NewWatcher(bufferSize)
	if err != nil {
		print.Panic("project", "%v\n", err)
	}
	stencilDir := filepath.Dir(stencilFile)
	watchr.WatchRecursive(stencilDir)
	watchr.ErrorHandler = func(err error) {
		print.Error("stencil", "Watcher error %v\n", err)
	}

	cmd, exe := run(false)
	// this function will block forever, Ctrl+C to quit app
	// var lastHappenedTime int64
	watchr.Start()
	print.Info("stencil", "watching %s\n", stencilDir)

	<-time.After(stencil.GetWatchDelay() + (300 * time.Millisecond))

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
			print.Debug("watchmain", "%+v\n", event)
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

	exeFile := "stencilbin-" + stencil.Version
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
		build = fs.Outdated([]string{dir + "/**/*.go"}, []string{exe})
		reasonFormat = "Stencil jobs changed. Rebuilding %s...\n"
	}

	if build {
		print.Debug("stencil", reasonFormat, exe)
		env := stencilenv(src)
		if env != "" {
			stencil.Env = env
		}
		_, err := stencil.Run("go build -a -o "+exeFile, stencil.M{"$in": dir})
		if err != nil {
			panic(fmt.Sprintf("Error building %s: %s\n", src, err.Error()))
		}
		// for some reason go build does not delete the exe named after the dir
		// which ends up with Stencildir/Stencildir
		if filepath.Base(dir) == "bin" {
			orphanedFile := filepath.Join(dir, filepath.Base(dir))
			if _, err := os.Stat(orphanedFile); err == nil {
				os.Remove(orphanedFile)
			}
		}
	}

	if isRebuild {
		print.Info("stencil", "ok\n")
	}

	return exe
}
