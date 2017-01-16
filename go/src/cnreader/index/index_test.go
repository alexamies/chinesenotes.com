package index

import (
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

// Simple test for corpus-wide word frequency reading
func TestBuildIndex1(t *testing.T) {
	Reset()
	w := "鐵"
	sw := SortedWordItem{w, 1}
	sortedWords := []SortedWordItem{sw}
	unknownChars := []SortedWordItem{}
	WriteIndexCorpus(sortedWords, unknownChars, 1)
	WriteIndexDoc(sortedWords, "file.txt", 1)
	BuildIndex()
	entries := wfdoc[w]
	expected := 1
	if len(entries) != expected {
		t.Error("index.TestReadWF1: Expected ", expected, " got ", len(entries))
	}
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
	WriteIndexCorpus(sortedWords, unknownChars, 1)
	readWFCorpus()
	entry := wf[w]
	expected := 1
	if entry.Count != expected {
		t.Error("index.TestReadWFCorpus1: Expected ", expected, " got ",
			entry.Count)
	}
}

// Trivial test for corpus index writing
func TestWriteIndexCorpus0(t *testing.T) {
	sortedWords := []SortedWordItem{}
	unknownChars := []SortedWordItem{}
	WriteIndexCorpus(sortedWords, unknownChars, 0)
}

// Simple test for corpus index writing
func TestWriteIndexCorpus1(t *testing.T) {
	sw := SortedWordItem{"鐵", 1}
	sortedWords := []SortedWordItem{sw}
	uw := SortedWordItem{"𣣌", 2}
	unknownChars := []SortedWordItem{uw}
	WriteIndexCorpus(sortedWords, unknownChars, 1)
}

// Trivial test for document index writing
func TestWriteIndexDoc0(t *testing.T) {
	sortedWords := []SortedWordItem{}
	WriteIndexDoc(sortedWords, "test.txt", 0)
}

// Simple test for document index writing
func TestWriteIndexDoc1(t *testing.T) {
	sw := SortedWordItem{"鐵", 1}
	sortedWords := []SortedWordItem{sw}
	WriteIndexDoc(sortedWords, "test.txt", 1)
}
