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
	expected := "literary_chinese_prose.csv"
	if collections[0].CollectionFile != expected {
		t.Error("Expected entry ", expected, ", got ", collections[0])
	}
}

// Test reading of corpus files
func TestCorpusEntries(t *testing.T) {
	corpusEntries := CorpusEntries("../../../../data/corpus/literary_chinese_prose.csv")
	if len(corpusEntries) == 0 {
		t.Error("No corpus entries found")
	}
	if corpusEntries[0].RawFile != "classical_chinese_text-raw.html" {
		t.Error("Expected entry classical_chinese_text-raw.html, got ",
			corpusEntries[0].RawFile)
	}
	if corpusEntries[0].GlossFile != "classical_chinese_text.html" {
		t.Error("Expected entry classical_chinese_text.html, got ",
			corpusEntries[0].GlossFile)
	}
}

// Test generating collection file
func TestWriteCollectionFile(t *testing.T) {
	WriteCollectionFile("erya.csv")
}