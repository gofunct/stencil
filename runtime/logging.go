package runtime

import (
	"fmt"
	"github.com/gofunct/iio"
	"os"
	"time"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LoggingMode represents a logging configuration specification.
type LoggingMode int

// LoggingMode values
const (
	LoggingNop LoggingMode = iota
	LoggingVerbose
	LoggingDebug
)

var (
	logging = LoggingNop

	// DebugLogConfig is used to generate a *zap.Logger for debug mode.
	DebugLogConfig = func() zap.Config {
		cfg := zap.NewProductionConfig()
		cfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
		cfg.DisableStacktrace = true
		return cfg
	}()
	// VerboseLogConfig is used to generate a *zap.Logger for verbose mode.
	VerboseLogConfig = func() zap.Config {
		cfg := zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		cfg.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
		cfg.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Local().Format("2006-01-02 15:04:05 MST"))
		}
		return cfg
	}()
)

// AddLoggingFlags sets "--debug" and "--verbose" flags to the given *cobra.Command instance.
func (u *UI) AddLoggingFlags(cmd *cobra.Command) {
	var (
		debugEnabled, verboseEnabled bool
	)
	cmd.PersistentFlags().BoolVar(&debugEnabled, "debug", false, fmt.Sprintf("Debug level output"))
	cmd.PersistentFlags().BoolVarP(&verboseEnabled, "verbose", "v", true, fmt.Sprintf("Verbose loggingoutput"))

	cobra.OnInitialize(func() {
		switch {
		case debugEnabled:
			u.Z.With(
				zap.String("exec", cmd.Name()),
				zap.String("version", cmd.Version),
				zap.Bool("runnable", cmd.Runnable()))
			Debug()
		case verboseEnabled:
			u.Z.With(
				zap.String("exec", cmd.Name()),
				zap.String("version", cmd.Version),
				zap.Bool("runnable", cmd.Runnable()),
				zap.Any("meta", cmd.Annotations),
				zap.Bool("is-root", cmd.HasSubCommands()))
			VerboseLog()
		}
	})
}

// Debug sets a debug logger in global.
func Debug() {
	logging = LoggingDebug
	ReplaceLogger(DebugLogConfig)
}

// Verbose sets a verbose logger in global.
func VerboseLog() {
	logging = LoggingVerbose
	ReplaceLogger(VerboseLogConfig)
}

// IsDebug returns true if a debug logger is used.
func IsDebugLog() bool { return logging == LoggingDebug }

// IsVerbose returns true if a verbose logger is used.
func IsVerboseLog() bool { return logging == LoggingVerbose }

// Logging returns a current logging mode.
func Mode() LoggingMode {
	return logging
}

func ReplaceLogger(cfg zap.Config) {
	l, err := cfg.Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to initialize a debug logger: %v\n", err)
	}

	iio.AddCloseFunc(func() { l.Sync() })
	iio.AddCloseFunc(zap.ReplaceGlobals(l))
}
