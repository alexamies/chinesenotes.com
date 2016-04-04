package alignment

import (
	"cnreader/config"
	"cnreader/corpus"
	"cnreader/dictionary"
	"fmt"
	"testing"
)

func init() {
	config.SetProjectHome("../../../..")
}

// Basic test
func TestAddCognate1(t *testing.T) {
	fmt.Printf("TestAddCognate1: Begin unit tests\n")
	corpusEntry := corpus.CorpusEntry{}
	testTitle := "Test Title"
	corpusEntry.Title = testTitle
	corpEntryCogs := NewCorpEntryCognates(corpusEntry)
	ws1 := dictionary.WordSenseEntry{
		Id: 1,
		Simplified: "佛", 
		Traditional: "\\N",
		Pinyin: "Fó",
		Grammar: "proper noun",
	}
	corpEntryCogs.AddCognate(&ws1)
	properNouns := corpEntryCogs.GetProperNouns()
	result := len(properNouns)
	expected := 1
	if result != expected {
		t.Error("TestAddCognate1: Expected ", expected, ", got ", result)
	}
	if corpEntryCogs.CorpEntry.Title != testTitle {
		t.Error("TestAddCognate1: Expected ", testTitle, ", got ",
			corpEntryCogs.CorpEntry.Title)
	}
}

// Test that the same cognate is not added multiple times
func TestAddCognate2(t *testing.T) {
	corpusEntry := corpus.CorpusEntry{}
	corpEntryCogs := NewCorpEntryCognates(corpusEntry)
	ws1 := dictionary.WordSenseEntry{
		Id: 1,
		Simplified: "佛", 
		Traditional: "\\N",
		Pinyin: "Fó",
		Grammar: "proper noun",
	}
	corpEntryCogs.AddCognate(&ws1)
	corpEntryCogs.AddCognate(&ws1)
	properNouns := corpEntryCogs.GetProperNouns()
	result := len(properNouns)
	expected := 1
	if result != expected {
		t.Error("TestAddCognate2: Expected ", expected, ", got ", result)
	}
}

// Test numeric cognate
func TestAddCognate3(t *testing.T) {
	corpusEntry := corpus.CorpusEntry{}
	corpEntryCogs := NewCorpEntryCognates(corpusEntry)
	ws1 := dictionary.WordSenseEntry{
		Id: 1,
		Simplified: "四圣谛", 
		Traditional: "四聖諦",
		Pinyin: "Sì Shèng Dì",
		Grammar: "proper noun",
	}
	corpEntryCogs.AddCognate(&ws1)
	corpEntryCogs.AddCognate(&ws1)
	properNouns := corpEntryCogs.GetProperNouns()
	result := len(properNouns)
	expected := 1
	if result != expected {
		t.Error("TestAddCognate2: Expected ", expected, ", got ", result)
	}
}