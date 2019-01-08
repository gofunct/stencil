package fs

import (
	"github.com/gofunct/stencil/pkg/logging"
	"github.com/gofunct/stencil/pkg/ui"
	"os"
	"path/filepath"
	"strings"
)

var srcPaths []string

// Project contains name, license and paths to projects.
type Project struct {
	absPath string
	cmdPath string
	srcPath string
	name    string
	FS      *FS
	UI      *ui.UI
}

// NewProject returns Project with specified project name.
func NewProject(projectName string, fs *FS) *Project {
	if projectName == "" {
		logging.Exit("can't create project with blank name")
	}

	p := &Project{
		name: projectName,
		FS: &FS{
			Afero:  NewAfero(),
			Assets: NewAssets(),
		},
		UI: ui.NewUI(),
	}

	// 1. Find already created protect.
	p.absPath = p.FS.FindPackage(projectName)

	// 2. If there are no created project with this path, and user is in GOPATH,
	// then use GOPATH/src/projectName.
	if p.absPath == "" {
		wd, err := os.Getwd()
		if err != nil {
			logging.Er("faied to get working directory", err)
		}
		for _, srcPath := range srcPaths {
			goPath := filepath.Dir(srcPath)
			if p.FS.FilepathHasPrefix(wd, goPath) {
				p.absPath = filepath.Join(srcPath, projectName)
				break
			}
		}
	}

	// 3. If user is not in GOPATH, then use (first GOPATH)/src/projectName.
	if p.absPath == "" {
		p.absPath = filepath.Join(srcPaths[0], projectName)
	}

	return p
}

// NewProjectFromPath returns Project with specified absolute path to
// package.
func NewProjectFromPath(absPath string) *Project {
	if absPath == "" {
		logging.Exit("can't create project: absPath can't be blank")
	}
	if !filepath.IsAbs(absPath) {
		logging.Exit("can't create project: absPath is not absolute")
	}

	// If absPath is symlink, use its destination.
	fi, err := os.Lstat(absPath)
	if err != nil {
		logging.Er("can't read path info", err)
	}
	if fi.Mode()&os.ModeSymlink != 0 {
		path, err := os.Readlink(absPath)
		if err != nil {
			logging.Er("can't read the destination of symlink", err)
		}
		absPath = path
	}

	p := &Project{
		FS: &FS{
			Afero:  NewAfero(),
			Assets: NewAssets(),
		},
		UI: ui.NewUI(),
	}
	p.absPath = strings.TrimSuffix(absPath, p.FS.FindCmdDir(absPath))
	p.name = filepath.ToSlash(p.FS.TrimSrcPath(p.absPath, p.SrcPath()))
	return p
}

// Name returns the name of project, e.g. "github.com/spf13/cobra"
func (p Project) Name() string {
	return p.name
}

// CmdPath returns absolute path to directory, where all commands are located.
func (p *Project) CmdPath() string {
	if p.absPath == "" {
		return ""
	}
	if p.cmdPath == "" {
		p.cmdPath = filepath.Join(p.absPath, p.FS.FindCmdDir(p.absPath))
	}
	return p.cmdPath
}

// AbsPath returns absolute path of project.
func (p Project) AbsPath() string {
	return p.absPath
}

// SrcPath returns absolute path to $GOPATH/src where project is located.
func (p *Project) SrcPath() string {
	if p.srcPath != "" {
		return p.srcPath
	}
	if p.absPath == "" {
		p.srcPath = srcPaths[0]
		return p.srcPath
	}

	for _, srcPath := range srcPaths {
		if p.FS.FilepathHasPrefix(p.absPath, srcPath) {
			p.srcPath = srcPath
			break
		}
	}

	return p.srcPath
}
