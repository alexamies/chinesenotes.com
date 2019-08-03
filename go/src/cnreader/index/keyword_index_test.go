package index

import (
	"github.com/alexamies/cnreader/ngram"
	"testing"
)

// Empty test for corpus index writing
func TestReset(t *testing.T) {
	Reset()
}

// Trivial test for corpus-wide word frequency reading
func TestBuildIndex0(t *testing.T) {
	BuildIndex()
}

// Trivial test for corpus-wide word frequency reading
func TestReadWFCorpus0(t *testing.T) {
	readWFCorpus()
}

// Trivial test for corpus-wide word frequency reading
func TestReadWFCorpus1(t *testing.T) {
	w := "鐵"
	sw := SortedWordItem{w, 1}
	sortedWords := []SortedWordItem{sw}
	unknownChars := []SortedWordItem{}
	bFreq := []ngram.BigramFreq{}
	WriteWFCorpus(sortedWords, unknownChars, bFreq, 1)
	readWFCorpus()
	entry := wf[w]
	expected := 1
	if entry.Count != expected {
		t.Error("index.TestReadWFCorpus1: Expected ", expected, " got ",
			entry.Count)
	}
}

// Trivial test for corpus index writing
func TestWriteWFCorpus0(t *testing.T) {
	sortedWords := []SortedWordItem{}
	unknownChars := []SortedWordItem{}
	bFreq := []ngram.BigramFreq{}
	WriteWFCorpus(sortedWords, unknownChars, bFreq, 0)
}

// Simple test for corpus index writing
func TestWriteWFCorpus1(t *testing.T) {
	sw := SortedWordItem{"鐵", 1}
	sortedWords := []SortedWordItem{sw}
	uw := SortedWordItem{"𣣌", 2}
	unknownChars := []SortedWordItem{uw}
	bFreq := []ngram.BigramFreq{}
	WriteWFCorpus(sortedWords, unknownChars, bFreq, 1)
}
