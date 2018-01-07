/*
Library for documents retrieval
*/
package index

import (
	"bufio"
	"cnreader/corpus"
	"cnreader/dictionary"
	"cnreader/config"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

// Maximum number of documents displayed on a web page
const MAX_DOCS_DISPLAYED = 10

// A document-specific word frequency entry record
type RetrievalResult struct {
	HTMLFile, Title, ColTitle string
	Count int
}

// Retrieves documents with title for a single keyword
func FindDocsForKeyword(keyword dictionary.HeadwordDef,
		corpusEntryMap map[string]corpus.CorpusEntry) []RetrievalResult {
	docs := make([]RetrievalResult, 0)
	if !keywordIndexReady {
		log.Printf("index.FindForKeyword, Warning: index not yet ready")
		entry := RetrievalResult{"Index not ready", "", "", 0}
		return []RetrievalResult{entry}
	}
	// TODO - separate corpora into simplified and traditional. At the moment
	// only traditional will work
	kw := *keyword.Simplified
	if *keyword.Traditional != "\\N" {
		kw = *keyword.Traditional
	}
	i := 0
	for _, raw := range wfdoc[kw] {
		if i < MAX_DOCS_DISPLAYED {
			corpusEntry, ok := corpusEntryMap[raw.Filename]
			if !ok {
				//log.Printf("index.FindForKeyword, no title for %s\n",
				//		raw.Filename)
				continue
			}
			item := RetrievalResult{corpusEntry.GlossFile, corpusEntry.Title,
				corpusEntry.ColTitle, raw.Count}
			docs = append(docs, item)
			i++
		} else {
			break
		}
	}
	return docs
}

// Retrieves raw results for a single keyword
func FindForKeyword(keyword string) []WFDocEntry {
	if !keywordIndexReady {
		log.Printf("index.FindForKeyword, Warning: index not yet ready")
		entry := WFDocEntry{"Index not ready", 0, 0.0}
		return []WFDocEntry{entry}
	}
	return wfdoc[keyword]
}

// Loads the keyword index from disk
func LoadKeywordIndex() {
	dir := config.IndexDir()

	fname := dir + "/" + KEYWORD_INDEX_FILE
	f, err := os.Open(fname)
	if err != nil {
		log.Printf("index.LoadKeywordIndex, error opening index file: ", err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	dec := json.NewDecoder(r)
	for {
		if err := dec.Decode(&wfdoc); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}
	keywordIndexReady = true
	fmt.Printf("index.LoadKeywordIndex, index loaded: %d entries", len(wfdoc))
}
