// Test sorting of word frequencies
package analysis

import (
	"log"
	"testing"
)

// Test merging of word frequencies
func TestMerge(t *testing.T) {
	log.Printf("analysis.TestMerge: Begin ******** \n")	
	classicalWF1 := NewWordFreqByGenre("classical")
	classicalWF1.WF["thou"] = 1
	classicalWF2 := NewWordFreqByGenre("classical")
	classicalWF2.WF["thou"] = 2
	genreWF1 := WFArrayByGenre{classicalWF1}
	genreWF1.Merge(classicalWF2)
	wf := genreWF1.Get("classical")
	result := wf["thou"]
	expected := 3
	if result != result {
		t.Error("analysis.TestMerge: Expected ", expected, ", got ", result)
	}
	log.Printf("analysis.TestMerge: End ******** \n")	
}
