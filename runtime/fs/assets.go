package fs

import (
	"github.com/jessevdk/go-assets"
	"go.uber.org/zap"
	"net/http"
	"os"
	"time"
)

type Assets struct {
	fs *assets.FileSystem
	g  *assets.Generator
}

func NewAssets() *Assets {
	return &Assets{
		fs: &assets.FileSystem{},
		g:  &assets.Generator{},
	}
}

///////////////////////////////FS///////////////////////////////

func (a *Assets) ListFSFiles() map[string]*assets.File {
	return a.fs.Files
}

func (a *Assets) OpenFSFile(path string) (http.File, error) {
	f, err := a.fs.Open(path)
	zap.L().Debug("Opening assetfs file", zap.String("path", path), zap.Error(err))
	return f, err
}

func (a *Assets) NewFSFile(path string, data []byte) *assets.File {
	return a.fs.NewFile(path, 0755, time.Now(), data)
}

///////////////////////////////Generator///////////////////////////////

func (a *Assets) GetFSPackage() string {
	return a.g.PackageName
}

func (a *Assets) GetFSPrefix() string {
	return a.g.StripPrefix
}

func (a *Assets) AddFileToFS(path string) error {
	err := a.g.Add(path)
	zap.L().Debug("Adding asset-generator file", zap.String("path", path), zap.Error(err))
	return err
}

func (a *Assets) WriteFSFile(path string) error {
	err := a.g.Write(os.Stdout)
	zap.L().Debug("Writing asset-generator file", zap.String("path", path), zap.Error(err))
	return err
}

func (a *Assets) GetFSVariable(f string) string {
	return a.g.VariableName
}
