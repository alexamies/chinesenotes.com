// Unit tests for query parsing functions
package find

import (
	"log"
	"testing"
)

// Test trivial query with empty dictionary
func TestFindWordsByEnglish(t *testing.T) {
	log.Printf("TestFindWordsByEnglish1: Begin unit tests\n")
	senses, error := findWordsByEnglish("hello")
	log.Printf("TestFindWordsByEnglish1: senses: %v, error: %v\n", senses, error)
}

// Test trivial query with empty dictionary
func TestLoadDict(t *testing.T) {
	wdict, err := LoadDict()
	if err != nil {
		t.Error("TestLoadDict: encountered error: ", err)
		return
	}
	if len(wdict) == 0 {
		t.Error("TestLoadDict: len(wdict) == 0")
	}
	w1 := wdict["猴"]
	if w1.HeadwordId == 0 {
		t.Error("TestLoadDict: w.HeadwordId == 0")
	}
	if w1.Pinyin != "hóu" {
		t.Error("TestLoadDict: w1.Pinyin != hóu", w1.Pinyin)
	}
	w2 := wdict["與"]
	if w2.HeadwordId == 0 {
		t.Error("TestLoadDict: w.HeadwordId == 0")
	}
	if w2.Pinyin == "" {
		t.Error("TestLoadDict: w2.Pinyin == ''")
	}
	w3 := wdict["來"]
	if len(w3.Senses) < 2 {
		t.Error("len(w3.Senses) < 2, ", len(w3.Senses))
	}
	w4 := wdict["發"]
	if len(w4.Senses) < 2 {
		t.Error("len(w4.Senses) < 2, ", len(w4.Senses))
	}
}