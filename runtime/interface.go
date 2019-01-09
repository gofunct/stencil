package scripts

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
)

type CloudScripter interface {
	Scripter
	Gcloud(args ...string) []byte
	Aws(args ...string) []byte
	Kubectl(args ...string) []byte
}

type RunFunc func(command *cobra.Command, args []string)

type Scripter interface {
	WithArgs(args ...string)
	WithName(args ...string)
	MergeV(v *viper.Viper)
	MergeD(map[string]interface{})
	Execute(ctx *context.Context, args ...string) error
	io.ReadCloser
	CloseIfNotRunnable()
}

type DevScripter interface {
	Scripter
	Make(args ...string) []byte
	Docker(args ...string) []byte
	Git(args ...string) []byte
	Bash(args ...string) []byte
	Go(args ...string) []byte
	Stencil(args ...string) []byte
	Gex(args ...string) []byte
	Protoc(args ...string) []byte
	Terraform(args ...string) []byte
}

type TmplScripter interface {
	Scripter
	WithTemplate(string)
	WithData(v *viper.Viper) map[string]interface{}
}
