package project

import (
	"github.com/gofunct/common/errors"
	"github.com/gofunct/common/ui"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"strings"
)

// CreateProjectFromPath returns Project with specified absolute path to
// package.
func CreateProjectFromPath(AbsolutePath string, v *viper.Viper) *Project {
	p := new(Project)
	if p.Notify == nil {
		p.Notify = ui.NewUI()
	}
	if p.Os == nil {
		p.Os = afero.NewOsFs()
	}

	if AbsolutePath == "" {
		exit("can't create project: AbsolutePath can't be blank")
	}
	if !filepath.IsAbs(AbsolutePath) {
		exit("can't create project: AbsolutePath is not absolute")
	}

	// If AbsolutePath is symlink, use its destination.
	fi, err := os.Lstat(AbsolutePath)
	if err != nil {
		exit("can't read path info: " + err.Error())
	}
	if fi.Mode()&os.ModeSymlink != 0 {
		path, err := os.Readlink(AbsolutePath)
		if err != nil {
			exit("can't read the destination of symlink: " + err.Error())
		}
		AbsolutePath = path
	}

	ask := p.Notify.Ask
	p.AbsolutePath = strings.TrimSuffix(AbsolutePath, FindCmdDir(AbsolutePath, v))
	p.ProjName = filepath.ToSlash(TrimSrcPath(p.AbsolutePath, p.SrcPath(v)))
	p.Image = ask(imageAsk)
	p.Version = ask(versionAsk)
	p.CloudRegion = ask(cloudRegionAsk)
	p.RunvarConfig = ask(runvarConfigAsk)
	p.DbPassword = ask(dbPasswordAsk)
	p.DbName = ask(dbNameAsk)
	p.DbUser = ask(dbUserAsk)
	p.DbHost = ask(dbHostAsk)
	p.Bucket = ask(bucketAsk)
	p.Lis = ask(lisAsk)
	p.Dockerhub = ask(dockerhubAsk)
	p.Github = ask(githubAsk)
	p.Services = strings.Split(ask(bucketAsk), ",")
	p.Author = ask(authorAsk)
	p.Tools = strings.Split(ask(toolsAsk), ",")
	p.Scope = ask(scopeAsk)
	p.Import = ask(imageAsk)
	p.LogLevel = ask(logLevelAsk)
	p.Logo = ask(logoAsk)
	p.Notify.Success("all required config values are set!")

	return p
}

// CreateProject returns Project with specified project ProjName.
func CreateProject(v *viper.Viper) *Project {
	p := new(Project)
	if p.Notify == nil {
		p.Notify = ui.NewUI()
	}
	if p.Os == nil {
		p.Os = afero.NewOsFs()
	}

	ask := p.Notify.Ask
	projectName := ask(projNameAsk)
	p.ProjName = projectName

	srcPaths := FindSourcePath(v)

	if projectName == "" {
		exit("can't create project with blank ProjName")
	}

	// 1. Find already created protect.
	p.AbsolutePath = FindPackage(projectName, v)

	// 2. If there are no created project with this path, and user is in GOPATH,
	// then use GOPATH/src/projectName.
	if p.AbsolutePath == "" {
		wd, err := os.Getwd()
		if err != nil {
			exit(err)
		}
		for _, srcPath := range srcPaths {
			goPath := filepath.Dir(srcPath)
			if CheckFilepathHasPrefix(wd, goPath) {
				p.AbsolutePath = filepath.Join(srcPath, projectName)
				break
			}
		}
	}

	// 3. If user is not in GOPATH, then use (first GOPATH)/src/projectName.
	if p.AbsolutePath == "" {
		p.AbsolutePath = filepath.Join(srcPaths[0], projectName)
	}

	p.Image = ask(imageAsk)
	p.Version = ask(versionAsk)
	p.CloudRegion = ask(cloudRegionAsk)
	p.RunvarConfig = ask(runvarConfigAsk)
	p.DbPassword = ask(dbPasswordAsk)
	p.DbName = ask(dbNameAsk)
	p.DbUser = ask(dbUserAsk)
	p.DbHost = ask(dbHostAsk)
	p.Bucket = ask(bucketAsk)
	p.Lis = ask(lisAsk)
	p.Dockerhub = ask(dockerhubAsk)
	p.Github = ask(githubAsk)
	p.Services = strings.Split(ask(bucketAsk), ",")
	p.Author = ask(authorAsk)
	p.Tools = strings.Split(ask(toolsAsk), ",")
	p.Scope = ask(scopeAsk)
	p.Import = ask(imageAsk)
	p.LogLevel = ask(logLevelAsk)
	p.Logo = ask(logoAsk)
	p.Notify.Success("all required config values are set!")

	return p
}

// CreateDirIfNotExists creates a directory if it does not exist.
func (p *Project) CreateDirIfNotExists(fs afero.Fs, path string) (err error) {
	err = fs.MkdirAll(path, 0755)
	zap.L().Debug("CreateDirIfNotExists", zap.String("path", path), zap.Error(err))
	return errors.Wrapf(err, "failed to create %q directory", path)
}
