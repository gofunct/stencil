package ui

import (
	"github.com/spf13/cobra"
	"strings"
	"testing"
)

func TestLogging(t *testing.T) {
	var u = NewPrinter()

	cases := []struct {
		args      []string
		mode      LoggingMode
		isDebug   bool
		isVerbose bool
	}{
		{
			mode: LoggingNop,
		},
		{
			args:      []string{"-v"},
			mode:      LoggingVerbose,
			isVerbose: true,
		},
		{
			args:      []string{"--verbose"},
			mode:      LoggingVerbose,
			isVerbose: true,
		},
		{
			args:    []string{"--debug"},
			mode:    LoggingDebug,
			isDebug: true,
		},
	}

	for _, tc := range cases {
		t.Run(strings.Join(tc.args, " "), func(t *testing.T) {
			defer u.Close()

			var (
				mode               LoggingMode
				isDebug, isVerbose bool
			)

			cmd := &cobra.Command{
				Run: func(*cobra.Command, []string) {
					mode = u.LoggingMode()
					isDebug = u.IsDebugLog()
					isVerbose = u.IsVerboseLog()
				},
			}

			u.AddLoggingFlags(cmd)
			cmd.SetArgs(tc.args)
			err := cmd.Execute()

			if err != nil {
				t.Errorf("Execute() returned an error: %v", err)
			}

			if got, want := mode, tc.mode; got != want {
				t.Errorf("LoggingMode() returned %v, want %v", got, want)
			}

			if got, want := isVerbose, tc.isVerbose; got != want {
				t.Errorf("IsVerbose() returned %t, want %t", got, want)
			}

			if got, want := isDebug, tc.isDebug; got != want {
				t.Errorf("IsDebug() returned %t, want %t", got, want)
			}
		})
	}
}
