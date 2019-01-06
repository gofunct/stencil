package app

import (
	"github.com/gofunct/stencil/pkg/runtime"
	"github.com/izumin5210/gex"
	"github.com/spf13/cobra"
	"net/http"
)

type Command struct {
	Cobra       *cobra.Command
	Runtime     *runtime.Runtime
	GexCfg      *gex.Config
	BuildParams func(c *Command, args []string) (interface{}, error)
	TemplateFS  http.FileSystem
}

func NewCommand() *Command {
	return &Command{}
}

func (c *Command) Execute() {

}
