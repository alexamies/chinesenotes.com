package index

import (
	"fmt"
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
	WriteWFCorpus(sortedWords, unknownChars, 1)
	WriteWFDoc(sortedWords, "file.txt", 1)
	BuildIndex()
	entries := wfdoc[w]
	expected := 1
	if len(entries) != expected {
		t.Error("index.TestBuildIndex1: Expected ", expected, " got ", len(entries))
	}
	documents := FindForKeyword(w)
	retExpected := 1
	if len(documents) != retExpected {
		t.Error("index.TestReadWF1: retExpected ", retExpected, " got ", len(documents))
	}
	fmt.Println("index.TestBuildIndex1 ", documents)
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
	WriteWFCorpus(sortedWords, unknownChars, 1)
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
	WriteWFCorpus(sortedWords, unknownChars, 0)
}

// Simple test for corpus index writing
func TestWriteWFCorpus1(t *testing.T) {
	sw := SortedWordItem{"鐵", 1}
	sortedWords := []SortedWordItem{sw}
	uw := SortedWordItem{"𣣌", 2}
	unknownChars := []SortedWordItem{uw}
	WriteWFCorpus(sortedWords, unknownChars, 1)
}

// Trivial test for document index writing
func TestWriteWFDoc0(t *testing.T) {
	sortedWords := []SortedWordItem{}
	WriteWFDoc(sortedWords, "test.txt", 0)
}

// Simple test for document index writing
func TestWriteWFDoc1(t *testing.T) {
	sw := SortedWordItem{"鐵", 1}
	sortedWords := []SortedWordItem{sw}
	WriteWFDoc(sortedWords, "test.txt", 1)
}
