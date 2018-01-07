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
	expressions := []Expression{expr}
	emptyLoader := library.EmptyLibraryLoader{"Empty"}
	dateUpdated := time.Now().Format("2006-01-02")
	lib := library.Library{
		Title: "Library",
		Summary: "Top level collection in the Library",
		DateUpdated: dateUpdated,
		TargetStatus: "public",
		Loader: emptyLoader,
	}
	results := FindAndReplace(expressions, lib)
	fmt.Printf("replace.TestFind0: results = %v\n", results)
	if len(results) != 0 {
		t.Error("replace.TestFind0: len(results) != 0")
	}
}

func TestFind1(t *testing.T) {
	expr := Expression{"你好", "", false}
	expressions := []Expression{expr}
	mockLoader := library.MockLibraryLoader{"Mock"}
	dateUpdated := time.Now().Format("2006-01-02")
	lib := library.Library{
		Title: "Library",
		Summary: "Top level collection in the Library",
		DateUpdated: dateUpdated,
		TargetStatus: "public",
		Loader: mockLoader,
	}
	results := FindAndReplace(expressions, lib)
	fmt.Printf("replace.TestFind0: results = %v\n", results)
	if len(results) != 1 {
		t.Error("replace.TestFind1: len(results) != 1")
	}
}

func TestWriteReport0(t *testing.T) {
	results := []Result{}
	WriteReport(results)
}

func TestUnmarshalExp(t *testing.T) {
	jsonStr := `{"Expressions":[{"Find":"你好","Replacement":"再见","Replace":true}]}`
	_, err := UnmarshalExp(jsonStr)
	if err != nil {
		t.Error("replace.TestUnmarshalExp: ", err)
	}
}