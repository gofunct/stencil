package project

import (
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

//Checks return bools

// CheckIfCmdDir checks if base of name is one of cmdDir.
func (p *Project) CheckIfCmdDir(name string) bool {
	name = filepath.Base(name)
	for _, cmdDir := range []string{"cmd", "cmds", "command", "commands"} {
		if name == cmdDir {
			return true
		}
	}
	return false
}

// exists checks if a file or directory exists.
func (p *Project) CheckPathExists(path string) bool {
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

// isEmpty checks if a given path is empty.
// Hidden files in path are ignored.
func (p *Project) CheckPathEmpty(path string) bool {
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

func (p *Project) CheckFilepathHasPrefix(path string, prefix string) bool {
	if len(path) <= len(prefix) {
		return false
	}
	if runtime.GOOS == "windows" {
		// Paths in windows are case-insensitive.
		return strings.EqualFold(path[0:len(prefix)], prefix)
	}
	return path[0:len(prefix)] == prefix

}

// CheckFileExists checks if a file or directory exists.
func (p *Project) CheckFileExists(path string) bool {
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
