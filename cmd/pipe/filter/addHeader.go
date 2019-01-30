package filter

import (
	"bytes"
	"github.com/gofunct/stencil/cmd/pipe"
)

// AddHeader prepends header to each asset's buffer unless it is already
// prefixed with the header.
func AddHeader(header string) func(*pipe.Asset) error {
	return func(asset *pipe.Asset) error {
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
