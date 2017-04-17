// Functions and methods to order keywords by tf-idf weight
package index

import (
	"cnreader/dictionary"
	"log"
	"sort"
)

// A keyword in a document
type Keyword struct {
	Term   string
	Weight float64 // from the tf-idf formula
}

type Keywords []Keyword


func (kws Keywords) Len() int {
	return len(kws)
}

func (kws Keywords) Swap(i, j int) {
	kws[i], kws[j] = kws[j], kws[i]
}

func (kws Keywords) Less(i, j int) bool {
	return kws[i].Weight > kws[j].Weight
}

// Gets the dictionary definition of a slice of strings
// Parameters
//   terms: The Chinese (simplified or traditional) text of the words
// Return
//   hws: an array of word senses
func GetHeadwordArray(keywords Keywords) ([]dictionary.HeadwordDef) {
	hws := []dictionary.HeadwordDef{}
	for _, kw := range keywords {
		hw, ok := dictionary.GetHeadword(kw.Term)
		if ok {
			hws = append(hws, hw)
		} else {
			log.Printf("index.GetHeadwordArray %s not found\n", kw.Term)
		}
	}
	return hws
}

// Orders the keyword with given frequency in a document by tf-idf weight
// Param:
//   vocab - word frequencies for a particular document
func SortByWeight(vocab map[string]int) Keywords {
	kws := []Keyword{}
	for t, count := range vocab {
		weight, ok := tfIdf(t, count)
		if ok {
			kws = append(kws, Keyword{Term: t, Weight: weight})
		} 
	}
	keywords := Keywords(kws)
	sort.Sort(keywords)
	return keywords
}