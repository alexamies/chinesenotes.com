// Test document frequency
package index

import (
	"log"
	"testing"
)

// Trivial test for document frequency
func TestAddVocabulary0(t *testing.T) {
	log.Println("index.DocumentFrequency.TestAddVocabulary enter")
	df := NewDocumentFrequency()
	vocab := map[string]int{}
	df.AddVocabulary(vocab)
	lenExpected := 0
	lenReturned := len(df.DocFreq)
	if lenReturned != lenExpected {
		t.Error("index.TestAddVocabulary0: lenExpected ", lenExpected, " got ",
			lenReturned)
	}
}

// Trivial test for document frequency
func TestIDF0(t *testing.T) {
	df := NewDocumentFrequency()
	vocab := map[string]int{}
	df.AddVocabulary(vocab)
	_, ok := df.IDF("")
	okExpected := false
	if ok != okExpected {
		t.Error("index.TestIDF0: okExpected ", okExpected, " got ", ok)
	}
	n := df.N
	nExpected := 1
	if n != nExpected {
		t.Error("index.TestIDF0: nExpected ", nExpected, " got ", n)
	}
}

// Simple test for document frequency
func TestIDF1(t *testing.T) {
	df := NewDocumentFrequency()
	term := "car"
	vocab := map[string]int{
		term: 1,
	}
	df.AddVocabulary(vocab)
	val, ok := df.IDF(term)
	okExpected := true
	if ok != okExpected {
		t.Error("index.TestIDF1: okExpected ", okExpected, " got ",
			ok)
	}
	valExpected := 0.0
	if val != valExpected {
		t.Error("index.TestIDF1: valExpected ", valExpected, " got ",
			val)
	}
}

// Slightly harder test
func TestIDF2(t *testing.T) {
	df := NewDocumentFrequency()
	terms := []string{"car", "auto", "insurance", "best"}
	vocab1 := map[string]int{
		terms[0]: 27,
		terms[1]: 3,
		terms[3]: 14,
	}
	df.AddVocabulary(vocab1)
	vocab2 := map[string]int{
		terms[0]: 4,
		terms[1]: 33,
		terms[2]: 33,
	}
	df.AddVocabulary(vocab2)
	vocab3 := map[string]int{
		terms[0]: 24,
		terms[2]: 29,
		terms[3]: 17,
	}
	df.AddVocabulary(vocab3)
	val0, ok0 := df.IDF(terms[0])
	okExpected0 := true
	if ok0 != okExpected0 {
		t.Error("index.TestIDF2: okExpected0 ", okExpected0, " got ",
			ok0)
	}
	valExpected0 := 0.0
	if val0 != valExpected0 {
		t.Error("index.TestIDF2: valExpected0 ", valExpected0, " got ",
			val0, ", df: ", df)
	}
}

// Example from p. 109 of Manning, Christopher D., Prabhakar Raghavan, and
// Hinrich Schütze. Introduction to information retrieval, Cambridge: Cambridge 
// university press, 2008.
func TestIDF3(t *testing.T) {
	df := NewDocumentFrequency()
	terms := []string{"car", "auto", "insurance", "best"}
	df.DocFreq[terms[0]] = 18165
	df.DocFreq[terms[1]] = 6723
	df.DocFreq[terms[2]] = 19241
	df.DocFreq[terms[3]] = 25235
	df.N = 806791
	v0, ok := df.IDF(terms[0])
	okExpected := true
	if ok != okExpected {
		t.Error("index.TestIDF0: okExpected ", okExpected, " got ", ok)
	}
	v1, ok := df.IDF(terms[1])
	v2, ok := df.IDF(terms[2])
	v3, ok := df.IDF(terms[3])
	log.Printf("index.TestIDF3 idf = (%v, %v, %v, %v)\n", v0, v1, v2, v3)
}

// Simple test for reading the document frequency data
func TestReadDocumentFrequency(t *testing.T) {
	df, err := ReadDocumentFrequency()
	if err != nil {
		t.Error("index.TestReadDocumentFrequency: error ", err)
	}
	log.Printf("index.TestReadDocumentFrequency df = %v\n", df)
}

// Trivial test for tf-idf
func TestTfIdf(t *testing.T) {
	df := NewDocumentFrequency()
	term := "car"
	vocab := map[string]int{
		term: 1,
	}
	df.AddVocabulary(vocab)
	completeDF = df
	_, ok := tfIdf("car", 1)
	if !ok {
		t.Error("index.TestTfIdf: not ok")
	}
}

// Simple test for saving document frequency data
func TestWriteToFile(t *testing.T) {
	df := NewDocumentFrequency()
	term := "car"
	vocab := map[string]int{
		term: 1,
	}
	df.AddVocabulary(vocab)
	df.WriteToFile()
}
