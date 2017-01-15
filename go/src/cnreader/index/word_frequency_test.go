// Test sorting of word frequencies
package index

import (
	"fmt"
	"testing"
)

// Test sorting of word frequencies
func TestSortedFreq1(t *testing.T) {
	fmt.Printf("TestSortedFreq: Begin unit tests\n")
	wordFreq := map[string]int{"one": 1, "three": 3, "two": 2}
	sortedWords := SortedFreq(wordFreq)
	if sortedWords == nil {
		t.Error("Expected non-nil sortedWords")
	}
	if sortedWords[0].Word != "three" {
		t.Error("Expected that 'three' to be the most frequent word")
	}
	/*
		for _, w := range sortedWords {
			fmt.Printf("TestSortedFreq: %v : %v\n", w, wordFreq[w])
		}
	*/
}
