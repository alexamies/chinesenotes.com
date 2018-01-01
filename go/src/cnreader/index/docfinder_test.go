// Test document retrieval
package index

import (
	"cnreader/corpus"
	"cnreader/dictionary"
	"fmt"
	"testing"
)

// Trivial test for document retrieval
func TestFindForKeyword0(t *testing.T) {
	BuildIndex()
	documents := FindForKeyword("你")
	fmt.Println("index.TestFindForKeyword0 ", documents)
}

// Trivial test for loading index
func TestLoadKeywordIndex0(t *testing.T) {
	LoadKeywordIndex()
}

// Trivial test for loading index
func TestFindDocsForKeyword0(t *testing.T) {
	BuildIndex()
	s1 := "海"
	s2 := "\\N"
	hw := dictionary.HeadwordDef{
		Id:          1,
		Simplified:  &s1,
		Traditional: &s2,
		Pinyin:      []string{"hǎi"},
		WordSenses:  &[]dictionary.WordSenseEntry{},
	}
	fileLoader := corpus.FileCorpusLoader{"File"}
	corpusEntryMap := fileLoader.LoadAll(corpus.COLLECTIONS_FILE)
	documents := FindDocsForKeyword(hw, corpusEntryMap)
	fmt.Println("index.TestFindDocsForKeyword0 ", documents)
}
