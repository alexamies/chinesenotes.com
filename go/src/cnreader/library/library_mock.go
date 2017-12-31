/*
 * Mock objects for library package
 */
package library

import (
	"cnreader/corpus"
	"time"
)

// Mock loader that has no data
type EmptyLibraryLoader struct {Label string}

func (loader EmptyLibraryLoader) LoadLibrary() []CorpusData {
	return []CorpusData{}
}

func (loader EmptyLibraryLoader) LoadCorpus(fName string) []corpus.CollectionEntry {
	return []corpus.CollectionEntry{}
}

// Mock loader that generates static data
type MockLibraryLoader struct {Label string}

func (loader MockLibraryLoader) LoadLibrary() []CorpusData {
	c := CorpusData{
		Title: "Title",
		ShortName: "ShortName",
		Status: "Status",
		FileName: "FileName",
	}
	return []CorpusData{c}
}

func (loader MockLibraryLoader) LoadCorpus(fName string) []corpus.CollectionEntry {
	dateUpdated := time.Now().Format("2006-01-02")
	c := corpus.CollectionEntry{
		CollectionFile: "a_collection_file.txt",
		GlossFile: "a_collection_file.html",
		Title: "A Corpus Collection",
		Summary: "A summary",
		Intro: "An introduction",
		DateUpdated: dateUpdated,
		Corpus: "A Corpus",
		CorpusEntries: []corpus.CorpusEntry{},
		AnalysisFile: "collection_analysis_file.html",
		Format: "prose",
		Date: "1984",
		Genre: "Science Fiction",
	}
	return []corpus.CollectionEntry{c}
}