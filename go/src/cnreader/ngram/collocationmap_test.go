// Test counting of bigram frequencies
package ngram

import (
	"cnreader/dictionary"
	"fmt"
	"testing"
)

// Test basic Bigram functions
func TestCMPutBigram(t *testing.T) {
	fmt.Printf("TestCMPutBigram: Begin unit test\n")

	// Data to test
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
	cm := CollocationMap{}

	// Method being tested
	cm.PutBigram(hw1.Id, b1)
	cm.PutBigram(hw1.Id, b1)

	// check result
	bfm := cm[hw1.Id]

	r1 := bfm.GetBigram(b1)
	e1 := 2
	if r1.Frequency != e1 {
		t.Error("TestCMPutBigram, expected ", e1, " got, ", r1.Frequency)
	}
}

// Test basic Bigram functions
func TestCMPutBigramFreq(t *testing.T) {

	// Data to test
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
	cm := CollocationMap{}
	cm.PutBigram(hw1.Id, b1)
	bf := BigramFreq{*b1, 3}

	// Method being tested
	cm.PutBigramFreq(hw1.Id, bf)

	// check result
	bfm := cm[hw1.Id]

	r1 := bfm.GetBigram(b1)
	e1 := 4
	if r1.Frequency != e1 {
		t.Error("TestCMPutBigramFreq, expected ", e1, " got, ", r1.Frequency)
	}
}


// Test basic Bigram functions
func TestMergeCollocationMap(t *testing.T) {

	// Data to test
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
	cm := CollocationMap{}
	cm.PutBigram(hw1.Id, b1)
	cm2 := CollocationMap{}
	cm2.PutBigram(hw1.Id, b1)

	// Method being tested
	cm.MergeCollocationMap(cm2)

	// check result
	bfm := cm[hw1.Id]

	r1 := bfm.GetBigram(b1)
	e1 := 2
	if r1.Frequency != e1 {
		t.Error("TestCMPutBigramFreq, expected ", e1, " got, ", r1.Frequency)
	}
}