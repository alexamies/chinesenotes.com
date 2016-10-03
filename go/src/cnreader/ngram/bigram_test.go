// Test bigram functions
package ngram

import (
	"cnreader/dictionary"
	"fmt"
	"testing"
)

// Test basic Bigram functions
func TestBigram(t *testing.T) {
	fmt.Printf("TestBigram: Begin unit test\n")
	hw1 := dictionary.HeadwordDef{
		Id: 1,
		Simplified: "蓝", 
		Traditional: "藍",
		Pinyin: []string{},
		WordSenses: &[]dictionary.WordSenseEntry{},
	}
	hw2 := dictionary.HeadwordDef{
		Id: 2,
		Simplified: "天", 
		Traditional: "\\N",
		Pinyin: []string{},
		WordSenses: &[]dictionary.WordSenseEntry{},
	}
	b := Bigram{
		HeadwordDef1: &hw1,
		HeadwordDef2: &hw2,
		Example: "",
		ExFile: "",
		ExDocTitle: "",
		ExColTitle: "",
	}
	r := b.Traditional()
	e := "藍天"
	if r != e {
		t.Error("TestBigram, expected ", e, " got, ", r)
	}
}
