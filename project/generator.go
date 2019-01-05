package project

import (
	"fmt"
	"github.com/gofunct/common/errors"
	"github.com/spf13/afero"
	"path/filepath"
	"sort"
	"strings"
)

func (p *Project) Generate(dir string, data interface{}) error {
	for _, tmplPath := range p.sortedEntryPaths() {

		//passes tmplPath without .tmpl suffix to
		path, err := TemplateString(strings.TrimSuffix(tmplPath, ".tmpl")).Compile(data)
		if err != nil {
			return errors.Wrapf(err, "failed to parse path: %s", path)
		}
		absPath := filepath.Join(dir, path)
		dirPath := filepath.Dir(absPath)

		// create directory if not exists
		if err := p.CreateDirIfNotExists(p.Os, dirPath); err != nil {
			return errors.WithStack(err)
		}

		entry := p.Assets.Files[tmplPath]

		// generate content
		body, err := TemplateString(string(entry.Data)).Compile(data)
		if err != nil {
			return errors.Wrapf(err, "failed to generate %s", path)
		}
		if ok, err := afero.Exists(p.Os, absPath); err != nil {
			p.Exit("failed to check if path exists" + errors.WithStack(err).Error())
		} else if ok {
			existedBody, err := afero.ReadFile(p.Os, absPath)
			if err != nil {
				p.Exit("failed to read file" + errors.WithStack(err).Error())
			}
			if string(existedBody) == body {
				p.Exit(path[1:] + " is identical.")
			} else {
				p.Exit(path[1:] + " is conflicted.")
			}
		}

		err = afero.WriteFile(p.Os, absPath, []byte(body), 0644)
		if err != nil {
			return errors.Wrapf(err, "failed to write %s", path)
		}

		fmt.Print(p.Notify, path[1:])
	}

	return nil
}

func (p *Project) Destroy(dir string, data interface{}) error {
	for _, tmplPath := range p.sortedEntryPaths() {
		path, err := TemplateString(strings.TrimSuffix(tmplPath, ".tmpl")).Compile(data)
		if err != nil {
			return errors.Wrapf(err, "failed to parse path: %s", path)
		}
		absPath := filepath.Join(dir, path)

		if ok, err := afero.Exists(p.Os, absPath); err != nil {
			p.Notify.Error("failed to get " + path[1:])
			return errors.WithStack(err)
		} else if ok {
			err = p.Os.Remove(absPath)
			if err != nil {
				p.Notify.Error("failed to remove " + path[1:])
				return errors.WithStack(err)
			}
		}

		fmt.Printf("%s, %v", p.Notify, path[1:])

		dirPath := filepath.Dir(path)
		absDirPath := filepath.Dir(absPath)
		if ok, err := afero.DirExists(p.Os, absDirPath); err == nil && ok {
			if r, err := afero.Glob(p.Os, filepath.Join(absDirPath, "*")); err == nil && len(r) == 0 {
				err = p.Os.Remove(absDirPath)
				if err != nil {
					p.Notify.Error("failed to remove " + dirPath[1:])
					return errors.Wrapf(err, "failed to remove %q", dirPath[1:])
				}
				fmt.Printf("%s, %v", p.Notify, dirPath[1:])
			}
		}
	}

	return nil
}

func (p *Project) sortedEntryPaths() []string {
	rootFiles := make([]string, 0, len(p.Assets.Files))
	tmplPaths := make([]string, 0, len(p.Assets.Files))
	for path, entry := range p.Assets.Files {
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
