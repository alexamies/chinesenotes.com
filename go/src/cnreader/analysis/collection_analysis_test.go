// Test sorting of word frequencies
package analysis

import (
	"fmt"
	"testing"
)

// Test sorting of word frequencies
func TestAddResults(t *testing.T) {
	fmt.Printf("TestAddResults: Begin unit tests\n")
	vocab := map[string]int{"one":1, "three":3, "two":2}
	usage := map[string]string {"one": "one banana"}
	unknown := map[string]int{"x":1}
	results := CollectionAResults{vocab, usage, 3, unknown}
	moreVocab := map[string]int{"one":1, "three":1, "four":4}
	moreUsage := map[string]string {"two": "two banana"}
	unknown1 := map[string]int{"x":1}
	results.AddResults(moreVocab, moreUsage, 4, unknown1)
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
