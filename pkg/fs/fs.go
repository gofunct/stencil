package fs

import (
	"github.com/gofunct/stencil/pkg/logging"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
)

type FS struct {
	Afero  *Afero
	Assets *Assets
}

// exists checks if a file or directory exists.
func (f *FS) Exists(path string) bool {
	if path == "" {
		return false
	}
	_, err := f.Afero.Stat(path)
	if err == nil {
		return true
	}
	if !os.IsNotExist(err) {
		logging.Er("file or directory already exists", err)
	}
	return false
}

// findCmdDir checks if base of absPath is cmd dir and returns it or
// looks for existing cmd dir in absPath.
func (f *FS) FindCmdDir(absPath string) string {
	if !f.Exists(absPath) || f.Afero.CheckIfThisDirEmpty(absPath) {
		return "cmd"
	}

	if f.Afero.CheckIfCmdDir(absPath) {
		return filepath.Base(absPath)
	}

	files, _ := filepath.Glob(filepath.Join(absPath, "c*"))
	for _, file := range files {
		if f.Afero.CheckIfCmdDir(file) {
			return filepath.Base(file)
		}
	}

	return "cmd"
}

// findPackage returns full path to existing go package in GOPATHs.
func (f *FS) FindPackage(packageName string) string {
	if packageName == "" {
		return ""
	}

	for _, srcPath := range srcPaths {
		packagePath := filepath.Join(srcPath, packageName)
		if f.Exists(packagePath) {
			return packagePath
		}
	}

	return ""
}

// trimSrcPath trims at the beginning of absPath the srcPath.
func (f *FS) TrimSrcPath(absPath, srcPath string) string {
	relPath, err := filepath.Rel(srcPath, absPath)
	if err != nil {
		logging.Er("failed to trim absPath from srcPath.", err)
	}
	return relPath
}

// isCmdDir checks if base of name is one of cmdDir.
func (f *FS) IsCmdDir(name string) bool {
	name = filepath.Base(name)
	for _, cmdDir := range []string{"cmd", "cmds", "command", "commands"} {
		if name == cmdDir {
			return true
		}
	}
	return false
}

func (f *FS) FilepathHasPrefix(path string, prefix string) bool {
	if len(path) <= len(prefix) {
		return false
	}
	if runtime.GOOS == "windows" {
		// Paths in windows are case-insensitive.
		return strings.EqualFold(path[0:len(prefix)], prefix)
	}
	return path[0:len(prefix)] == prefix

}

func (f *FS) GetCurrentUser() string {
	usr, _ := user.Current()
	return usr.Username
}

func (f *FS) SortedRootAndTemplateFSFiles() []string {
	rootFiles := make([]string, 0, len(f.Assets.fs.Files))
	tmplPaths := make([]string, 0, len(f.Assets.fs.Files))
	for path, entry := range f.Assets.fs.Files {
		if entry.IsDir() {
			continue
		}
		if strings.Count(entry.Path[1:], "/") == 0 {
			rootFiles = append(rootFiles, path)
		} else {
			tmplPaths = append(tmplPaths, path)
		}
	}
	sort.Strings(rootFiles)
	sort.Strings(tmplPaths)
	return append(rootFiles, tmplPaths...)
}
