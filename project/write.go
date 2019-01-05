package project

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

//Write writes to a file and returns an error

// WriteToFile writes r to file with path only
// if file/directory on given path doesn't exist.
func writeToFile(path string, r io.Reader) error {
	if CheckPathExists(path) {
		return fmt.Errorf("%v already exists", path)
	}

	dir := filepath.Dir(path)
	if dir != "" {
		if err := os.MkdirAll(dir, 0777); err != nil {
			return err
		}
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, r)
	return err
}

func (p *Project) WriteStringToFile(path string, s string) error {
	return writeToFile(path, strings.NewReader(s))
}
