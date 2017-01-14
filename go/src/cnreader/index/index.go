/*
Package for index building
*/
package index

import (
	"bufio"
	"cnreader/config"
	"fmt"
	"log"
	"os"
)

// Resets the document analysis plain text file
func Reset() {

	dir := config.IndexDir()

	// Word frequencies
	fname := dir + "/word_freq_doc.txt"
	wfFile, _ := os.Create(fname)
	wfFile.Close()
}

// Write corpus analysis to plain text files in the index directory
func WriteIndexCorpus(sortedWords, sortedUnknownWords []SortedWordItem) {
	dir := config.IndexDir()

	// Word frequencies
	wfFile, err := os.Create(dir + "/word_frequencies.txt")
	if err != nil {
		log.Printf("Could not open write wfFile", err)
		return
	}
	defer wfFile.Close()
	wfWriter := bufio.NewWriter(wfFile)
	for _, wordItem := range sortedWords {
		fmt.Fprintf(wfWriter, "%s\t%d\n", wordItem.Word, wordItem.Freq)
	}
	wfWriter.Flush()

	// Write unknown characters to a text file
	unknownCharsFile, err := os.Create(dir + "/unknown.txt")
	if err != nil {
		log.Printf("Could not open write unknownCharsFile", err)
		return
	}
	defer unknownCharsFile.Close()
	w := bufio.NewWriter(unknownCharsFile)
	for _, wordItem := range sortedUnknownWords {
		for _, r := range wordItem.Word {
			fmt.Fprintf(w, "U+%X\t%c", r, r)
		}
		fmt.Fprintln(w)
	}
	w.Flush()
}

// Append document analysis to a plain text file in the index directory
func WriteIndexDoc(sortedWords []SortedWordItem, srcFile string) {

	if srcFile == "lunyu/lunyu001.txt" {
		log.Printf("index.WriteIndexDoc: file = %s, len(sortedWords) = %d",
			srcFile, len(sortedWords))
	}

	dir := config.IndexDir()

	// Word frequencies
	fname := dir + "/word_freq_doc.txt"
	wfFile, err := os.OpenFile(fname, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Printf("Could not open write wfFile", err)
		return
	}
	defer wfFile.Close()
	wfWriter := bufio.NewWriter(wfFile)
	for _, wordItem := range sortedWords {
		fmt.Fprintf(wfWriter, "%s\t%d\t%s\n", wordItem.Word, wordItem.Freq,
			srcFile)
	}
	wfWriter.Flush()
}
