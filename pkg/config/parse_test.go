package config

import "testing"

func TestTemplateString_Compile(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		s       TemplateString
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Compile(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("TemplateString.Compile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TemplateString.Compile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTemplateString_ExecuteTemplate(t *testing.T) {
	type args struct {
		tmplStr string
		data    interface{}
	}
	tests := []struct {
		name    string
		t       *TemplateString
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.t.ExecuteTemplate(tt.args.tmplStr, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("TemplateString.ExecuteTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TemplateString.ExecuteTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTemplateString_CommentifyString(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		t    *TemplateString
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.CommentifyString(tt.args.in); got != tt.want {
				t.Errorf("TemplateString.CommentifyString() = %v, want %v", got, tt.want)
			}
		})
	}
}
