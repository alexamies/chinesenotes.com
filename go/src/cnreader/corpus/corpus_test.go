package corpus

import (
	"fmt"
	"testing"
)

// Test reading of files for HTML conversion
func TestCollections(t *testing.T) {
	fmt.Printf("Collections: Begin unit tests\n")
	collections := Collections()
	if len(collections) == 0 {
		t.Error("No collections found")
	}
	if collections[0] != "literary_chinese_prose.csv" {
		t.Error("Expected fentry html-conversion.csv, got ",
			collections[0])
	}
}

// Test reading of corpus files
func TestCorpusEntries(t *testing.T) {
	corpusEntries := CorpusEntries("../../../../data/corpus/literary_chinese_prose.csv")
	if len(corpusEntries) == 0 {
		t.Error("No corpus entries found")
	}
	if corpusEntries[0].RawFile != "../corpus/classical_chinese_text-raw.html" {
		t.Error("Expected entry ../corpus/classical_chinese_text-raw.html, got ",
			corpusEntries[0].RawFile)
	}
	if corpusEntries[0].GlossFile != "classical_chinese_text.html" {
		t.Error("Expected entry classical_chinese_text.html, got ",
			corpusEntries[0].GlossFile)
	}
}