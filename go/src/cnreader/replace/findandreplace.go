/*
Library for finding and replacing strings in the library.
*/
package replace

import (
	"cnreader/library"
	"log"
)

// Finds occurrences of the expression in the library
// Parameters
//   expr - the string to find
//   lib - the library to search over
func Find(expr string, lib library.Library) {
	log.Printf("Find expr: %s\n", expr)
	corpora := lib.Loader.LoadLibrary()
	for i, corpus := range corpora {
		log.Printf("Find %d: corpus: %v\n", i, corpus)
		collections := lib.Loader.GetCorpusLoader().LoadCorpus(corpus.FileName)
		for j, col := range collections {
			log.Printf("Find j: %d: col: %v\n", j, col)
			for k, entry := range col.CorpusEntries {
				log.Printf("Find k: %d: entry: %v\n", k, entry)
			}
		}
	}
}