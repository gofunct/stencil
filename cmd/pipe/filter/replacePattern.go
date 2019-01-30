package filter

import (
	"regexp"

	"github.com/gofunct/stencil/cmd/pipe"
)

// ReplacePattern replaces the leading part of a path in all assets.
//
//      ReplacePath("views/", "dist/views")
//
// This should be used before the Write() filter.
func ReplacePattern(pattern, repl string) func(*pipe.Asset) error {
	re := regexp.MustCompile(pattern)
	return func(asset *pipe.Asset) error {
		if asset.IsText() {
			s := asset.String()
			if s != "" {
				asset.RewriteString(re.ReplaceAllString(s, repl))
			}
		}
		return nil
	}
}
