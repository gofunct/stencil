package runtime

import (
	"github.com/gofunct/stencil/pkg/runtime/plugins"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type Project interface {
	GenerateProject(rootDir, pkgName string) error
}

// New creates a module.Generator instance.
func NewProject(plugin, path string, cmd *cobra.Command) (Project, error) {
	var r = new(Runtime)
	r.FS = NewFileSystem(path)
	if err := r.initConfig(cmd); err != nil {
		return nil, errors.Wrapf(err, "%s", "failed to setup runtime viper config")
	}
	if err := r.mapFromDefaultConfig(path); err != nil {
		return nil, errors.Wrapf(err, "%s", "failed to merge config with template funcmap")
	}

	switch plugin {
	case "certs":
		r.TmplFS = plugins.Certs
	case "grpc":
		r.TmplFS = plugins.Grpc
	default:
		r.TmplFS = plugins.Init
	}

	return r, nil
}
