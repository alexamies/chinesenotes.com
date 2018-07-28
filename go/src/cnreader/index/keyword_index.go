/*
Package for keyword index building and use
*/
package index

import (
	"bufio"
	"cnreader/config"
	"cnreader/ngram"
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

// ngram frequencies for corpus
const NGRAM_CORPUS_FILE = "ngram_frequencies.txt"

// A word frequency entry record
type WFEntry struct {
	Chinese string
	Count   int
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
	log.Printf("index.readWFCorpus: reading %s\n", fname)
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
		wfentry := WFEntry{
			Chinese: w,
			Count:   int(count),
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
		filename := row[2]
		entry := WFDocEntry{filename, int(count)}
		if entryarr, ok := wfdoc[w]; !ok {
			wfslice := make([]WFDocEntry, 1)
			wfslice[0] = entry
			wfdoc[w] = wfslice
		} else {
			wfdoc[w] = append(entryarr, entry)
		}
	}
	log.Printf("index.ReadWFDoc: read %d records\n", len(wfdoc))
}

// Resets the document analysis plain text file
func Reset() {
	dir := config.IndexDir()
	fname := dir + "/" + WF_DOC_FILE
	wfFile, err := os.Create(fname)
	if err != nil {
		log.Fatal("index.Reset: Could not reset file", err)
	}
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
func WriteWFCorpus(sortedWords, sortedUnknownWords []SortedWordItem,
		bFreq []ngram.BigramFreq, wc int) {

	log.Printf("index.WriteWFCorpus: enter")

	// Word frequencies
	dir := config.IndexDir()
	fname := dir + "/" + WF_CORPUS_FILE
	wfFile, err := os.Create(fname)
	if err != nil {
		log.Fatal("Could not open write wfFile", err)
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

	// Write ngrams to a file
	ngramFile, err := os.Create(dir + "/" + NGRAM_CORPUS_FILE)
	if err != nil {
		log.Printf("Could not open write ngramFile", err)
		return
	}
	defer ngramFile.Close()
	nWriter := bufio.NewWriter(ngramFile)
	for _, ngramItem := range bFreq {
		rel_freq := 0.0
		if wc > 0 {
			rel_freq = float64(ngramItem.Frequency) * 10000.0 / float64(wc)
		}
		fmt.Fprintf(nWriter, "%s\t%d\t%f\n", ngramItem.BigramVal.Traditional(),
			ngramItem.Frequency,	rel_freq)
	}
	nWriter.Flush()
}