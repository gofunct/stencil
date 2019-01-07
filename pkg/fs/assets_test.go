package fs

import (
	"github.com/jessevdk/go-assets"
	"net/http"
	"reflect"
	"testing"
)

func TestNewAssets(t *testing.T) {
	tests := []struct {
		name string
		want *Assets
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAssets(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAssets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAssets_ListFSFiles(t *testing.T) {
	type fields struct {
		fs *assets.FileSystem
		g  *assets.Generator
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]*assets.File
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Assets{
				fs: tt.fields.fs,
				g:  tt.fields.g,
			}
			if got := a.ListFSFiles(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Assets.ListFSFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAssets_OpenFSFile(t *testing.T) {
	type fields struct {
		fs *assets.FileSystem
		g  *assets.Generator
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    http.File
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Assets{
				fs: tt.fields.fs,
				g:  tt.fields.g,
			}
			got, err := a.OpenFSFile(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Assets.OpenFSFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Assets.OpenFSFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAssets_NewFSFile(t *testing.T) {
	type fields struct {
		fs *assets.FileSystem
		g  *assets.Generator
	}
	type args struct {
		path string
		data []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *assets.File
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Assets{
				fs: tt.fields.fs,
				g:  tt.fields.g,
			}
			if got := a.NewFSFile(tt.args.path, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Assets.NewFSFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAssets_GetFSPackage(t *testing.T) {
	type fields struct {
		fs *assets.FileSystem
		g  *assets.Generator
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Assets{
				fs: tt.fields.fs,
				g:  tt.fields.g,
			}
			if got := a.GetFSPackage(); got != tt.want {
				t.Errorf("Assets.GetFSPackage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAssets_GetFSPrefix(t *testing.T) {
	type fields struct {
		fs *assets.FileSystem
		g  *assets.Generator
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Assets{
				fs: tt.fields.fs,
				g:  tt.fields.g,
			}
			if got := a.GetFSPrefix(); got != tt.want {
				t.Errorf("Assets.GetFSPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAssets_AddFileToFS(t *testing.T) {
	type fields struct {
		fs *assets.FileSystem
		g  *assets.Generator
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
			a := &Assets{
				fs: tt.fields.fs,
				g:  tt.fields.g,
			}
			if err := a.AddFileToFS(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("Assets.AddFileToFS() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAssets_WriteFSFile(t *testing.T) {
	type fields struct {
		fs *assets.FileSystem
		g  *assets.Generator
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
			a := &Assets{
				fs: tt.fields.fs,
				g:  tt.fields.g,
			}
			if err := a.WriteFSFile(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("Assets.WriteFSFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAssets_GetFSVariable(t *testing.T) {
	type fields struct {
		fs *assets.FileSystem
		g  *assets.Generator
	}
	type args struct {
		f string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Assets{
				fs: tt.fields.fs,
				g:  tt.fields.g,
			}
			if got := a.GetFSVariable(tt.args.f); got != tt.want {
				t.Errorf("Assets.GetFSVariable() = %v, want %v", got, tt.want)
			}
		})
	}
}
