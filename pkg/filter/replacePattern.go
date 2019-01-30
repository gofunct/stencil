package filter

import (
	"github.com/gofunct/stencil"
	"regexp"
)

// ReplacePattern replaces the leading part of a path in all assets.
//
//      ReplacePath("views/", "dist/views")
//
// This should be used before the Write() filter.
func ReplacePattern(pattern, repl string) func(*stencil.Asset) error {
	re := regexp.MustCompile(pattern)
	return func(asset *stencil.Asset) error {
		if asset.IsText() {
			s := asset.String()
			if s != "" {
				asset.RewriteString(re.ReplaceAllString(s, repl))
			}
		}
		return nil
	}
}
