package index

import (
	"testing"
)

// Empty test for corpus index writing
func TestReset(t *testing.T) {
	Reset()
}

// Empty test for corpus index writing
func TestWriteIndexCorpus0(t *testing.T) {
	sortedWords := []SortedWordItem{}
	unknownChars := []SortedWordItem{}
	WriteIndexCorpus(sortedWords, unknownChars)
}

// Trivial test for corpus index writing
func TestWriteIndexCorpus1(t *testing.T) {
	sw := SortedWordItem{"鐵", 1}
	sortedWords := []SortedWordItem{sw}
	uw := SortedWordItem{"𣣌", 2}
	unknownChars := []SortedWordItem{uw}
	WriteIndexCorpus(sortedWords, unknownChars)
}

// Empty test for document index writing
func TestWriteIndexDoc0(t *testing.T) {
	sortedWords := []SortedWordItem{}
	WriteIndexDoc(sortedWords, "test.txt")
}

// Trivial test for document index writing
func TestWriteIndexDoc1(t *testing.T) {
	sw := SortedWordItem{"鐵", 1}
	sortedWords := []SortedWordItem{sw}
	WriteIndexDoc(sortedWords, "test.txt")
}
