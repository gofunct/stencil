package runtime

import (
	"github.com/gofunct/stencil/pkg/ui"
	"github.com/jessevdk/go-assets"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

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
		st := ui.StatusCreate
		if ok, err := afero.Exists(r.FS, absPath); err != nil {
			// TODO: handle an error
			st = ui.StatusSkipped
		} else if ok {
			existedBody, err := afero.ReadFile(r.FS, absPath)
			if err != nil {
				// TODO: handle an error
				st = ui.StatusSkipped
			}
			if string(existedBody) == body {
				st = ui.StatusIdentical
			} else {
				st = ui.StatusSkipped
				r.UI.ItemFailure(path[1:] + " is conflicted.")
				if ok, err := r.UI.Confirm("Overwite it?"); err != nil {
					zap.L().Error("failed to confirm to apply", zap.Error(err))
					return errors.WithStack(err)
				} else if ok {
					st = ui.StatusCreate
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
