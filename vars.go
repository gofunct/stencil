package stencil

import (
	"github.com/gofunct/stencil/pkg/wait"
	"github.com/mgutz/minimist"
	"time"
)

const (
	defaultWatchDelay = 1200 * time.Millisecond
)

var (
	watching           bool
	help               bool
	verbose            bool
	version            bool
	deprecatedWarnings bool

	// WaitMs is the default time (1500 ms) to debounce task events in watch mode.
	Wait            time.Duration
	runnerWaitGroup = &wait.WaitGroupN{}
	waitExit        bool
	argm            minimist.ArgMap
	wd              string
	watchDelay      = defaultWatchDelay
)
