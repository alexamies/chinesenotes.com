// Test counting of bigram frequencies
package ngram

import (
	"cnreader/dictionary"
	"fmt"
	"testing"
)

// Test basic Bigram functions
func TestBigramFreqMap(t *testing.T) {
	fmt.Printf("TestBigramFreqMap: Begin unit test\n")
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
	b1 := Bigram{
		HeadwordDef1: hw1, 
		HeadwordDef2: hw2,
		Example: "",
		ExFile: "",
		ExDocTitle: "",
		ExColTitle: "",
	}
	bm := NewBigramFreqMap()
	bm.PutBigram(b1)
	bm.PutBigram(b1)
	r1 := bm.GetBigram(b1)
	e1 := 2
	if r1.Frequency != e1 {
		t.Error("TestBigramFreqMap, expected ", e1, " got, ", r1.Frequency)
	}
	ws3 := dictionary.WordSenseEntry{
		Id: 3,
		Simplified: "海", 
		Traditional: "\\N",
		Pinyin: "hǎi",
		Grammar: "noun",
	}
	hw3 := dictionary.HeadwordDef{
		Id: 3,
		Simplified: "海",
		Traditional: "\\N",
		Pinyin: []string{},
		WordSenses: []dictionary.WordSenseEntry{ws3},
	}
	b2 := Bigram{
		HeadwordDef1: hw1,
		HeadwordDef2: hw3,
		Example: "",
		ExFile: "",
		ExDocTitle: "",
		ExColTitle: "",
	}
	r2 := bm.GetBigram(b2)
	e2 := 0
	if r2.Frequency != e2 {
		t.Error("TestBigramFreqMap, expected ", e2, " got, ", r2.Frequency)
	}
}
