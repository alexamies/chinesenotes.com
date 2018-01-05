/*
 * Unit tests for the replace package
 */
package replace

import (
	"cnreader/library"
	"fmt"
	"testing"
	"time"
)

func TestFind0(t *testing.T) {
	expr := Expression{"見", "", false}
	emptyLoader := library.EmptyLibraryLoader{"Empty"}
	dateUpdated := time.Now().Format("2006-01-02")
	lib := library.Library{
		Title: "Library",
		Summary: "Top level collection in the Library",
		DateUpdated: dateUpdated,
		TargetStatus: "public",
		Loader: emptyLoader,
	}
	results := FindAndReplace(expr, lib)
	fmt.Printf("corpus.TestFind0: results = %v\n", results)
	if len(results) != 0 {
		t.Error("corpus.TestFind0: len(results) != 0")
	}
}

func TestFind1(t *testing.T) {
	expr := Expression{"你好", "", false}
	mockLoader := library.MockLibraryLoader{"Mock"}
	dateUpdated := time.Now().Format("2006-01-02")
	lib := library.Library{
		Title: "Library",
		Summary: "Top level collection in the Library",
		DateUpdated: dateUpdated,
		TargetStatus: "public",
		Loader: mockLoader,
	}
	results := FindAndReplace(expr, lib)
	fmt.Printf("corpus.TestFind0: results = %v\n", results)
	if len(results) != 1 {
		t.Error("corpus.TestFind1: len(results) != 1")
	}
}
