package fs

import (
	"reflect"
	"testing"
)

func TestNewProject(t *testing.T) {
	type args struct {
		projectName string
		fs          *FS
	}
	tests := []struct {
		name string
		args args
		want *Project
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProject(tt.args.projectName, tt.args.fs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewProjectFromPath(t *testing.T) {
	type args struct {
		absPath string
	}
	tests := []struct {
		name string
		args args
		want *Project
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProjectFromPath(tt.args.absPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProjectFromPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProject_Name(t *testing.T) {
	type fields struct {
		absPath string
		cmdPath string
		srcPath string
		name    string
		FS      *FS
		UI      *ui.UI
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
			p := Project{
				absPath: tt.fields.absPath,
				cmdPath: tt.fields.cmdPath,
				srcPath: tt.fields.srcPath,
				name:    tt.fields.name,
				FS:      tt.fields.FS,
				UI:      tt.fields.UI,
			}
			if got := p.Name(); got != tt.want {
				t.Errorf("Project.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProject_CmdPath(t *testing.T) {
	type fields struct {
		absPath string
		cmdPath string
		srcPath string
		name    string
		FS      *FS
		UI      *ui.UI
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
			p := &Project{
				absPath: tt.fields.absPath,
				cmdPath: tt.fields.cmdPath,
				srcPath: tt.fields.srcPath,
				name:    tt.fields.name,
				FS:      tt.fields.FS,
				UI:      tt.fields.UI,
			}
			if got := p.CmdPath(); got != tt.want {
				t.Errorf("Project.CmdPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProject_AbsPath(t *testing.T) {
	type fields struct {
		absPath string
		cmdPath string
		srcPath string
		name    string
		FS      *FS
		UI      *ui.UI
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
			p := Project{
				absPath: tt.fields.absPath,
				cmdPath: tt.fields.cmdPath,
				srcPath: tt.fields.srcPath,
				name:    tt.fields.name,
				FS:      tt.fields.FS,
				UI:      tt.fields.UI,
			}
			if got := p.AbsPath(); got != tt.want {
				t.Errorf("Project.AbsPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProject_SrcPath(t *testing.T) {
	type fields struct {
		absPath string
		cmdPath string
		srcPath string
		name    string
		FS      *FS
		UI      *ui.UI
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
			p := &Project{
				absPath: tt.fields.absPath,
				cmdPath: tt.fields.cmdPath,
				srcPath: tt.fields.srcPath,
				name:    tt.fields.name,
				FS:      tt.fields.FS,
				UI:      tt.fields.UI,
			}
			if got := p.SrcPath(); got != tt.want {
				t.Errorf("Project.SrcPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
