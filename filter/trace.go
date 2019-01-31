package filter

import (
	"github.com/gofunct/stencil"
	"github.com/gofunct/stencil/pkg/print"
)

// Trace traces an asset, printing key properties of asset to the console.
func Trace() func(*stencil.Asset) error {
	return func(asset *stencil.Asset) error {
		print.Debug("filter", asset.Dump())
		return nil
	}
}
