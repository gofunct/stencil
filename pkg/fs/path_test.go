package fs

import "testing"

func TestPath_String(t *testing.T) {
	tests := []struct {
		name string
		p    Path
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("Path.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPath_Join(t *testing.T) {
	type args struct {
		elem []string
	}
	tests := []struct {
		name string
		p    Path
		args args
		want Path
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Join(tt.args.elem...); got != tt.want {
				t.Errorf("Path.Join() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRootDir_String(t *testing.T) {
	tests := []struct {
		name string
		d    RootDir
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.String(); got != tt.want {
				t.Errorf("RootDir.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRootDir_Join(t *testing.T) {
	type args struct {
		elem []string
	}
	tests := []struct {
		name string
		d    RootDir
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Join(tt.args.elem...); got != tt.want {
				t.Errorf("RootDir.Join() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRootDir_BinDir(t *testing.T) {
	tests := []struct {
		name string
		d    RootDir
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.BinDir(); got != tt.want {
				t.Errorf("RootDir.BinDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRootDir_CmdDir(t *testing.T) {
	tests := []struct {
		name string
		d    RootDir
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.CmdDir(); got != tt.want {
				t.Errorf("RootDir.CmdDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRootDir_VendorDir(t *testing.T) {
	tests := []struct {
		name string
		d    RootDir
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.VendorDir(); got != tt.want {
				t.Errorf("RootDir.VendorDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRootDir_ProtoDir(t *testing.T) {
	tests := []struct {
		name string
		d    RootDir
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.ProtoDir(); got != tt.want {
				t.Errorf("RootDir.ProtoDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRootDir_ApiDir(t *testing.T) {
	tests := []struct {
		name string
		d    RootDir
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.ApiDir(); got != tt.want {
				t.Errorf("RootDir.ApiDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRootDir_DocsDir(t *testing.T) {
	tests := []struct {
		name string
		d    RootDir
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.DocsDir(); got != tt.want {
				t.Errorf("RootDir.DocsDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRootDir_StaticDir(t *testing.T) {
	tests := []struct {
		name string
		d    RootDir
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.StaticDir(); got != tt.want {
				t.Errorf("RootDir.StaticDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
