/*
Package for keyword index building and use
*/
package index

import (
	"bufio"
	"cnreader/config"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// File name for keyword index
const KEYWORD_INDEX_FILE = "keyword_index.json"

// Maximum number of keywords in index
const MAX_FILES_PER_KEYWORD = 50

// Unknown characters file
const UNKNOWN_FILE = "unknown.txt"

// Word frequencies for corpus
const WF_CORPUS_FILE = "word_frequencies.txt"

// Word frequencies for each document
const WF_DOC_FILE = "word_freq_doc.txt"

// A word frequency entry record
type WFEntry struct {
	Chinese string
	Count   int
	Freq    float64 // frequency per 10,000 words
}

// Keep corpus-wide word frequency map in memory
var wf map[string]WFEntry

// Keep document-specific word frequency map in memory
var wfdoc map[string][]WFDocEntry

// For checking the status of the keyword index
var keywordIndexReady bool

// Reads word frequencies data from files into memory and builds the keyword
// index
func BuildIndex() {
	readWFCorpus()
	readWFDoc()
	writeKeywordIndex()
	keywordIndexReady = true
}

// Reads corpus-wide word frequencies from file into memory
func readWFCorpus() {
	wf = make(map[string]WFEntry)
	dir := config.IndexDir()
	fname := dir + "/" + WF_CORPUS_FILE
	log.Printf("index.ReadWFCorpus: reading %s\n", fname)
	wffile, err := os.Open(fname)
	if err != nil {
		log.Fatal("index.ReadWFCorpus, error opening word freq file: ", err)
	}
	defer wffile.Close()
	reader := csv.NewReader(wffile)
	reader.FieldsPerRecord = -1
	reader.Comma = rune('\t')
	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		log.Fatal("index.ReadWFCorpus: Could not wf file ", err)
	}
	for i, row := range rawCSVdata {
		w := row[0] // Chinese text for word
		count, err := strconv.ParseInt(row[1], 10, 0)
		if err != nil {
			log.Fatal("Could not parse word count ", i, err)
		}
		freq, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			log.Fatal("Could not parse word freq ", i, err)
		}
		wfentry := WFEntry{
			Chinese: w,
			Count:   int(count),
			Freq:    freq,
		}
		wf[w] = wfentry
	}
}

// Reads document-specific word frequencies from file into memory
func readWFDoc() {
	wfdoc = make(map[string][]WFDocEntry)
	dir := config.IndexDir()
	fname := dir + "/" + WF_DOC_FILE
	log.Printf("index.ReadWFDoc: reading %s\n", fname)
	wffile, err := os.Open(fname)
	if err != nil {
		log.Fatal("index.ReadWFDoc, error opening word freq file: ", err)
	}
	defer wffile.Close()
	reader := csv.NewReader(wffile)
	reader.FieldsPerRecord = -1
	reader.Comma = rune('\t')
	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		log.Fatal("index.ReadWFDoc: Could not wf file ", err)
	}
	for i, row := range rawCSVdata {
		w := row[0] // Chinese text for word
		count, err := strconv.ParseInt(row[1], 10, 0)
		if err != nil {
			log.Fatal("Could not parse word count ", i, err)
		}
		freq, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			log.Fatal("Could not parse word freq ", i, err)
		}
		filename := row[3]
		entry := WFDocEntry{filename, int(count), freq}
		if entryarr, ok := wfdoc[w]; !ok {
			wfslice := make([]WFDocEntry, 1)
			wfslice[0] = entry
			wfdoc[w] = wfslice
		} else {
			wfdoc[w] = append(entryarr, entry)
		}
	}
}

// Resets the document analysis plain text file
func Reset() {
	dir := config.IndexDir()
	fname := dir + "/word_freq_doc.txt"
	wfFile, _ := os.Create(fname)
	wfFile.Close()
	keywordIndexReady = false
}

// Writes a JSON format keyword index to look up top documents for each keyword
func writeKeywordIndex() {
	dir := config.IndexDir()

	// Word frequencies
	fname := dir + "/" + KEYWORD_INDEX_FILE
	f, err := os.Create(fname)
	if err != nil {
		log.Printf("index.writeKeywordIndex: Could not create file", err)
		return
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	for k, items := range wfdoc {
		sort.Sort(ByFrequencyDoc(items))
		if len(items) > MAX_FILES_PER_KEYWORD {
			wfdoc[k] = items[:MAX_FILES_PER_KEYWORD]
		}
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(wfdoc)
	w.Flush()
}

// Write corpus analysis to plain text files in the index directory
func WriteWFCorpus(sortedWords, sortedUnknownWords []SortedWordItem, wc int) {

	// Word frequencies
	dir := config.IndexDir()
	fname := dir + "/" + WF_CORPUS_FILE
	wfFile, err := os.Create(fname)
	if err != nil {
		log.Printf("Could not open write wfFile", err)
		return
	}
	defer wfFile.Close()
	wfWriter := bufio.NewWriter(wfFile)
	for _, wordItem := range sortedWords {
		rel_freq := 0.0
		if wc > 0 {
			rel_freq = float64(wordItem.Freq) * 10000.0 / float64(wc)
		}
		fmt.Fprintf(wfWriter, "%s\t%d\t%f\n", wordItem.Word, wordItem.Freq,
			rel_freq)
	}
	wfWriter.Flush()

	// Write unknown characters to a text file
	unknownCharsFile, err := os.Create(dir + "/" + UNKNOWN_FILE)
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
func WriteWFDoc(sortedWords []SortedWordItem, srcFile string, wc int) {

	dir := config.IndexDir()

	// Word frequencies
	fname := dir + "/" + WF_DOC_FILE
	wfFile, err := os.OpenFile(fname, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Printf("Could not open write wfFile", err)
		return
	}
	defer wfFile.Close()
	wfWriter := bufio.NewWriter(wfFile)
	for _, wordItem := range sortedWords {
		rel_freq := 0.0
		if wc > 0 {
			rel_freq = float64(wordItem.Freq) * 10000.0 / float64(wc)
		}
		fmt.Fprintf(wfWriter, "%s\t%d\t%f\t%s\n", wordItem.Word, wordItem.Freq,
			rel_freq, srcFile)
	}
	wfWriter.Flush()
}