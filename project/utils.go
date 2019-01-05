package project

import (
	"context"
	"github.com/gofunct/common/errors"
	"github.com/gofunct/common/ui"
	"github.com/izumin5210/gex"
	"github.com/spf13/viper"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func exit(msg interface{}) {
	ui.NewUI().Error(msg.(string))
	os.Exit(1)
}

func (p *Project) InstallDeps(rootDir string, cfg gex.Config) error {
	p.GexCfg.WorkingDir = rootDir
	repo, err := p.GexCfg.Create()
	if err == nil {
		spec := p.BuildSpec()
		err = repo.Add(
			context.TODO(),
			"github.com/gofunct/stencil"+spec,
			"github.com/golang/protobuf/protoc-gen-go",
			"github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway",
			"github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger",
		)
	}
	return errors.WithStack(err)
}

func FindSourcePath(v *viper.Viper) []string {
	var srcPaths []string
	// Initialize srcPaths.
	envGoPath := os.Getenv("GOPATH")
	v.BindEnv("GOPATH")
	goPaths := filepath.SplitList(envGoPath)
	if len(goPaths) == 0 {
		// Adapted from https://github.com/Masterminds/glide/pull/798/files.
		// As of Go 1.8 the GOPATH is no longer required to be set. Instead there
		// is a default value. If there is no GOPATH check for the default value.
		// Note, checking the GOPATH first to avoid invoking the go toolchain if
		// possible.

		goExecutable := os.Getenv("COBRA_GO_EXECUTABLE")
		v.BindEnv("COBRA_GO_EXECUTABLE")
		if len(goExecutable) <= 0 {
			goExecutable = "go"
		}

		out, err := exec.Command(goExecutable, "env", "GOPATH").Output()
		if err != nil {
			exit(err)
		}

		toolchainGoPath := strings.TrimSpace(string(out))
		goPaths = filepath.SplitList(toolchainGoPath)
		if len(goPaths) == 0 {
			exit("$GOPATH is not set")
		}
	}
	srcPaths = make([]string, 0, len(goPaths))
	for _, goPath := range goPaths {
		srcPaths = append(srcPaths, filepath.Join(goPath, "src"))
	}
	return srcPaths
}

func CheckFilepathHasPrefix(path string, prefix string) bool {
	if len(path) <= len(prefix) {
		return false
	}
	if runtime.GOOS == "windows" {
		// Paths in windows are case-insensitive.
		return strings.EqualFold(path[0:len(prefix)], prefix)
	}
	return path[0:len(prefix)] == prefix

}

// findPackage returns full path to existing go package in GOPATHs.
func FindPackage(packageName string, v *viper.Viper) string {
	srcPaths := FindSourcePath(v)
	if packageName == "" {
		return ""
	}

	for _, srcPath := range srcPaths {
		packagePath := filepath.Join(srcPath, packageName)
		if CheckPathExists(packagePath) {
			return packagePath
		}
	}

	return ""
}

// exists checks if a file or directory exists.
func CheckPathExists(path string) bool {
	if path == "" {
		return false
	}
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if !os.IsNotExist(err) {
		exit(err)
	}
	return false
}

// findCmdDir checks if base of absPath is cmd dir and returns it or
// looks for existing cmd dir in absPath.
func FindCmdDir(absPath string, v *viper.Viper) string {
	if !CheckPathExists(absPath) || CheckPathEmpty(absPath) {
		return "cmd"
	}

	if CheckIfCmdDir(absPath) {
		return filepath.Base(absPath)
	}

	files, _ := filepath.Glob(filepath.Join(absPath, "c*"))
	for _, file := range files {
		if CheckIfCmdDir(file) {
			return filepath.Base(file)
		}
	}

	return "cmd"
}

// CheckIfCmdDir checks if base of name is one of cmdDir.
func CheckIfCmdDir(name string) bool {
	name = filepath.Base(name)
	for _, cmdDir := range []string{"cmd", "cmds", "command", "commands"} {
		if name == cmdDir {
			return true
		}
	}
	return false
}

// isEmpty checks if a given path is empty.
// Hidden files in path are ignored.
func CheckPathEmpty(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		exit(err)
	}

	if !fi.IsDir() {
		return fi.Size() == 0
	}

	f, err := os.Open(path)
	if err != nil {
		exit(err)
	}
	defer f.Close()

	names, err := f.Readdirnames(-1)
	if err != nil && err != io.EOF {
		exit(err)
	}

	for _, name := range names {
		if len(name) > 0 && name[0] != '.' {
			return false
		}
	}
	return true
}

// TrimSrcPath trims at the beginning of absPath the srcPath.
func TrimSrcPath(absPath, srcPath string) string {
	relPath, err := filepath.Rel(srcPath, absPath)
	if err != nil {
		exit(err)
	}
	return relPath
}
