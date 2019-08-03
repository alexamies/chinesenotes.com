/*
 * Unit tests for corpus package
 */
package corpus

import (
	"fmt"
	"github.com/alexamies/cnreader/config"
	"testing"
)

func init() {
	config.SetProjectHome("../../../..")
}

// Trivial test to load a collection file
func TestLoadAll0(t *testing.T) {
	fmt.Printf("corpus.TestLoadAll: Begin unit test\n")
	loader := EmptyCorpusLoader{"File"}
	corpusEntryMap := loader.LoadAll(COLLECTIONS_FILE)
	if len(corpusEntryMap) != 0 {
		t.Error("corpus.TestLoadAll0: Non zero num. corpus entries found")
	}
}

// Easy test to load a collection file
func TestLoadAll1(t *testing.T) {
	fmt.Printf("corpus.TestLoadAll1: Begin unit test\n")
	loader := MockCorpusLoader{"File"}
	corpusEntryMap := loader.LoadAll(COLLECTIONS_FILE)
	if len(corpusEntryMap) != 1 {
		t.Error("corpus.TestLoadAll1: No corpus entries found")
	}
	_, ok := corpusEntryMap["raw_file.txt"]
	if !ok {
		fmt.Printf("corpus.TestLoadAll1: corpusEntryMap %v\n", corpusEntryMap)
		t.Error("corpus.TestLoadAll1: Corpus entry not found")
	}
}

func TestLoadAll2(t *testing.T) {
	fmt.Printf("corpus.TestLoadAll2: Begin unit test\n")
	fileLoader := FileCorpusLoader{"File"}
	corpusEntryMap := fileLoader.LoadAll(COLLECTIONS_FILE)
	if len(corpusEntryMap) == 0 {
		t.Error("corpus.TestLoadAll: No corpus entries found")
	} else {
		for _, v := range corpusEntryMap {
			entry := corpusEntryMap[v.RawFile]
			fmt.Printf("corpus.TestLoadAll2: first entry: %v\n", entry)
			break
		}
	}
	fmt.Printf("corpus.TestLoadAll: End unit test\n")
}

// Test reading of files for HTML conversion
func TestCollections(t *testing.T) {
	fmt.Printf("corpus.TestCollections: Begin unit test\n")
	collections := loadCorpusCollections(COLLECTIONS_FILE)
	if len(collections) == 0 {
		t.Error("No collections found")
	} else {
		genre := "Confucian"
		if collections[0].Genre != genre {
			t.Error("Expected genre ", genre, ", got ",
				collections[0].Genre)
		}
	}
	fmt.Printf("corpus.TestCollections: End unit test\n")
}

// Test reading of corpus files
func TestLoadCollection0(t *testing.T) {
	fmt.Printf("corpus.TestLoadCollection0: Begin unit test\n")
	emptyLoader := EmptyCorpusLoader{"Empty"}
	corpusEntries := emptyLoader.LoadCollection("literary_chinese_prose.csv", "")
	if len(corpusEntries) != 0 {
		t.Error("Non zero corpus entries found")
	}
	fmt.Printf("corpus.TestLoadCollection0: End unit test\n")
}

// Test reading of corpus files
func TestLoadCollection1(t *testing.T) {
	fmt.Printf("corpus.TestLoadCollection1: Begin unit test\n")
	mockLoader := MockCorpusLoader{"Mock"}
	corpusEntries := mockLoader.LoadCollection("literary_chinese_prose.csv", "")
	if len(corpusEntries) != 1 {
		t.Error("Num corpus entries found != 1")
	}
	fmt.Printf("corpus.TestLoadCollection1: End unit test\n")
}

// Test reading of corpus files
func TestLoadCollection2(t *testing.T) {
	fileLoader := FileCorpusLoader{"File"}
	corpusEntries := fileLoader.LoadCollection("literary_chinese_prose.csv", "")
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
func TestReadIntroFile(t *testing.T) {
	ReadIntroFile("erya00.txt")
}
