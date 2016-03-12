// Test sorting of word frequencies
package analysis

import (
	"cnreader/dictionary"
	"cnreader/ngram"
	"fmt"
	"testing"
)

// Test sorting of word frequencies
func TestAddResults(t *testing.T) {
	fmt.Printf("TestAddResults: Begin unit tests\n")
	vocab := map[string]int{"one":1, "three":3, "two":2}
	usage := map[string]string {"one": "one banana"}
	unknown := map[string]int{"x":1}
	ws1 := dictionary.WordSenseEntry{
		Id: 1,
		Simplified: "蓝", 
		Traditional: "藍",
		Pinyin: "lán",
		Grammar: "adjective",
	}
	hw1 := dictionary.HeadwordDef{
		Id: 1,
		Simplified: "蓝", 
		Traditional: "藍",
		Pinyin: []string{},
		WordSenses: []dictionary.WordSenseEntry{ws1},
	}
	ws2 := dictionary.WordSenseEntry{
		Id: 1,
		Simplified: "天", 
		Traditional: "\\N",
		Pinyin: "tiān",
		Grammar: "noun",
	}
	hw2 := dictionary.HeadwordDef{
		Id: 2,
		Simplified: "天",
		Traditional: "\\N",
		Pinyin: []string{},
		WordSenses: []dictionary.WordSenseEntry{ws2},
	}
	b1 := ngram.Bigram{
		HeadwordDef1: hw1, 
		HeadwordDef2: hw2,
		Example: "",
		ExFile: "",
		ExDocTitle: "",
		ExColTitle: "",
	}
	bm := ngram.NewBigramFreqMap()
	bm.PutBigram(b1)
	bm.PutBigram(b1)
	results := CollectionAResults{
		Vocab: vocab,
		Usage: usage,
		BigramFrequencies: *bm,
		WC: 3,
		UnknownChars: unknown,
	}
	moreVocab := map[string]int{"one":1, "three":1, "four":4}
	moreUsage := map[string]string {"two": "two banana"}
	unknown1 := map[string]int{"x":1}
	more := CollectionAResults{
		Vocab: moreVocab,
		Usage: moreUsage,
		BigramFrequencies: *bm,
		WC: 4,
		UnknownChars: unknown1,
	}
	results.AddResults(more)
	r := results.Vocab["three"]
	e := 4
	if r != e {
		t.Error("TestAddResults, three expected ", e, " got, ", r)
	}
	r = results.Vocab["four"]
	e = 4
	if r != e {
		t.Error("TestAddResults, four expected ", e, " got, ", r)
	}
	r = results.WC
	e = 7
	if r != e {
		t.Error("TestAddResults, word count expected ", e, " got, ", r)
	}
	r = results.BigramFrequencies.GetBigram(b1).Frequency
	e = 4
	if r != e {
		t.Error("TestAddResults, bigram count expected ", e, " got, ", r)
	}
}
