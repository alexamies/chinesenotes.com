// Test keyword ordering
package index

import (
	"cnreader/config"
	"cnreader/dictionary"
	"log"
	"testing"
)

func TestGetHeadwordArray0(t *testing.T) {
	keywords := Keywords{}
	hws := GetHeadwordArray(keywords)
	nReturned := len(hws)
	nExpected := 0
	if nReturned != nExpected {
		t.Error("index.TestGetHeadwordArray0: nExpected ", nExpected, " got ",
			nReturned)
	}
}

func TestGetHeadwordArray1(t *testing.T) {
	kw := Keyword{"多", 1.1}
	keywords := Keywords{kw}
	dictionary.ReadDict(config.LUFileNames())
	hws := GetHeadwordArray(keywords)
	nReturned := len(hws)
	nExpected := 1
	if nReturned != nExpected {
		t.Error("index.TestGetHeadwordArray1: nExpected ", nExpected, " got ",
			nReturned)
	}
	nPinyin := len(hws[0].Pinyin)
	nPExpected := 1
	if nPinyin != nPExpected {
		t.Error("index.TestGetHeadwordArray1: nPExpected ", nPExpected, " got ",
			nPinyin)
	}
}

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
// vocab1: term1 = 1*log10(1/1) = 0.0, term2 = 2*log10(2/1) = 0.60
//
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
	topExpected := term2
	if top.Term != topExpected {
		log.Printf("index.TestSortByWeight2 keywords = %v\n", keywords)
		t.Error("index.TestSortByWeight2: topExpected ", topExpected, " got ",
			top)
	}
}
