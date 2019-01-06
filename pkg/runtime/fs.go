package runtime

import (
	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"go.uber.org/zap"
	"go/build"
	"io"
	"os"
	"os/user"
	"path/filepath"
)

type getOSUserFunc func() (*user.User, error)

const (
	// PackageSeparator is a package separator string on protobuf.
	PackageSeparator = "."
)

// Make visible for testing
var (
	BuildContext build.Context
	GetOSUser    getOSUserFunc
)

func init() {
	BuildContext = build.Default
	GetOSUser = user.Current
}

type FileSystem struct {
	*afero.Afero
	RootFS afero.Fs
}

func NewFileSystem(root string) *FileSystem {
	fs := afero.NewMemMapFs()

	return &FileSystem{
		Afero: &afero.Afero{
			Fs: fs,
		},
		RootFS: afero.NewBasePathFs(fs, root),
	}
}

func (f *FileSystem) MkDir(path string) error {
	err := f.MkdirAll(path, 0755)
	zap.L().Debug("CreateDirIfNotExists", zap.String("path", path), zap.Error(err))
	return errors.Wrapf(err, "failed to create %q directory", path)
}

func (f *FileSystem) IsDirectory(path string) (bool, error) {
	return f.IsDir(path)
}

func (f *FileSystem) DoesFileContainThis(filename string, this []byte) (bool, error) {
	return f.FileContainsBytes(filename, this)
}

func (f *FileSystem) IsDirEmpty(path string) (bool, error) {
	return f.IsEmpty(path)
}

func (f *FileSystem) ReadFromDir(path string) ([]os.FileInfo, error) {
	return f.ReadDir(path)
}

func (f *FileSystem) ReadFromFile(path string) ([]byte, error) {
	return f.ReadFile(path)
}

func (f *FileSystem) WalkPath(root string, walkFn filepath.WalkFunc) error {
	return f.Walk(root, walkFn)
}

func (f *FileSystem) WriteToFile(filename string, data []byte) error {
	return f.WriteFile(filename, data, 0755)
}

func (f *FileSystem) NewTempDir(dir, prefix string) (string, error) {
	return f.TempDir(dir, prefix)
}

func (f *FileSystem) NewTempFile(dir, prefix string) (afero.File, error) {
	return f.TempFile(dir, prefix)
}

func (f *FileSystem) WriteToReader(path string, r io.Reader) (err error) {
	return f.WriteReader(path, r)
}
