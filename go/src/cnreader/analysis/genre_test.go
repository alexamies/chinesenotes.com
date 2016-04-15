// Test sorting of word frequencies
package analysis

import (
	"log"
	"testing"
)

// Test merging of word frequencies
func TestMergeByGenre1(t *testing.T) {
	classicalWF1 := NewWordFreqByGenre("classical")
	classicalWF1.WF["thou"] = 1
	classicalWF2 := NewWordFreqByGenre("classical")
	classicalWF2.WF["thou"] = 2
	genreWF1 := WFArrayByGenre{classicalWF1}
	genreWF := MergeByGenre(genreWF1, classicalWF2)
	wf := genreWF.Get("classical")
	result := wf["thou"]
	expected := 3
	if result != expected {
		t.Error("analysis.TestMergeByGenre1: Expected ", expected, ", got ", result)
	}
}

// Test length of merged results
func TestMergeByGenre2(t *testing.T) {
	log.Printf("analysis.TestMergeByGenre2: Begin ******** \n")	
	classicalWF1 := NewWordFreqByGenre("classical")
	classicalWF1.WF["thou"] = 1
	classicalWF2 := NewWordFreqByGenre("modern")
	classicalWF2.WF["you"] = 2
	genreWF1 := WFArrayByGenre{classicalWF1}
	genreWF := MergeByGenre(genreWF1, classicalWF2)
	result := len(genreWF)
	expected := 2
	if result != expected {
		t.Error("analysis.TestMergeByGenre2: Expected ", expected, ", got ", result)
	}
	log.Printf("analysis.TestMergeByGenre2: End ******** \n")	
}
