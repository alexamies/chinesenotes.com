/*
 * Unit tests for the replace package
 */
package replace

import (
	"cnreader/library"
	"testing"
	"time"
)

func TestFind0(t *testing.T) {
	expr := "見"
	emptyLoader := library.EmptyLibraryLoader{"Empty"}
	dateUpdated := time.Now().Format("2006-01-02")
	lib := library.Library{
		Title: "Library",
		Summary: "Top level collection in the Library",
		DateUpdated: dateUpdated,
		TargetStatus: "public",
		Loader: emptyLoader,
	}
	Find(expr, lib)
}

func TestFind1(t *testing.T) {
	expr := "見"
	mockLoader := library.MockLibraryLoader{"Mock"}
	dateUpdated := time.Now().Format("2006-01-02")
	lib := library.Library{
		Title: "Library",
		Summary: "Top level collection in the Library",
		DateUpdated: dateUpdated,
		TargetStatus: "public",
		Loader: mockLoader,
	}
	Find(expr, lib)
}
