package project

import (
	"github.com/gofunct/common/errors"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"sync"
)

//Finding returns a string or stringslice

// findPackage returns full path to existing go package in GOPATHs.
func (p *Project) FindPackage(packageName string, v *viper.Viper) string {
	srcPaths := p.FindSourcePath(v)
	if packageName == "" {
		return ""
	}

	for _, srcPath := range srcPaths {
		packagePath := filepath.Join(srcPath, packageName)
		if p.CheckPathExists(packagePath) {
			return packagePath
		}
	}

	return ""
}

// findCmdDir checks if base of absPath is cmd dir and returns it or
// looks for existing cmd dir in absPath.
func (p *Project) FindCmdDir(absPath string, v *viper.Viper) string {
	if !p.CheckPathExists(absPath) || p.CheckPathEmpty(absPath) {
		return "cmd"
	}

	if p.CheckIfCmdDir(absPath) {
		return filepath.Base(absPath)
	}

	files, _ := filepath.Glob(filepath.Join(absPath, "c*"))
	for _, file := range files {
		if p.CheckIfCmdDir(file) {
			return filepath.Base(file)
		}
	}

	return "cmd"
}

func (p *Project) FindExecutables(fs afero.Fs, v *viper.Viper) []string {
	var wg sync.WaitGroup
	ch := make(chan string)

	for _, path := range filepath.SplitList(os.Getenv("PATH")) {
		wg.Add(1)
		go func(path string) {
			defer wg.Done()

			files, err := afero.ReadDir(fs, path)
			if err != nil {
				return
			}

			for _, f := range files {
				if m := f.Mode(); !f.IsDir() && m&0111 != 0 {
					ch <- f.Name()
				}
			}
		}(path)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	execs := make([]string, 0, 100)
	for exec := range ch {
		execs = append(execs, exec)
	}
	sort.Strings(execs)

	return execs
}

func (p *Project) FindSourcePath(v *viper.Viper) []string {
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

// FindMainPackagesAndSources returns go source file names by main package directories.
func (p *Project) FindMainPackagesAndSources(fs afero.Fs, dir string, v *viper.Viper) (map[string][]string, error) {
	out := make(map[string][]string)
	fset := token.NewFileSet()
	err := afero.Walk(fs, dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return errors.WithStack(err)
		}
		if info.IsDir() || filepath.Ext(info.Name()) != ".go" || strings.HasSuffix(info.Name(), "_test.go") {
			return nil
		}
		data, err := afero.ReadFile(fs, path)
		if err != nil {
			zap.L().Warn("failed to read a file", zap.Error(err), zap.String("path", path))
			return nil
		}
		f, err := parser.ParseFile(fset, "", data, parser.PackageClauseOnly)
		if err != nil {
			zap.L().Warn("failed to parse a file", zap.Error(err), zap.String("path", path), zap.String("body", string(data)))
			return nil
		}
		if f.Package.IsValid() && f.Name.Name == "main" {
			dir := filepath.Dir(path)
			out[dir] = append(out[dir], info.Name())
		}
		return nil
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return out, nil
}

// FindPackageName generates the package name of this application from the given path and envs.
func (p *Project) FindPackageName(rootPath string, v *viper.Viper) (string, error) {
	importPath, err := p.FindImportPath(rootPath, v)
	if err != nil {
		return "", errors.WithStack(err)
	}
	entries := strings.Split(importPath, string(filepath.Separator))
	if len(entries) < 2 {
		u, err := GetOSUser()
		if err != nil {
			return "", errors.WithStack(err)
		}
		entries = []string{u.Username, entries[0]}
	}
	entries = entries[len(entries)-2:]
	if strings.Contains(entries[0], PackageSeparator) {
		s := strings.Split(entries[0], PackageSeparator)
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
		entries[0] = strings.Join(s, PackageSeparator)
	}
	pkgName := strings.Join(entries[len(entries)-2:], PackageSeparator)
	pkgName = strings.Replace(pkgName, "-", "_", -1)
	return pkgName, nil
}

// GetImportPath creates the golang package path from the given path.
func (p *Project) FindImportPath(rootPath string, v *viper.Viper) (importPath string, err error) {
	for _, gopath := range filepath.SplitList(BuildContext.GOPATH) {
		prefix := filepath.Join(gopath, "src") + string(filepath.Separator)
		// FIXME: should not use strings.HasPrefix
		if strings.HasPrefix(rootPath, prefix) {
			importPath = filepath.ToSlash(strings.Replace(rootPath, prefix, "", 1))
			break
		}
	}
	if importPath == "" {
		err = errors.New("failed to get the import path")
	}
	return
}

// BinDir returns the directory path contains executable binaries.
func (d *RootDir) FindBinDir() Path {
	return d.Join("bin")
}
