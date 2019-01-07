package config

import (
	"html/template"
	"reflect"
	"testing"

	"github.com/gofunct/stencil/pkg/fs"
	"github.com/spf13/viper"
)

func TestConfig_MergeWithProject(t *testing.T) {
	type fields struct {
		FMap        *template.FuncMap
		TemplString *TemplateString
		V           *viper.Viper
		P           *fs.Project
		Meta        map[string]interface{}
	}
	type args struct {
		p *fs.Project
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				FMap:        tt.fields.FMap,
				TemplString: tt.fields.TemplString,
				V:           tt.fields.V,
				P:           tt.fields.P,
				Meta:        tt.fields.Meta,
			}
			c.MergeWithProject(tt.args.p)
		})
	}
}

func TestConfig_AddToSettings(t *testing.T) {
	type fields struct {
		FMap        *template.FuncMap
		TemplString *TemplateString
		V           *viper.Viper
		P           *fs.Project
		Meta        map[string]interface{}
	}
	type args struct {
		m map[string]interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				FMap:        tt.fields.FMap,
				TemplString: tt.fields.TemplString,
				V:           tt.fields.V,
				P:           tt.fields.P,
				Meta:        tt.fields.Meta,
			}
			if got := c.AddToSettings(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Config.AddToSettings() = %v, want %v", got, tt.want)
			}
		})
	}
}
