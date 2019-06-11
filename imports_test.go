package imports

import (
	"os"
	"testing"
)

func TestReadImports(t *testing.T) {
	r, err := os.Open("imports_test_file.txt")
	if err != nil {
		t.Fatal(err)
	}

	list, err := readImports(r)
	if err != nil {
		t.Error(err)
	}

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

	for i := range list {
		if i >= len(listExpected) {
			t.Errorf("extra import found: %s", list[i])
		} else if list[i] != listExpected[i] {
			t.Errorf("import mismatch: actual %s expected %s", list[i], listExpected[i])
		}
	}
}

func TestReadImportsNil(t *testing.T) {
	_, err := readImports(nil)
	if err == nil {
		t.Errorf("no error returned for nil Reader")
	}
}
