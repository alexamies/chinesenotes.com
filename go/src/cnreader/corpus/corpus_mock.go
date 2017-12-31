/*
 * Mock objects for the corpus package
 */
package corpus

import (
	"time"
)


// Implements the CorpusLoader interface with trivial implementation
type EmptyCorpusLoader struct {Label string}

func (loader EmptyCorpusLoader) LoadCollection(fName string) (CollectionEntry, error) {
	return CollectionEntry{}, nil
}

func (loader EmptyCorpusLoader) LoadCorpus(fName string) []CollectionEntry {
	return []CollectionEntry{}
}

// Implements the CorpusLoader interface with mock data
type MockCorpusLoader struct {Label string}

func (loader MockCorpusLoader) LoadCollection(fName string) (CollectionEntry, error) {
	entry := CorpusEntry{
		RawFile: "corpus_doc.txt",
		GlossFile: "corpus_doc.html",
		Title: "corpus doc title",
		ColTitle: "A Corpus Collection",
	}
	dateUpdated := time.Now().Format("2006-01-02")
	c := CollectionEntry{
		CollectionFile: "a_collection_file.txt",
		GlossFile: "a_collection_file.html",
		Title: "A Corpus Collection",
		Summary: "A summary",
		Intro: "An introduction",
		DateUpdated: dateUpdated,
		Corpus: "A Corpus",
		CorpusEntries: []CorpusEntry{entry},
		AnalysisFile: "collection_analysis_file.html",
		Format: "prose",
		Date: "1984",
		Genre: "Science Fiction",
	}
	return c, nil
}

func (loader MockCorpusLoader) LoadCorpus(fName string) []CollectionEntry {
	c, _ := loader.LoadCollection(fName)
	return []CollectionEntry{c}
}
