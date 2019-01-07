package fs

import (
	"io"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/spf13/afero"
)

func TestNewAfero(t *testing.T) {
	tests := []struct {
		name string
		want *Afero
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAfero(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAfero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAfero_CheckFilepathHasPrefix(t *testing.T) {
	type fields struct {
		af *afero.Afero
	}
	type args struct {
		path   string
		prefix string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Afero{
				af: tt.fields.af,
			}
			if got := c.CheckFilepathHasPrefix(tt.args.path, tt.args.prefix); got != tt.want {
				t.Errorf("Afero.CheckFilepathHasPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAfero_CheckIfCmdDir(t *testing.T) {
	type fields struct {
		af *afero.Afero
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Afero{
				af: tt.fields.af,
			}
			if got := r.CheckIfCmdDir(tt.args.name); got != tt.want {
				t.Errorf("Afero.CheckIfCmdDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAfero_CheckIfThisIsDir(t *testing.T) {
	type fields struct {
		af *afero.Afero
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Afero{
				af: tt.fields.af,
			}
			got, err := a.CheckIfThisIsDir(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Afero.CheckIfThisIsDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Afero.CheckIfThisIsDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAfero_CheckIfFileContainThis(t *testing.T) {
	type fields struct {
		af *afero.Afero
	}
	type args struct {
		filename string
		this     []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Afero{
				af: tt.fields.af,
			}
			got, err := a.CheckIfFileContainThis(tt.args.filename, tt.args.this)
			if (err != nil) != tt.wantErr {
				t.Errorf("Afero.CheckIfFileContainThis() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Afero.CheckIfFileContainThis() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAfero_CheckIfThisDirEmpty(t *testing.T) {
	type fields struct {
		af *afero.Afero
	}
	type args struct {
		path string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Afero{
				af: tt.fields.af,
			}
			if got := a.CheckIfThisDirEmpty(tt.args.path); got != tt.want {
				t.Errorf("Afero.CheckIfThisDirEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAfero_FindAllThisPattern(t *testing.T) {
	type fields struct {
		af *afero.Afero
	}
	type args struct {
		pattern string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Afero{
				af: tt.fields.af,
			}
			got, err := a.FindAllThisPattern(tt.args.pattern)
			if (err != nil) != tt.wantErr {
				t.Errorf("Afero.FindAllThisPattern() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Afero.FindAllThisPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAfero_FindAllProtoFiles(t *testing.T) {
	type fields struct {
		af *afero.Afero
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Afero{
				af: tt.fields.af,
			}
			got, err := a.FindAllProtoFiles()
			if (err != nil) != tt.wantErr {
				t.Errorf("Afero.FindAllProtoFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Afero.FindAllProtoFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAfero_FindAllGoFiles(t *testing.T) {
	type fields struct {
		af *afero.Afero
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Afero{
				af: tt.fields.af,
			}
			got, err := a.FindAllGoFiles()
			if (err != nil) != tt.wantErr {
				t.Errorf("Afero.FindAllGoFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Afero.FindAllGoFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAfero_FindAllYamlFiles(t *testing.T) {
	type fields struct {
		af *afero.Afero
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Afero{
				af: tt.fields.af,
			}
			got, err := a.FindAllYamlFiles()
			if (err != nil) != tt.wantErr {
				t.Errorf("Afero.FindAllYamlFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Afero.FindAllYamlFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAfero_FindAllJsonFiles(t *testing.T) {
	type fields struct {
		af *afero.Afero
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Afero{
				af: tt.fields.af,
			}
			got, err := a.FindAllJsonFiles()
			if (err != nil) != tt.wantErr {
				t.Errorf("Afero.FindAllJsonFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Afero.FindAllJsonFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAfero_FindAllMdFiles(t *testing.T) {
	type fields struct {
		af *afero.Afero
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Afero{
				af: tt.fields.af,
			}
			got, err := a.FindAllMdFiles()
			if (err != nil) != tt.wantErr {
				t.Errorf("Afero.FindAllMdFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Afero.FindAllMdFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAfero_FindAllPBFiles(t *testing.T) {
	type fields struct {
		af *afero.Afero
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Afero{
				af: tt.fields.af,
			}
			got, err := a.FindAllPBFiles()
			if (err != nil) != tt.wantErr {
				t.Errorf("Afero.FindAllPBFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Afero.FindAllPBFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAfero_FindAllShelllFiles(t *testing.T) {
	type fields struct {
		af *afero.Afero
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Afero{
				af: tt.fields.af,
			}
			got, err := a.FindAllShelllFiles()
			if (err != nil) != tt.wantErr {
				t.Errorf("Afero.FindAllShelllFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Afero.FindAllShelllFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAfero_MakeDir(t *testing.T) {
	type fields struct {
		af *afero.Afero
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Afero{
				af: tt.fields.af,
			}
			if err := a.MakeDir(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("Afero.MakeDir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAfero_MakeTempFile(t *testing.T) {
	type fields struct {
		af *afero.Afero
	}
	type args struct {
		dir    string
		prefix string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    afero.File
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Afero{
				af: tt.fields.af,
			}
			got, err := a.MakeTempFile(tt.args.dir, tt.args.prefix)
			if (err != nil) != tt.wantErr {
				t.Errorf("Afero.MakeTempFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Afero.MakeTempFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAfero_MakeTempDir(t *testing.T) {
	type fields struct {
		af *afero.Afero
	}
	type args struct {
		dir    string
		prefix string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Afero{
				af: tt.fields.af,
			}
			got, err := a.MakeTempDir(tt.args.dir, tt.args.prefix)
			if (err != nil) != tt.wantErr {
				t.Errorf("Afero.MakeTempDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Afero.MakeTempDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAfero_WriteToFile(t *testing.T) {
	type fields struct {
		af *afero.Afero
	}
	type args struct {
		filename string
		data     []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Afero{
				af: tt.fields.af,
			}
			if err := a.WriteToFile(tt.args.filename, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Afero.WriteToFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAfero_WriteToReader(t *testing.T) {
	type fields struct {
		af *afero.Afero
	}
	type args struct {
		path string
		r    io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Afero{
				af: tt.fields.af,
			}
			if err := a.WriteToReader(tt.args.path, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("Afero.WriteToReader() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAfero_ReadFromDir(t *testing.T) {
	type fields struct {
		af *afero.Afero
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []os.FileInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Afero{
				af: tt.fields.af,
			}
			got, err := a.ReadFromDir(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Afero.ReadFromDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Afero.ReadFromDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAfero_ReadFromFile(t *testing.T) {
	type fields struct {
		af *afero.Afero
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Afero{
				af: tt.fields.af,
			}
			got, err := a.ReadFromFile(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Afero.ReadFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Afero.ReadFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAfero_OpenFile(t *testing.T) {
	type fields struct {
		af *afero.Afero
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    afero.File
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Afero{
				af: tt.fields.af,
			}
			got, err := a.OpenFile(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Afero.OpenFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Afero.OpenFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAfero_WalkPath(t *testing.T) {
	type fields struct {
		af *afero.Afero
	}
	type args struct {
		path   string
		walkFn filepath.WalkFunc
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Afero{
				af: tt.fields.af,
			}
			if err := a.WalkPath(tt.args.path, tt.args.walkFn); (err != nil) != tt.wantErr {
				t.Errorf("Afero.WalkPath() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAfero_Remove(t *testing.T) {
	type fields struct {
		af *afero.Afero
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Afero{
				af: tt.fields.af,
			}
			if err := a.Remove(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("Afero.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAfero_Rename(t *testing.T) {
	type fields struct {
		af *afero.Afero
	}
	type args struct {
		old string
		new string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Afero{
				af: tt.fields.af,
			}
			if err := a.Rename(tt.args.old, tt.args.new); (err != nil) != tt.wantErr {
				t.Errorf("Afero.Rename() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAfero_ChangePermissions(t *testing.T) {
	type fields struct {
		af *afero.Afero
	}
	type args struct {
		path string
		o    os.FileMode
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Afero{
				af: tt.fields.af,
			}
			if err := a.ChangePermissions(tt.args.path, tt.args.o); (err != nil) != tt.wantErr {
				t.Errorf("Afero.ChangePermissions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAfero_Stat(t *testing.T) {
	type fields struct {
		af *afero.Afero
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    os.FileInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Afero{
				af: tt.fields.af,
			}
			got, err := a.Stat(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Afero.Stat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Afero.Stat() = %v, want %v", got, tt.want)
			}
		})
	}
}
