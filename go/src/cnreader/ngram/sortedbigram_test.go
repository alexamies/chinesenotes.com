// Test sorting of bigram frequencies
package ngram

import (
	"cnreader/dictionary"
	"fmt"
	"testing"
)

// Test basic Bigram functions
func TestSortedBFM(t *testing.T) {
	fmt.Printf("TestSortedBFM: Begin unit test\n")
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
		HeadwordDef1: &hw1, 
		HeadwordDef2: &hw2,
		Example: "",
		ExFile: "",
		ExDocTitle: "",
		ExColTitle: "",
	}
	bm := BigramFreqMap{}
	bm.PutBigram(b1)
	bm.PutBigram(b1)
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
		HeadwordDef1: &hw1,
		HeadwordDef2: &hw3,
		Example: "",
		ExFile: "",
		ExDocTitle: "",
		ExColTitle: "",
	}
	bm.PutBigram(b2)
	sbf := SortedFreq(bm)
	r1 := len(sbf)
	e1 := 2
	if r1 != e1 {
		t.Error("TestSortedBFM, expected ", e1, " got, ", r1)
	}
	r2 := sbf[0].Frequency
	e2 := 2
	if r2 != e2 {
		t.Error("TestSortedBFM, expected ", e2, " got, ", r2, "sbf[0]", sbf[0])
	}
	r3 := sbf[0].BigramVal.HeadwordDef1.Simplified
	e3 := "蓝"
	if r3 != e3 {
		t.Error("TestSortedBFM, expected ", e3, " got, ", r3, "sbf[0]", sbf[0])
	}
}
