package filter

import (
	"bytes"
	"github.com/gofunct/stencil"
)

// Cat concatenates all assets with a join string. Cat clears all assets
// from the pipeline replacing it with a single asset of the concatenated value.
func Cat(join string, dest string) func(*stencil.Pipeline) error {
	return func(pipeline *stencil.Pipeline) error {
		var buffer bytes.Buffer
		for i, asset := range pipeline.Assets {
			if i > 0 {
				buffer.WriteString(join)
			}
			buffer.Write(asset.Bytes())
		}

		// removes existing assets
		pipeline.Truncate()

		// add new asset for the concatenated buffer
		asset := &stencil.Asset{WritePath: dest}
		asset.Write(buffer.Bytes())
		pipeline.AddAsset(asset)
		return nil
	}
}
