package root

import (
	"fmt"
	"github.com/gofunct/stencil/runtime/ui"
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
func (s *scriptr) AddLoggingFlags(cmd *cobra.Command) {
	var (
		debugEnabled, verboseEnabled bool
	)
	cmd.PersistentFlags().BoolVar(&debugEnabled, ui.Blue("debug"), false, ui.Blue("Debug level output"))
	cmd.PersistentFlags().BoolVarP(&verboseEnabled, "verbose", "v", true, ui.Blue("Verbose loggingoutput"))

	cobra.OnInitialize(func() {
		switch {
		case debugEnabled:
			s.Z.With(
				zap.String("exec", cmd.Name()),
				zap.String("version", cmd.Version),
				zap.Bool("runnable", cmd.Runnable()))
			s.Debug()
		case verboseEnabled:
			s.Z.With(
				zap.String("exec", cmd.Name()),
				zap.String("version", cmd.Version),
				zap.Bool("runnable", cmd.Runnable()),
				zap.Any("meta", cmd.Annotations),
				zap.Bool("is-root", cmd.HasSubCommands()))
			s.VerboseLog()
		}
	})
}

// Debug sets a debug logger in global.
func (s *scriptr) Debug() {
	logging = LoggingDebug
	s.ReplaceLoggerConfig(DebugLogConfig)
}

// Verbose sets a verbose logger in global.
func (s *scriptr) VerboseLog() {
	logging = LoggingVerbose
	s.ReplaceLoggerConfig(VerboseLogConfig)
}

// IsDebug returns true if a debug logger is used.
func IsDebugLog() bool { return logging == LoggingDebug }

// IsVerbose returns true if a verbose logger is used.
func IsVerboseLog() bool { return logging == LoggingVerbose }

// Logging returns a current logging mode.
func Mode() LoggingMode {
	return logging
}

func (s *scriptr) ReplaceLoggerConfig(cfg zap.Config) {
	l, err := cfg.Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to initialize a debug logger: %v\n", err)
	}

	s.AddCloser(func() { l.Sync() })
	s.AddCloser(zap.ReplaceGlobals(l))
}

func (s *scriptr) DebugC(msg string) {
	s.Z.Debug(ui.Cyan(msg))
}

func (s *scriptr) FatalC(msg string, err error) {
	s.Z.Fatal(ui.Red(msg), zap.Error(err))
}

func (s *scriptr) InfoC(msg string, err error) {
	s.Z.Info(ui.Blue(msg), zap.Error(err))
}

func (s *scriptr) WarnC(args ...interface{}) {
		s.Z.Warn(ui.Yello(args))
}