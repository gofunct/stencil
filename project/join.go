package project

import "path/filepath"

//Join returns a new path

// Join joins path elements to the path.
func (p Path) Join(elem ...string) Path {
	return Path(filepath.Join(append([]string{p.ConvertToString()}, elem...)...))
}
