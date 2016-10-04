// Test counting of bigram frequencies
package ngram

import (
	"cnreader/dictionary"
	"fmt"
	"testing"
)

// Test basic Bigram functions
func TestMerge(t *testing.T) {
	fmt.Printf("TestMerge: Begin unit test\n")

	// Set up test
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

	// Invoke twice
	bm.PutBigram(b1)
	bm.PutBigram(b1)

	// Second map
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
	bm2 := BigramFreqMap{}
	bm2.PutBigram(b1)
	bm2.PutBigram(b2)

	// Method to test
	bm.Merge(bm2)

	// Check result
	r1 := bm.GetBigram(b1)
	e1 := 3
	if r1.Frequency != e1 {
		t.Error("TestPutBigram, expected ", e1, " got, ", r1.Frequency)
	}
	r2 := bm.GetBigram(b2)
	e2 := 1
	if r2.Frequency != e2 {
		t.Error("TestPutBigram, expected ", e2, " got, ", r2.Frequency)
	}
}

// Test basic Bigram functions
func TestPutBigram(t *testing.T) {

	// Set up test
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

	// Method to test (invoke twice)
	bm.PutBigram(b1)
	bm.PutBigram(b1)

	// Check result
	r1 := bm.GetBigram(b1)
	e1 := 2
	if r1.Frequency != e1 {
		t.Error("TestPutBigram, expected ", e1, " got, ", r1.Frequency)
	}
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
	r2 := bm.GetBigram(b2)
	e2 := 0
	if r2.Frequency != e2 {
		t.Error("TestPutBigram, expected ", e2, " got, ", r2.Frequency)
	}
}

// Test basic Bigram functions
func TestPutBigramFreq(t *testing.T) {

	// Set up test
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
	bf := BigramFreq{*b1, 2}

	// Method to test (invoke twice)
	bm.PutBigramFreq(bf)

	// Check result
	r1 := bm.GetBigram(b1)
	e1 := 3
	if r1.Frequency != e1 {
		t.Error("TestPutBigramFreq, expected ", e1, " got, ", r1.Frequency)
	}
}