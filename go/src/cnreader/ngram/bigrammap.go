/*
Library for storing bigram frequencies in a map
*/
package ngram

// Single record of the frequency of occurence of a bigram
type BigramFreq struct {
	BigramVal Bigram
	Frequency int
}

// Map of the frequency of occurence of a bigram in a collection of texts
type BigramFreqMap map[string]BigramFreq

// Does the Bigram map contain a bigram with this combination of words?
func (bfmPtr *BigramFreqMap) GetBigramVal(id1, id2 int) (*Bigram, bool) {
	bfm := *bfmPtr
	key := bigramKey(id1, id2)
	bf, ok := bfm[key]
	if ok {
		return &bf.BigramVal, ok
	}
	return NullBigram(), ok
}

// Put the bigram in the bigram frequency map
func (bfmPtr *BigramFreqMap) GetBigram(bigram *Bigram) BigramFreq {
	bfm := *bfmPtr
	return bfm[bigram.String()]
}

// Merge another bigram frequency map
func (bfmPtr *BigramFreqMap) Merge(more BigramFreqMap) {
	bfm := *bfmPtr
	for k, v := range more {
    	if bf, ok := bfm[k]; ok {
    		bf.Frequency += v.Frequency
    		bfm[k] = bf
    	} else {
    		bfm[k] = v
    	}
	}
}

// Put the bigram in the bigram frequency map
func (bfmPtr *BigramFreqMap) PutBigram(bigram *Bigram) {
	if !bigram.ContainsFunctionWord() {
		bfm := *bfmPtr
		if bf, ok := bfm[bigram.String()]; !ok {
			bfm[bigram.String()] = BigramFreq{*bigram, 1}
		} else {
			bf.Frequency++
			bfm[bigram.String()] = bf
		}
	}
}

// Put the bigram in the bigram frequency map
func (bfmPtr *BigramFreqMap) PutBigramFreq(bigramFreq BigramFreq) {
	bfm := *bfmPtr
	key := bigramFreq.BigramVal.String()
	if bf, ok := bfm[key]; !ok {
		bfm[key] = bigramFreq
	} else {
		bf.Frequency += bigramFreq.Frequency
		bfm[key] = bf
	}
}