package filter

import (
	"github.com/gofunct/stencil/cmd/pipe"
	"github.com/gofunct/stencil/glob"
	"io/ioutil"
)

// Load loads all the files from glob patterns and creates the initial
// asset array for a pipeline. This loads the entire contents of the file, binary
// or text, into a buffer. Consider creating your own loader if dealing
// with large files.
func Load(patterns ...string) func(*pipe.Pipeline) error {
	return func(pipeline *pipe.Pipeline) error {
		fileAssets, _, err := glob.Glob(patterns)
		if err != nil {
			return err
		}

		for _, info := range fileAssets {
			if !info.IsDir() {
				data, err := ioutil.ReadFile(info.Path)
				if err != nil {
					return err
				}
				asset := &pipe.Asset{Info: info}
				asset.Write(data)
				asset.WritePath = info.Path
				pipeline.AddAsset(asset)
			}
		}
		return nil
	}
}
