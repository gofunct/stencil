package filter

import (
	"github.com/gofunct/gofs"
	"github.com/gofunct/stencil"
)

// Trace traces an asset, printing key properties of asset to the console.
func Trace() func(*stencil.Asset) error {
	return func(asset *stencil.Asset) error {
		gofs.Debug("filter", asset.Dump())
		return nil
	}
}
