package filter

import (
	"github.com/gofunct/stencil"
	"github.com/gofunct/stencil/pkg/print"
	"strings"
)

// ReplacePath replaces the leading part of a path in all assets.
//
//      ReplacePath("src/", "dist/")
//
// This should be used before the Write() filter.
func ReplacePath(from string, to string) func(*stencil.Asset) error {
	return func(asset *stencil.Asset) error {
		oldPath := asset.WritePath
		if !strings.HasPrefix(oldPath, from) {
			return nil
		}
		asset.WritePath = to + oldPath[len(from):]
		if Verbose {
			print.Debug("pipeline", "ReplacePath %s => %s\n", oldPath, asset.WritePath)
		}
		return nil
	}
}
