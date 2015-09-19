package config

import (
	"fmt"
	"testing"
)

// Test reading of files for HTML conversion
func TestGetHTMLConversions(t *testing.T) {
	fmt.Printf("TestGetHTMLConversions: Begin unit tests\n")
	conversions := GetHTMLConversions()
	fmt.Printf("TestGetHTMLConversions: # conversions: %d\n", len(conversions))
	if len(conversions) == 0 {
		t.Error("No conversions found")
	}
	if conversions[0].SrcFile != "../corpus/classical_chinese-raw.html" {
		t.Error("Expected source file classical_chinese-raw.html, got ",
			conversions[0].SrcFile)
	}
	if conversions[0].DestFile != "classical_chinese.html" {
		t.Error("Expected dest file classical_chinese.html, got ",
			conversions[0].DestFile)
	}
}

