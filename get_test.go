package imports

import (
	"sort"
	"testing"
)

func TestGet(t *testing.T) {
	imps, err := Get("testdata")
	if err != nil {
		t.Errorf("Error retrieving imports from testdata: %s", err)
		return
	}

	// Expected output, file list sorted manually.
	impsExp := Imports{
		"fmt": []string{"a.go", "b.go", "read_test.go"},
		"github.com/user/pkg":   []string{"read_test.go"},
		"github.com/user2/pkg2": []string{"read_test.go"},
		"github.com/user3/pkg3": []string{"read_test.go"},
		"github.com/user4/pkg4": []string{"read_test.go"},
		"github.com/user6/pkg6": []string{"a.go"},
		"github.com/user7/pkg7": []string{"b.go"},
		"io":      []string{"read_test.go"},
		"os":      []string{"read_test.go"},
		"path":    []string{"a.go"},
		"runtime": []string{"b.go", "read_test.go"},
	}

	for path, fileList := range imps {
		fileListExp := impsExp[path]
		if fileListExp == nil {
			t.Errorf("unexpected path %s", path)
			continue
		}

		sort.StringSlice(fileListExp).Sort()
		for i, fileName := range fileList {
			if i >= len(fileListExp) {
				t.Errorf("unexpected file %s", fileName)
				continue
			}
			if fileName != fileListExp[i] {
				t.Errorf("file name mismatch: expected %s actual %s",
					fileListExp[i], fileName)
			}
		}
	}
}
