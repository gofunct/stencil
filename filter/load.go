package filter

import (
	"github.com/gofunct/gofs"
	"github.com/gofunct/stencil"
	"io/ioutil"
)

// Load loads all the files from glob patterns and creates the initial
// asset array for a pipeline. This loads the entire contents of the file, binary
// or text, into a buffer. Consider creating your own loader if dealing
// with large files.
func Load(patterns ...string) func(*stencil.Pipeline) error {
	return func(pipeline *stencil.Pipeline) error {
		fileAssets, _, err := gofs.Glob(patterns)
		if err != nil {
			return err
		}

		for _, info := range fileAssets {
			if !info.IsDir() {
				data, err := ioutil.ReadFile(info.Path)
				if err != nil {
					return err
				}
				asset := &stencil.Asset{Info: info}
				asset.Write(data)
				asset.WritePath = info.Path
				pipeline.AddAsset(asset)
			}
		}
		return nil
	}
}
