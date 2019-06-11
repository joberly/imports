package imports

import (
	"os"
	"path/filepath"
)

// Imports is a map of Go import paths to the list of source files
// in which those imports are found.
type Imports map[string][]string

// Get the Imports for all Go source files in the given directory.
// This function returns immediately with no Imports data in the
// event of any error.
func Get(dirName string) (Imports, error) {
	// Get the list of Go files in the directory.
	fileList, err := findGoFiles(dirName)
	if err != nil {
		return nil, err
	}

	imps := Imports{}

	// For each Go file, add the file name to the list of files
	// for each import found in that file.
	for _, fileName := range fileList {
		// Open the Go source file.
		path := filepath.Join(dirName, fileName)
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}

		// Read its import ist.
		importList, err := readImports(file)
		if err != nil {
			return nil, err
		}

		// Add this file to the list for each import path.
		for _, importPath := range importList {
			imps[importPath] = append(imps[importPath], fileName)
		}
	}

	return imps, nil
}
