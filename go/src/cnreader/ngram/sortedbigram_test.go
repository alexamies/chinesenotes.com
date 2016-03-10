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
	hw1 := dictionary.HeadwordDef{
		Id: 1,
		Simplified: "蓝", 
		Traditional: "藍",
		Pinyin: []string{},
		WordSenses: []dictionary.WordSenseEntry{},
	}
	hw2 := dictionary.HeadwordDef{
		Id: 2,
		Simplified: "天",
		Traditional: "\\N",
		Pinyin: []string{},
		WordSenses: []dictionary.WordSenseEntry{},
	}
	b1 := Bigram{hw1, hw2}
	bm := NewBigramFreqMap()
	bm.PutBigram(b1)
	bm.PutBigram(b1)
	hw3 := dictionary.HeadwordDef{
		Id: 3,
		Simplified: "海",
		Traditional: "\\N",
		Pinyin: []string{},
		WordSenses: []dictionary.WordSenseEntry{},
	}
	b2 := Bigram{hw1, hw3}
	bm.PutBigram(b2)
	sbf := SortedFreq(*bm)
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
