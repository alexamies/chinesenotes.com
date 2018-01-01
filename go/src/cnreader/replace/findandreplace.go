/*
Library for finding and replacing strings in the library.
*/
package replace

import (
	"cnreader/library"
	"log"
	"strings"
)

// Finds occurrences of the expression in the library
// Parameters
//   substr - the string to find
//   lib - the library to search over
func Find(substr string, lib library.Library) {
	log.Printf("replace.Find substr: %s\n", substr)
	corpora := lib.Loader.LoadLibrary()
	corpLoader := lib.Loader.GetCorpusLoader()
	for i, corpus := range corpora {
		log.Printf("replace.Find %d: corpus: %v\n", i, corpus)
		collections := corpLoader.LoadCorpus(corpus.FileName)
		for j, col := range collections {
			log.Printf("replace.Find j: %d: col: %v\n", j, col)
			entries := corpLoader.LoadCollection(col.CollectionFile, col.Title)
			for _, entry := range entries {
				text := corpLoader.ReadText(entry.RawFile)
				res := strings.Contains(text, substr)
				log.Printf("replace.Find title: %s: found: %v\n", col.Title, res)
			}
		}
	}
}