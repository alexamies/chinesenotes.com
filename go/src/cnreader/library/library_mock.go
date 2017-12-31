/*
 * Mock objects for library package
 */
package library

import (
	"cnreader/corpus"
)

// Mock loader that has no data
type EmptyLibraryLoader struct {Label string}

// Implements the method from the LibraryLoader interface
func (loader EmptyLibraryLoader) GetCorpusLoader() corpus.CorpusLoader {
	return corpus.EmptyCorpusLoader{loader.Label}
}

func (loader EmptyLibraryLoader) LoadLibrary() []CorpusData {
	return []CorpusData{}
}

// Mock loader that generates static data
type MockLibraryLoader struct {Label string}

// Implements the method from the LibraryLoader interface
func (loader MockLibraryLoader) GetCorpusLoader() corpus.CorpusLoader {
	return corpus.MockCorpusLoader{loader.Label}
}

func (loader MockLibraryLoader) LoadLibrary() []CorpusData {
	c := CorpusData{
		Title: "Title",
		ShortName: "ShortName",
		Status: "Status",
		FileName: "FileName",
	}
	return []CorpusData{c}
}
