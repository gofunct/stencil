package root

import (
	"github.com/spf13/afero"
	"io"
)

type Context interface {
	logS
	configS
	ioS
	debugS
}

type logS interface {
	Exit(args ...interface{})
	Er(error, ...interface{})
	Info(msg string, args ...interface{})
}

type debugS interface {
	Debug(msg string)
}

type configS interface {
	Query() error
	Value(key string) interface{}
	MergeMeta(map[string]interface{})
	GetMeta() map[string]interface{}
	Unmarshal(interface{}) error
	Bytes(interface{}) ([]byte, error)
	GetEnv() ([]string, error)
	SetFs(afero.Fs)
}

// IO contains an input reader, an output writer and an error writer.
type ioS interface {
	In() io.Reader
	Out() io.Writer
	Err() io.Writer
	Close() error
}

type Failer interface {
	Failed() error
}