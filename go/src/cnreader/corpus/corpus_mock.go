/*
 * Mock objects for the corpus package
 */
package corpus

import (
	"time"
)


// Implements the CorpusLoader interface with trivial implementation
type EmptyCorpusLoader struct {Label string}

func (loader EmptyCorpusLoader) GetCollectionEntry(fName string) (CollectionEntry, error) {
	return CollectionEntry{}, nil
}

func (loader EmptyCorpusLoader) LoadAll(fName string) (map[string]CorpusEntry) {
	return map[string]CorpusEntry{}
}

func (loader EmptyCorpusLoader) LoadCollection(fName, colTitle string) []CorpusEntry {
	return []CorpusEntry{}
}

func (loader EmptyCorpusLoader) LoadCorpus(fName string) []CollectionEntry {
	return []CollectionEntry{}
}

func (loader EmptyCorpusLoader) ReadText(fName string) string {
	return ""
}

// Implements the CorpusLoader interface with mock data
type MockCorpusLoader struct {Label string}

func (loader MockCorpusLoader) GetCollectionEntry(fName string) (CollectionEntry, error) {
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

func (loader MockCorpusLoader) LoadAll(fName string) (map[string]CorpusEntry) {
	corpusEntryMap := map[string]CorpusEntry{}
	collections := loader.LoadCorpus(fName)
	for _, collectionEntry := range collections {
		corpusEntries := loader.LoadCollection(collectionEntry.CollectionFile,
				collectionEntry.Title)
		for _, entry := range corpusEntries {
			corpusEntryMap[entry.RawFile] = entry
		}
	}
	return corpusEntryMap
}

func (loader MockCorpusLoader) LoadCollection(fName, colTitle string) []CorpusEntry {
	entry := CorpusEntry{
		RawFile: "raw_file.txt",
		GlossFile: "gloss_file.html",
		Title: "Entry Title",
		ColTitle: "Collection Title",
	}
	return []CorpusEntry{entry}
}

func (loader MockCorpusLoader) LoadCorpus(fName string) []CollectionEntry {
	c, _ := loader.GetCollectionEntry(fName)
	return []CollectionEntry{c}
}

func (loader MockCorpusLoader) ReadText(fName string) string {
	return "你好 Hello!"
}