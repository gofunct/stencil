package app

import (
	"github.com/gofunct/stencil/pkg/runtime"
	"github.com/gofunct/stencil/pkg/ui"
)

type CreateAppFunc func(*Command) (*App, error)

// App contains dependencies to execute a generator.
type App struct {
	Project runtime.Project
	UI      ui.UI
}
