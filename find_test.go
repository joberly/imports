package imports

import "testing"

func TestFindGoFiles(t *testing.T) {
	fileList, err := findGoFiles("testdata")
	if err != nil {
		t.Error(err)
		return
	}

	// Map of expected file names and the count of
	// how many times they are found in the list.
	// Expect each key in the map to have a count of one.
	fileListExp := map[string]int{"a.go": 0, "b.go": 0, "read_test.go": 0}

	// Check that each file found was expected and count it.
	for _, filename := range fileList {
		if _, ok := fileListExp[filename]; !ok {
			t.Errorf("unexpected file found: %s", filename)
		}
		count := fileListExp[filename]
		fileListExp[filename] = count + 1
	}

	// Check if the function found files more than once.
	for filename, count := range fileListExp {
		if count > 1 {
			t.Errorf("file found more than once: %s count %d",
				filename, count)
		}
	}
}
