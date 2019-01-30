package filter

import (
	"github.com/gofunct/stencil/cmd/pipe"
	"github.com/gofunct/stencil/util"
)

// Trace traces an asset, printing key properties of asset to the console.
func Trace() func(*pipe.Asset) error {
	return func(asset *pipe.Asset) error {
		util.Debug("filter", asset.Dump())
		return nil
	}
}
