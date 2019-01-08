package runtime

import (
	"testing"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func TestAddLoggingFlags(t *testing.T) {
	type args struct {
		cmd *cobra.Command
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddLoggingFlags(tt.args.cmd)
		})
	}
}

func TestDebug(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Debug()
		})
	}
}

func TestVerboseLog(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			VerboseLog()
		})
	}
}

func TestIsDebugLog(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDebugLog(); got != tt.want {
				t.Errorf("IsDebugLog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsVerboseLog(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsVerboseLog(); got != tt.want {
				t.Errorf("IsVerboseLog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMode(t *testing.T) {
	tests := []struct {
		name string
		want LoggingMode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mode(); got != tt.want {
				t.Errorf("Mode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReplaceLogger(t *testing.T) {
	type args struct {
		cfg zap.Config
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReplaceLogger(tt.args.cfg)
		})
	}
}
