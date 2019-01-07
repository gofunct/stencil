package fs

import (
	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"go.uber.org/zap"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type Afero struct {
	af *afero.Afero
}

func NewAfero() *Afero {
	return &Afero{
		af: &afero.Afero{
			Fs: afero.NewOsFs(),
		},
	}
}

///////////////////////////CHECK///////////////////////////

func (c *Afero) CheckFilepathHasPrefix(path string, prefix string) bool {
	if len(path) <= len(prefix) {
		return false
	}
	if runtime.GOOS == "windows" {
		// Paths in windows are case-insensitive.
		return strings.EqualFold(path[0:len(prefix)], prefix)
	}
	return path[0:len(prefix)] == prefix

}

// isCmdDir checks if base of name is one of cmdDir.
func (r *Afero) CheckIfCmdDir(name string) bool {
	name = filepath.Base(name)
	for _, cmdDir := range []string{"cmd", "cmds", "command", "commands"} {
		if name == cmdDir {
			return true
		}
	}
	return false
}

func (a *Afero) CheckIfThisIsDir(path string) (bool, error) {
	return a.af.IsDir(path)
}

func (a *Afero) CheckIfFileContainThis(filename string, this []byte) (bool, error) {
	return a.af.FileContainsBytes(filename, this)
}

func (a *Afero) CheckIfThisDirEmpty(path string) bool {
	b, err := a.af.IsEmpty(path)
	zap.L().Fatal("Checking if directory is empty", zap.String("path", path), zap.Error(err))
	return b
}

func (a *Afero) FindAllThisPattern(pattern string) ([]string, error) {
	f, err := afero.Glob(a.af, pattern)
	zap.L().Debug("Finding all files with pattern", zap.String("pattern", pattern), zap.Error(err))
	return f, err
}

func (a *Afero) FindAllProtoFiles() ([]string, error) {
	f, err := afero.Glob(a.af, "*.proto")
	zap.L().Debug("Finding all proto files", zap.Error(err))
	return f, err
}

func (a *Afero) FindAllGoFiles() ([]string, error) {
	f, err := afero.Glob(a.af, "*.go")
	zap.L().Debug("Finding all go files", zap.Error(err))
	return f, err
}

func (a *Afero) FindAllYamlFiles() ([]string, error) {
	f, err := afero.Glob(a.af, "*.yaml")
	zap.L().Debug("Finding all yaml files", zap.Error(err))
	return f, err
}

func (a *Afero) FindAllJsonFiles() ([]string, error) {
	f, err := afero.Glob(a.af, "*.json")
	zap.L().Debug("Finding all json files", zap.Error(err))
	return f, err
}

func (a *Afero) FindAllMdFiles() ([]string, error) {
	f, err := afero.Glob(a.af, "*.md")
	zap.L().Debug("Finding all markdown files", zap.Error(err))
	return f, err
}

func (a *Afero) FindAllPBFiles() ([]string, error) {
	f, err := afero.Glob(a.af, "*.pb.go")
	zap.L().Debug("Finding all generated protobuf files", zap.Error(err))
	return f, err
}

func (a *Afero) FindAllShelllFiles() ([]string, error) {
	f, err := afero.Glob(a.af, "*.sh")
	zap.L().Debug("Finding all shell files", zap.Error(err))
	return f, err
}

///////////////////////////MAKE///////////////////////////

func (a *Afero) MakeDir(path string) error {
	err := a.af.MkdirAll(path, 0755)
	zap.L().Debug("Making All Directories", zap.String("path", path), zap.Error(err))
	return errors.Wrapf(err, "failed to create %q directory", path)
}

func (a *Afero) MakeTempFile(dir, prefix string) (afero.File, error) {
	f, err := a.af.TempFile(dir, prefix)
	zap.L().Debug("Making Temporary File", zap.String("dir", dir), zap.String("prefix", prefix), zap.Error(err))
	return f, err
}

func (a *Afero) MakeTempDir(dir, prefix string) (string, error) {
	s, err := a.af.TempDir(dir, prefix)
	zap.L().Debug("Making Temporary Directory", zap.String("dir", dir), zap.String("prefix", prefix), zap.Error(err))
	return s, err
}

///////////////////////////WRITE///////////////////////////

func (a *Afero) WriteToFile(filename string, data []byte) error {
	err := a.af.WriteFile(filename, data, 0755)
	zap.L().Debug("Writing to File", zap.String("filename", filename), zap.ByteString("data", data), zap.Error(err))
	return err
}

func (a *Afero) WriteToReader(path string, r io.Reader) error {
	err := a.af.WriteReader(path, r)
	zap.L().Debug("Writing to File", zap.String("path", path), zap.Any("reader", r), zap.Error(err))
	return err
}

///////////////////////////READ///////////////////////////

func (a *Afero) ReadFromDir(path string) ([]os.FileInfo, error) {
	i, err := a.af.ReadDir(path)
	zap.L().Debug("Reading directory", zap.String("path", path), zap.Error(err))

	return i, err
}

func (a *Afero) ReadFromFile(path string) ([]byte, error) {
	b, err := a.af.ReadFile(path)
	zap.L().Debug("Reading file", zap.String("path", path), zap.Error(err))
	return b, err
}

func (a *Afero) OpenFile(path string) (afero.File, error) {
	f, err := a.af.Open(path)
	zap.L().Debug("Opening file", zap.String("path", path), zap.Error(err))
	return f, err
}

///////////////////////////WALK///////////////////////////

func (a *Afero) WalkPath(path string, walkFn filepath.WalkFunc) error {
	err := a.af.Walk(path, walkFn)
	zap.L().Debug("Walking path with func", zap.String("path", path), zap.Error(err))
	return err
}

///////////////////////////LIST///////////////////////////

///////////////////////////DELETE///////////////////////////

func (a *Afero) Remove(path string) error {
	err := a.af.Remove(path)
	zap.L().Debug("Removing file", zap.String("path", path), zap.Error(err))
	return err
}

///////////////////////////OTHER///////////////////////////

func (a *Afero) Rename(old, new string) error {
	err := a.af.Rename(old, new)
	zap.L().Debug("Renaming", zap.String("old", old), zap.String("new", new), zap.Error(err))
	return err
}

func (a *Afero) ChangePermissions(path string, o os.FileMode) error {
	err := a.af.Chmod(path, o)
	zap.L().Debug("Changing permissions", zap.String("path", path), zap.Any("file-mode", o), zap.Error(err))
	return err
}

func (a *Afero) Stat(name string) (os.FileInfo, error) {
	o, err := a.af.Stat(name)
	zap.L().Debug("Changing permissions", zap.String("name", name), zap.Error(err))
	return o, err
}
