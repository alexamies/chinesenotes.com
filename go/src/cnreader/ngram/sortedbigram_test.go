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
	s1 := "蓝"
	s2 := "藍"
	ws1 := dictionary.WordSenseEntry{
		Id: 1,
		Simplified: s1, 
		Traditional: s2,
		Pinyin: "lán",
		Grammar: "adjective",
	}
	hw1 := dictionary.HeadwordDef{
		Id: 1,
		Simplified: &s1, 
		Traditional: &s2,
		Pinyin: []string{},
		WordSenses: &[]dictionary.WordSenseEntry{ws1},
	}
	s3 := "天"
	s4 := "\\N"
	ws2 := dictionary.WordSenseEntry{
		Id: 1,
		Simplified: s3, 
		Traditional: s4,
		Pinyin: "tiān",
		Grammar: "noun",
	}
	hw2 := dictionary.HeadwordDef{
		Id: 2,
		Simplified: &s3, 
		Traditional: &s4,
		Pinyin: []string{},
		WordSenses: &[]dictionary.WordSenseEntry{ws2},
	}
	example := ""
	exFile := ""
	exDocTitle := ""
	exColTitle := ""
	b1 := NewBigram(hw1, hw2, example, exFile, exDocTitle, exColTitle)
	bm := BigramFreqMap{}
	bm.PutBigram(b1)
	bm.PutBigram(b1)
	s5 := "海"
	s6 := "\\N"
	ws3 := dictionary.WordSenseEntry{
		Id: 3,
		Simplified: s5, 
		Traditional: s6,
		Pinyin: "hǎi",
		Grammar: "noun",
	}
	hw3 := dictionary.HeadwordDef{
		Id: 3,
		Simplified: &s5,
		Traditional: &s6,
		Pinyin: []string{},
		WordSenses: &[]dictionary.WordSenseEntry{ws3},
	}
	b2 := NewBigram(hw1, hw3, example, exFile, exDocTitle, exColTitle)
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
	r3 := *sbf[0].BigramVal.HeadwordDef1.Simplified
	e3 := "蓝"
	if r3 != e3 {
		t.Error("TestSortedBFM, expected ", e3, " got, ", r3, "sbf[0]", sbf[0])
	}
}
