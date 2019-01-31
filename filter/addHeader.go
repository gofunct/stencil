package filter

import (
	"bytes"
	"github.com/gofunct/stencil"
)

// AddHeader prepends header to each asset's buffer unless it is already
// prefixed with the header.
func AddHeader(header string) func(*stencil.Asset) error {
	return func(asset *stencil.Asset) error {
		if asset.IsText() {
			if bytes.HasPrefix(asset.Bytes(), []byte(header)) {
				return nil
			}
			buffer := bytes.NewBufferString(header)
			buffer.Write(asset.Bytes())
			asset.Buffer = *buffer
		}
		return nil
	}
}
