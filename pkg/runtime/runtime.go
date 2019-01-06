package runtime

import (
	"bytes"
	"fmt"
	"github.com/gofunct/stencil/pkg/ui"
	"github.com/jessevdk/go-assets"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go/parser"
	"go/token"
	"io"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"text/template"
)

var srcPaths []string

func init() {
	// Initialize srcPaths.
	envGoPath := os.Getenv("GOPATH")
	goPaths := filepath.SplitList(envGoPath)
	if len(goPaths) == 0 {
		// Adapted from https://github.com/Masterminds/glide/pull/798/files.
		// As of Go 1.8 the GOPATH is no longer required to be set. Instead there
		// is a default value. If there is no GOPATH check for the default value.
		// Note, checking the GOPATH first to avoid invoking the go toolchain if
		// possible.

		goExecutable := os.Getenv("COBRA_GO_EXECUTABLE")
		if len(goExecutable) <= 0 {
			goExecutable = "go"
		}

		out, err := exec.Command(goExecutable, "env", "GOPATH").Output()
		if err != nil {
			er(err)
		}

		toolchainGoPath := strings.TrimSpace(string(out))
		goPaths = filepath.SplitList(toolchainGoPath)
		if len(goPaths) == 0 {
			er("$GOPATH is not set")
		}
	}
	srcPaths = make([]string, 0, len(goPaths))
	for _, goPath := range goPaths {
		srcPaths = append(srcPaths, filepath.Join(goPath, "src"))
	}
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

type Runtime struct {
	FS     *FileSystem
	TmplFS *assets.FileSystem
	Config *viper.Viper
	Root   RootDir
	UI     ui.UI
	FMap   map[string]interface{}
}

// GetImportPath creates the golang package path from the given path.
func (r *Runtime) GetImportPath(rootPath string) (importPath string, err error) {
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

// GetPackageName generates the package name of this application from the given path and envs.
func (r *Runtime) GetPackageName(rootPath string) (string, error) {
	importPath, err := r.GetImportPath(rootPath)
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

// FindMainPackagesAndSources returns go source file names by main package directories.
func (r *Runtime) FindMainPackagesAndSources(dir string) (map[string][]string, error) {
	out := make(map[string][]string)
	fset := token.NewFileSet()
	err := afero.Walk(r.FS, dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return errors.WithStack(err)
		}
		if info.IsDir() || filepath.Ext(info.Name()) != ".go" || strings.HasSuffix(info.Name(), "_test.go") {
			return nil
		}
		data, err := afero.ReadFile(r.FS, path)
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

func (r *Runtime) Generate(dir string, data interface{}) error {
	for _, tmplPath := range r.sortedEntryPaths() {

		entry := r.TmplFS.Files[tmplPath]
		path, err := TemplateString(strings.TrimSuffix(tmplPath, ".tmpl")).Compile(data)
		if err != nil {
			return errors.Wrapf(err, "failed to parse path: %s", path)
		}
		absPath := filepath.Join(dir, path)
		dirPath := filepath.Dir(absPath)

		// create directory if not exists
		if err := r.FS.MkDir(dirPath); err != nil {
			return errors.WithStack(err)
		}

		// generate content
		body, err := TemplateString(string(entry.Data)).Compile(data)
		if err != nil {
			return errors.Wrapf(err, "failed to generate %s", path)
		}

		// check existed entries
		r.UI.Status = ui.StatusCreate
		if ok, err := afero.Exists(r.FS, absPath); err != nil {
			// TODO: handle an error
			r.UI.Status = ui.StatusSkipped
		} else if ok {
			existedBody, err := afero.ReadFile(r.FS, absPath)
			if err != nil {
				// TODO: handle an error
				r.UI.Status = ui.StatusSkipped
			}
			if string(existedBody) == body {
				r.UI.Status = ui.StatusIdentical
			} else {
				r.UI.Status = ui.StatusSkipped
				r.UI.ItemFailure(path[1:] + " is conflicted.")
				if ok, err := r.UI.Confirm("Overwite it?"); err != nil {
					zap.L().Error("failed to confirm to apply", zap.Error(err))
					return errors.WithStack(err)
				} else if ok {
					r.UI.Status = ui.StatusCreate
				}
			}
		}

		// create
		if r.UI.ShouldCreate() {
			err = afero.WriteFile(r.FS, absPath, []byte(body), 0644)
			if err != nil {
				return errors.Wrapf(err, "failed to write %s", path)
			}
		}

		r.UI.Print(path[1:])
	}

	return nil
}

func (r *Runtime) Destroy(dir string, data interface{}) error {
	for _, tmplPath := range r.sortedEntryPaths() {
		path, err := TemplateString(strings.TrimSuffix(tmplPath, ".tmpl")).Compile(data)
		if err != nil {
			return errors.Wrapf(err, "failed to parse path: %s", path)
		}
		absPath := filepath.Join(dir, path)

		r.UI.Status = ui.StatusSkipped
		if ok, err := afero.Exists(r.FS, absPath); err != nil {
			r.UI.ItemFailure("failed to get " + path[1:])
			return errors.WithStack(err)
		} else if ok {
			err = r.FS.Remove(absPath)
			if err != nil {
				r.UI.ItemFailure("failed to remove " + path[1:])
				return errors.WithStack(err)
			}
			r.UI.Status = ui.StatusDelete
		}

		r.UI.Print(path[1:])

		dirPath := filepath.Dir(path)
		absDirPath := filepath.Dir(absPath)
		if ok, err := afero.DirExists(r.FS, absDirPath); err == nil && ok {
			if x, err := afero.Glob(r.FS, filepath.Join(absDirPath, "*")); err == nil && len(x) == 0 {
				err = r.FS.Remove(absDirPath)
				if err != nil {
					r.UI.ItemFailure("failed to remove " + dirPath[1:])
					return errors.Wrapf(err, "failed to remove %q", dirPath[1:])
				}
				r.UI.Status = ui.StatusDelete
				r.UI.Print(dirPath[1:])
			}
		}
	}

	return nil
}

func (r *Runtime) sortedEntryPaths() []string {
	rootFiles := make([]string, 0, len(r.TmplFS.Files))
	tmplPaths := make([]string, 0, len(r.TmplFS.Files))
	for path, entry := range r.TmplFS.Files {
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

func (r *Runtime) GenerateProject(rootDir, pkgName string) error {
	err := r.mapFromDefaultConfig(rootDir)
	if err != nil {
		return errors.Wrap(err, "failed to create template funcmap from config")
	}
	return r.Generate(rootDir, r.FMap)
}

func (r *Runtime) initConfig(cmd *cobra.Command) error {
	r.Config = viper.New()
	c := r.Config
	c.AddConfigPath(".")
	c.AutomaticEnv()
	c.SetEnvPrefix("stencil")
	c.SetConfigName("stencil")
	c.SetFs(r.FS.RootFS)
	if err := c.SafeWriteConfig(); err != nil {
		return errors.Wrapf(err, "%s", "failed to write config")
	}
	c.BindPFlags(cmd.PersistentFlags())
	c.BindPFlags(cmd.Flags())
	return nil
}

// isEmpty checks if a given path is empty.
// Hidden files in path are ignored.
func (r *Runtime) IsEmpty(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		er(err)
	}

	if !fi.IsDir() {
		return fi.Size() == 0
	}

	f, err := os.Open(path)
	if err != nil {
		er(err)
	}
	defer f.Close()

	names, err := f.Readdirnames(-1)
	if err != nil && err != io.EOF {
		er(err)
	}

	for _, name := range names {
		if len(name) > 0 && name[0] != '.' {
			return false
		}
	}
	return true
}

// exists checks if a file or directory exists.
func (r *Runtime) Exists(path string) bool {
	if path == "" {
		return false
	}
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if !os.IsNotExist(err) {
		er(err)
	}
	return false
}

func (r *Runtime) ExecuteTemplate(tmplStr string, data interface{}) (string, error) {
	tmpl, err := template.New("").Funcs(template.FuncMap{"comment": r.CommentifyString}).Parse(tmplStr)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, data)
	return buf.String(), err
}

func (r *Runtime) WriteStringToFile(path string, s string) error {
	return r.WriteToFile(path)
}

// writeToFile writes r to file with path only
// if file/directory on given path doesn't exist.
func (r *Runtime) WriteToFile(path string) error {
	if r.Exists(path) {
		return fmt.Errorf("%v already exists", path)
	}

	dir := filepath.Dir(path)
	if dir != "" {
		if err := os.MkdirAll(dir, 0777); err != nil {
			return err
		}
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, r.UI.IO.In)
	return err
}

// commentfyString comments every line of in.
func (r *Runtime) CommentifyString(in string) string {
	var newlines []string
	lines := strings.Split(in, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "//") {
			newlines = append(newlines, line)
		} else {
			if line == "" {
				newlines = append(newlines, "//")
			} else {
				newlines = append(newlines, "// "+line)
			}
		}
	}
	return strings.Join(newlines, "\n")
}

func (r *Runtime) Username() string {
	usr, _ := user.Current()
	return usr.Username
}

// returns user home directory
func (r *Runtime) UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

// Generates App Path
func (r *Runtime) CreateAppPath(VCSHost string, user string, appname string) string {
	gopath := os.Getenv("GOPATH")
	if len(gopath) < 1 {
		gopath = r.UserHomeDir() + "/go"
	}
	apppath := gopath + "/src/" + VCSHost + "/" + user + "/" + appname
	os.MkdirAll(apppath, 0755)
	return apppath
}

// runs go format on all generated go files
func (r *Runtime) RunGoFormat(VCSHost string, user string, app string) error {
	gopath := VCSHost + "/" + user + "/" + app
	_, err := exec.Command("go", "fmt", gopath).Output()
	return errors.Wrapf(err, "%s", "failed to execute go format")
}

// trimSrcPath trims at the beginning of AbsPath the srcPath.
func (r *Runtime) TrimSrcPath(AbsPath, srcPath string) string {
	relPath, err := filepath.Rel(srcPath, AbsPath)
	if err != nil {
		er(err)
	}
	return relPath
}

// findPackage returns full path to existing go package in GOPATHs.
func (r *Runtime) FindPackage(packageName string) string {
	if packageName == "" {
		return ""
	}

	for _, srcPath := range srcPaths {
		packagePath := filepath.Join(srcPath, packageName)
		if r.Exists(packagePath) {
			return packagePath
		}
	}

	return ""
}

// findCmdDir checks if base of absPath is cmd dir and returns it or
// looks for existing cmd dir in absPath.
func (r *Runtime) FindCmdDir(absPath string) string {
	if !r.Exists(absPath) || r.IsEmpty(absPath) {
		return "cmd"
	}

	if r.IsCmdDir(absPath) {
		return filepath.Base(absPath)
	}

	files, _ := filepath.Glob(filepath.Join(absPath, "c*"))
	for _, file := range files {
		if r.IsCmdDir(file) {
			return filepath.Base(file)
		}
	}

	return "cmd"
}

// isCmdDir checks if base of name is one of cmdDir.
func (r *Runtime) IsCmdDir(name string) bool {
	name = filepath.Base(name)
	for _, cmdDir := range []string{"cmd", "cmds", "command", "commands"} {
		if name == cmdDir {
			return true
		}
	}
	return false
}

func (r *Runtime) FilepathHasPrefix(path string, prefix string) bool {
	if len(path) <= len(prefix) {
		return false
	}
	if runtime.GOOS == "windows" {
		// Paths in windows are case-insensitive.
		return strings.EqualFold(path[0:len(prefix)], prefix)
	}
	return path[0:len(prefix)] == prefix

}
