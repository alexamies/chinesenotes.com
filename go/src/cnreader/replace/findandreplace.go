/*
Library for finding and replacing strings in the library.
*/
package replace

import (
	"cnreader/config"
	"cnreader/corpus"
	"cnreader/library"
	"fmt"
	"log"
	"strings"
)

// Input into the find and replace feature
type Expression struct {
	Find, Replacement string
	Replace bool
}

// Encapsulates results from find and replace
type Result struct {
	Find string
	Replacement string
	Replace bool
	Occurences int
	Corpus string
	Collection string
	Document string
	
}

func (r Result) String() string {
	if !r.Replace {
		return fmt.Sprintf("Found %d of %s in %s, %s, %s", r.Occurences, r.Find,
			r.Corpus, r.Collection, r.Document)
	}
	return fmt.Sprintf("Replaced %d of %s with %s in %s, %s, %s", r.Occurences,
		r.Find, r.Replacement, r.Corpus, r.Collection, r.Document)
}

// Finds occurrences of the expression in the library
// Parameters
//   substr - the string to find
//   lib - the library to search over
func FindAndReplace(expr Expression, lib library.Library) []Result {
	log.Printf("replace.Find substr: %s\n", expr.Find)
	results := []Result{}
	corpora := lib.Loader.LoadLibrary()
	corpLoader := lib.Loader.GetCorpusLoader()
	for _, corpus := range corpora {
		//log.Printf("replace.Find %d: corpus: %v\n", i, corpus)
		collections := corpLoader.LoadCorpus(corpus.FileName)
		for _, col := range collections {
			//log.Printf("replace.Find j: %d: col: %v\n", j, col)
			documents := corpLoader.LoadCollection(col.CollectionFile, col.Title)
			for _, doc := range documents {
				src := config.CorpusDir() + "/" + doc.RawFile
				text := corpLoader.ReadText(src)
				res := strings.Contains(text, expr.Find)
				if res {
					result := Result{
							Find: expr.Find, 
							Replacement: expr.Replacement, 
							Replace: false,
							Occurences: 1, 
							Corpus: corpus.Title, 
							Collection: col.Title,  
							Document: doc.Title,
						}
					if expr.Replace {
						WriteReplacement(corpus, col, doc, text, expr.Find,
								expr.Replacement)
						result = Result{
							Find: expr.Find, 
							Replacement: expr.Replacement, 
							Replace: true,
							Occurences: 1, 
							Corpus: corpus.Title, 
							Collection: col.Title,  
							Document: doc.Title,
						}
					}
					results = append(results, result)
				}
			}
		}
	}
	return results
}

// Replace the expression in the given text with the replacement, writing to disk
func WriteReplacement(corpus library.CorpusData, col corpus.CollectionEntry,
		doc corpus.CorpusEntry, text string, find string, replacement string) {
	log.Printf("replace.WriteReplacement repacing: %s: with: %s\n", find,
		replacement)
}
