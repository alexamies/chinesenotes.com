/*
Library for word frequencies for each doc in the corpus
*/
package index

import (
	"bufio"
	"cnreader/config"
	"fmt"
	"log"
	"os"
)

// Word frequencies for each document
const WF_DOC_FILE = "word_freq_doc.txt"

// Remembers the word frequency for each term for each document in the corpus
type WordFreqDocRecord struct {
	Word string
	Freq int
	GlossFile string
}

// Remembers the word frequency for each term for each document in the corpus
type WordFreqDocMap map[string]WordFreqDocRecord

// Ads a map of word frequencies for a given document to the map
func (wfDocMap WordFreqDocMap) AddWF(vocab map[string]int, glossFile string) {
	//log.Printf("index.AddWF: enter %d, %d", len(wfDocMap), len(vocab))
	for word, count := range vocab {
		record := WordFreqDocRecord{word, count, glossFile}
		wfDocMap.Put(record)
	}
	//log.Printf("index.AddWF: exit %d", len(wfDocMap))
}

//Merge two WordFreqDocMap struct's together
func (wfDocMap WordFreqDocMap) Merge(wfDocMap2 WordFreqDocMap) {
	//log.Printf("index.Merge: enter %d, %d", len(wfDocMap), len(wfDocMap2))
	for _, record := range wfDocMap2 {
		wfDocMap.Put(record)
	}
}

// Adds a record to the map
func (wfDocMap WordFreqDocMap) Put(record WordFreqDocRecord) {
	key := record.Word + record.GlossFile
	_, ok := wfDocMap[key]
	if ok {
		log.Printf("index.WordFreqDocRecord: key, %s %s is already in map", 
			record.Word, record.GlossFile)
		return
	}
	wfDocMap[key] = record
}

// Append document analysis to a plain text file in the index directory
func (wordFreqDocMap WordFreqDocMap) WriteToFile(df DocumentFrequency) {
	log.Printf("index.WriteToFile: enter, %d", len(wordFreqDocMap))
	dir := config.IndexDir()
	fname := dir + "/" + WF_DOC_FILE
	wfFile, err := os.OpenFile(fname, os.O_WRONLY, 0600)
	if err != nil {
		log.Printf("Could not open write wfFile", err)
		return
	}
	defer wfFile.Close()
	wfWriter := bufio.NewWriter(wfFile)
	for _, record := range wordFreqDocMap {
		docFreq, ok := df.IDF(record.Word)
		if !ok {
			log.Printf("Could not compute document frequency for %s\n",
				record.Word)
			docFreq = 0.0
		}
		fmt.Fprintf(wfWriter, "%s,%d,%s,%.4f\n", record.Word, record.Freq,
			record.GlossFile, docFreq)
	}
	wfWriter.Flush()
}