package project

import "path/filepath"

//Trim returns a string

// TrimSrcPath trims at the beginning of absPath the srcPath.
func (p *Project) TrimSrcPath(absPath, srcPath string) string {
	relPath, err := filepath.Rel(srcPath, absPath)
	if err != nil {
		exit(err)
	}
	return relPath
}
