package filter

import (
	"github.com/gofunct/stencil/util"
	"strings"

	"github.com/gofunct/stencil/cmd/pipe"
)

// ReplacePath replaces the leading part of a path in all assets.
//
//      ReplacePath("src/", "dist/")
//
// This should be used before the Write() filter.
func ReplacePath(from string, to string) func(*pipe.Asset) error {
	return func(asset *pipe.Asset) error {
		oldPath := asset.WritePath
		if !strings.HasPrefix(oldPath, from) {
			return nil
		}
		asset.WritePath = to + oldPath[len(from):]
		if Verbose {
			util.Debug("pipeline", "ReplacePath %s => %s\n", oldPath, asset.WritePath)
		}
		return nil
	}
}
