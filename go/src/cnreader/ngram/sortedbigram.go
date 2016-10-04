/*
Library for sorting the bigram frequency map
*/
package ngram

import (
	"sort"
	)

// Sorted into descending order with most frequent bigram first
type SortedBFM struct {
	bfm BigramFreqMap
	f []BigramFreq
}

func (sbf *SortedBFM) Len() int {
	return len(sbf.bfm)
}

func (sbf *SortedBFM) Less(i, j int) bool {
	return sbf.bfm.GetBigram(&sbf.f[i].BigramVal).Frequency > sbf.bfm.GetBigram(&sbf.f[j].BigramVal).Frequency
}

func NewSortedBFM(bfm BigramFreqMap) *SortedBFM {
	return &SortedBFM{bfm, []BigramFreq{}}
}

func (sbf *SortedBFM) Swap(i, j int) {
	sbf.f[i], sbf.f[j] = sbf.f[j], sbf.f[i]
}

// Get the bigram frequencies as a sorted array
func SortedFreq(bfm BigramFreqMap) []BigramFreq {
	sortedBFM := new(SortedBFM)
	sortedBFM.bfm = bfm
	sortedBFM.f = make([]BigramFreq, len(bfm))
	i := 0
	for key, _ := range bfm {
		sortedBFM.f[i] = sortedBFM.bfm[key]
		i++
	}
	sort.Sort(sortedBFM)
	return sortedBFM.f
}