/*
Library for sotring bigram frequencies in a map
*/
package ngram

// Single record of the frequency of occurence of a bigram
type BigramFreq struct {
	BigramVal Bigram
	Frequency int
}

// Map of the frequency of occurence of a bigram in a collection of texts
type BigramFreqMap struct {
	BM map[string]BigramFreq
}

// Put the bigram in the bigram frequency map
func (bfm *BigramFreqMap) GetBigram(bigram Bigram) BigramFreq {
	return bfm.BM[bigram.String()]
}

// Constructore
func NewBigramFreqMap() *BigramFreqMap {
	return &BigramFreqMap{map[string]BigramFreq{}}
}

// Put the bigram in the bigram frequency map
func (bfm *BigramFreqMap) PutBigram(bigram Bigram) {
	if !bigram.ContainsFunctionWord() {
		if bf, ok := bfm.BM[bigram.String()]; !ok {
			bfm.BM[bigram.String()] = BigramFreq{bigram, 1}
		} else {
			bf.Frequency++
			bfm.BM[bigram.String()] = bf
		}
	}
}
