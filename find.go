package imports

import (
	"os"
	"path/filepath"
)

// findGoFiles finds all files with extension ".go" in the given directory
// and returns the base names of those files in the returned string slice.
func findGoFiles(dirname string) ([]string, error) {
	dir, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}

	// Note this will read file info for all files in the directory.
	// Initial assumption here is that a Go project directory
	// has a "reasonable" number of files to limit memory usage.
	var fiList []os.FileInfo
	fiList, err = dir.Readdir(0)
	if err != nil {
		return nil, err
	}

	fileList := make([]string, 0, len(fiList))
	for _, fi := range fiList {
		if !fi.IsDir() && filepath.Ext(fi.Name()) == ".go" {
			fileList = append(fileList, fi.Name())
		}
	}

	return fileList, nil
}
