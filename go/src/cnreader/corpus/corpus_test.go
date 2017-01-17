package corpus

import (
	"cnreader/config"
	"fmt"
	"testing"
)

func init() {
	config.SetProjectHome("../../../..")
}

// Test reading of files for HTML conversion
func TestCollections(t *testing.T) {
	fmt.Printf("corpus.TestCollections: Begin unit tests\n")
	collections := Collections()
	if len(collections) == 0 {
		t.Error("No collections found")
	} else {
		genre := "Confucian"
		if collections[0].Genre != genre {
			t.Error("Expected genre ", genre, ", got ",
				collections[0].Genre)
		}
	}
}

// Test reading of corpus files
func TestCorpusEntries(t *testing.T) {
	corpusEntries := CorpusEntries("../../../../data/corpus/literary_chinese_prose.csv", "")
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

// Trivial test to look up a collection file
func TestLoadCorpusEntries(t *testing.T) {
	loadCorpusEntries()
	if len(corpusEntryMap) == 0 {
		t.Error("corpus.TestLoadCorpusEntries: No corpus entries found")
	} else {
		for _, v := range corpusEntryMap {
			entry := GetCorpusEntry(v.RawFile)
			fmt.Printf("corpus.TestLoadCorpusEntries: first entry: %v\n", entry)
			break
		}
	}
}

// Test generating collection file
func TestReadIntroFile(t *testing.T) {
	ReadIntroFile("erya00.txt")
}

// Test generating collection file
func TestWriteCollectionFile(t *testing.T) {
	WriteCollectionFile("erya.csv", "analysis/erya-analysis-test.html")
}