/*
Library for Chinese vocabulary analysis based on literary genre
*/
package analysis

// Sorted list of word frequencies
type SortedByGenre struct {
	Genre string
	SWI   []WFResult
}

// Word frequency by genre
type WordFreqByGenre struct {
	Genre string
	WF    map[string]int
}

// An array of word frequency maps by genre
type WFArrayByGenre []WordFreqByGenre

// Gets the matching word frequency map
func (wfArray WFArrayByGenre) Get(genre string) map[string]int {
	for _, wf := range wfArray {
		if wf.Genre == genre {
			return wf.WF
		}
	}
	return map[string]int{}
}

// Merge the argument into the word frequency map for the matching genre
// more: a word frequency map for a given genre
func MergeByGenre(wfArray WFArrayByGenre, more WordFreqByGenre) WFArrayByGenre {
	found := false
	for _, wf := range wfArray {
		if wf.Genre == more.Genre {
			for k, v := range more.WF {
				wf.WF[k] += v
			}
			found = true
		}
	}
	if !found {
		wfArray = append(wfArray, more)
	}
	return wfArray
}

// Constructor
func NewWordFreqByGenre(genre string) WordFreqByGenre {
	return WordFreqByGenre{
		Genre: genre,
		WF:    map[string]int{},
	}
}
