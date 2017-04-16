// Test keyword ordering
package index

import (
	"log"
	"testing"
)

// Trivial test for keyword ordering
func TestSortByWeight0(t *testing.T) {
	vocab := map[string]int{}
	keywords := SortByWeight(vocab)
	nReturned := len(keywords)
	nExpected := 0
	if nReturned != nExpected {
		t.Error("index.TestSortByWeight0: nExpected ", nExpected, " got ",
			nReturned)
	}
}

// Simple test for keyword ordering
func TestSortByWeight1(t *testing.T) {
	df := NewDocumentFrequency()
	term0 := "好"
	vocab := map[string]int{
		term0: 1,
	}
	df.AddVocabulary(vocab)
	completeDF = df
	keywords := SortByWeight(vocab)
	nReturned := len(keywords)
	nExpected := 1
	if nReturned != nExpected {
		t.Error("index.TestSortByWeight1: nExpected ", nExpected, " got ",
			nReturned)
	}
	top := keywords[0]
	topExpected := term0
	if top.Term != topExpected {
		t.Error("index.TestSortByWeight1: topExpected ", topExpected, " got ",
			top)
	}
}

// Easy test for keyword ordering
func TestSortByWeight2(t *testing.T) {
	df := NewDocumentFrequency()
	term0 := "你"
	term1 := "好"
	term2 := "嗎"
	vocab0 := map[string]int{
		term0: 1,
		term1: 2,
	}
	df.AddVocabulary(vocab0)
	vocab1 := map[string]int{
		term1: 1,
		term2: 2,
	}
	df.AddVocabulary(vocab1)
	completeDF = df
	keywords := SortByWeight(vocab1)
	top := keywords[0]
	topExpected := term1
	if top.Term != topExpected {
		log.Printf("index.TestSortByWeight2 keywords = %v\n", keywords)
		t.Error("index.TestSortByWeight2: topExpected ", topExpected, " got ",
			top)
	}
}
