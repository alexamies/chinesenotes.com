package index

import (
	"testing"
)

// Trivial test for document index writing
func TestWriteWFDoc0(t *testing.T) {
	wfMap := WordFreqDocMap{}
	df := DocumentFrequency{}
	wfMap.WriteToFile(df)
}

// Simple test for document index writing
func TestWriteWFDoc1(t *testing.T) {
	wfRec := WordFreqDocRecord{"鐵", 1, "test.html"}
	wfMap := WordFreqDocMap{}
	wfMap.Put(wfRec)
	df := DocumentFrequency{}
	wfMap.WriteToFile(df)
}
