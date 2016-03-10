// Test sorting of word frequencies
package analysis

import (
	"cnreader/ngram"
	"fmt"
	"testing"
)

// Test sorting of word frequencies
func TestAddResults(t *testing.T) {
	fmt.Printf("TestAddResults: Begin unit tests\n")
	vocab := map[string]int{"one":1, "three":3, "two":2}
	usage := map[string]string {"one": "one banana"}
	unknown := map[string]int{"x":1}
	bigramFreq := ngram.NewBigramFreqMap()
	results := CollectionAResults{
		Vocab: vocab,
		Usage: usage,
		BigramFrequencies: *bigramFreq,
		WC: 3,
		UnknownChars: unknown,
	}
	moreVocab := map[string]int{"one":1, "three":1, "four":4}
	moreUsage := map[string]string {"two": "two banana"}
	unknown1 := map[string]int{"x":1}
	more := CollectionAResults{
		Vocab: moreVocab,
		Usage: moreUsage,
		BigramFrequencies: *bigramFreq,
		WC: 4,
		UnknownChars: unknown1,
	}
	results.AddResults(more)
	r := results.Vocab["three"]
	e := 4
	if r != e {
		t.Error("TestAddResults, three expected ", e, " got, ", r)
	}
	r = results.Vocab["four"]
	e = 4
	if r != e {
		t.Error("TestAddResults, four expected ", e, " got, ", r)
	}
	r = results.WC
	e = 7
	if r != e {
		t.Error("TestAddResults, word count expected ", e, " got, ", r)
	}
}
