package ui

import (
	"github.com/tcnksm/go-input"
	"reflect"
	"testing"

	"github.com/gofunct/stencil/pkg/iio"
)

func TestNewUI(t *testing.T) {
	tests := []struct {
		name string
		want *UI
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUI(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUI_ItemSuccess(t *testing.T) {
	type fields struct {
		io      *iio.IO
		inputUI *input.UI
		printer *printer
	}
	type args struct {
		msg string
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
			u := &UI{
				io:      tt.fields.io,
				inputUI: tt.fields.inputUI,
				printer: tt.fields.printer,
			}
			u.ItemSuccess(tt.args.msg)
		})
	}
}

func TestUI_ItemFailure(t *testing.T) {
	type fields struct {
		io      *iio.IO
		inputUI *input.UI
		printer *printer
	}
	type args struct {
		msg  string
		errs []error
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
			u := &UI{
				io:      tt.fields.io,
				inputUI: tt.fields.inputUI,
				printer: tt.fields.printer,
			}
			u.ItemFailure(tt.args.msg, tt.args.errs...)
		})
	}
}

func TestUI_Confirm(t *testing.T) {
	type fields struct {
		io      *iio.IO
		inputUI *input.UI
		printer *printer
	}
	type args struct {
		msg string
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
			u := &UI{
				io:      tt.fields.io,
				inputUI: tt.fields.inputUI,
				printer: tt.fields.printer,
			}
			got, err := u.Confirm(tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("UI.Confirm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UI.Confirm() = %v, want %v", got, tt.want)
			}
		})
	}
}
