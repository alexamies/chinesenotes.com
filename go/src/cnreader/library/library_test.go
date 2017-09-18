package corpus

import (
	"fmt"
	"testing"
)

func TestLibrary(t *testing.T) {
	fmt.Printf("library.TestLibrary: Begin unit tests\n")
	l := Library()
	if len(l) == 0 {
		t.Error("No collections found")
	}
}

func TestWriteLibraryFile(t *testing.T) {
	WriteLibraryFile()
}
