package imports

import (
	"os"
	"testing"
)

// TestReadImports tests good path read of imports from a test Go source file.
func TestReadImports(t *testing.T) {
	r, err := os.Open("read_test_file.txt")
	if err != nil {
		t.Fatal(err)
	}

	list, err := readImports(r)
	if err != nil {
		t.Error(err)
	}

	// Expected list of imports in correct order.
	listExpected := []string{
		"github.com/user/pkg",
		"github.com/user2/pkg2",
		"github.com/user3/pkg3",
		"github.com/user4/pkg4",
		"fmt",
		"io",
		"runtime",
		"os",
	}

	// Check iist against expected list.
	// Report error for extra imports and any import in the list
	// that doesn't match the expected list.
	for i := range list {
		if i >= len(listExpected) {
			t.Errorf("extra import found: %s", list[i])
		} else if list[i] != listExpected[i] {
			t.Errorf("import mismatch: actual %s expected %s", list[i], listExpected[i])
		}
	}
}

// TestReadImportsNil tests for graceful error when specifying a nil Reader.
func TestReadImportsNil(t *testing.T) {
	_, err := readImports(nil)
	if err == nil {
		t.Errorf("no error returned for nil Reader")
	}
}
