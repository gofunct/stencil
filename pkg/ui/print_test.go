package ui

import (
	"bytes"
	"io"
	"reflect"
	"testing"

	"github.com/dixonwille/wlog"
	"github.com/gofunct/stencil/pkg/iio"
)

func TestNewPrinter(t *testing.T) {
	tests := []struct {
		name string
		want *printer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPrinter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPrinter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_printer_ChangeReaderWriter(t *testing.T) {
	type fields struct {
		UI      *wlog.ConcurrentUI
		IO      *iio.IO
		Closers []func()
	}
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name            string
		fields          fields
		args            args
		wantWriter      string
		wantErrorWriter string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &printer{
				UI:      tt.fields.UI,
				IO:      tt.fields.IO,
				Closers: tt.fields.Closers,
			}
			writer := &bytes.Buffer{}
			errorWriter := &bytes.Buffer{}
			r.ChangeReaderWriter(tt.args.reader, writer, errorWriter)
			if gotWriter := writer.String(); gotWriter != tt.wantWriter {
				t.Errorf("printer.ChangeReaderWriter() = %v, want %v", gotWriter, tt.wantWriter)
			}
			if gotErrorWriter := errorWriter.String(); gotErrorWriter != tt.wantErrorWriter {
				t.Errorf("printer.ChangeReaderWriter() = %v, want %v", gotErrorWriter, tt.wantErrorWriter)
			}
		})
	}
}

func Test_printer_Close(t *testing.T) {
	type fields struct {
		UI      *wlog.ConcurrentUI
		IO      *iio.IO
		Closers []func()
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &printer{
				UI:      tt.fields.UI,
				IO:      tt.fields.IO,
				Closers: tt.fields.Closers,
			}
			u.Close()
		})
	}
}

func Test_printer_AddCloseFunc(t *testing.T) {
	type fields struct {
		UI      *wlog.ConcurrentUI
		IO      *iio.IO
		Closers []func()
	}
	type args struct {
		f func()
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
			u := &printer{
				UI:      tt.fields.UI,
				IO:      tt.fields.IO,
				Closers: tt.fields.Closers,
			}
			u.AddCloseFunc(tt.args.f)
		})
	}
}
