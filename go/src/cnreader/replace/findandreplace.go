/*
Library for finding and replacing strings in the library.
*/
package replace

import (
	"cnreader/config"
	"cnreader/corpus"
	"cnreader/library"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Input into the find and replace feature
type Expression struct {
	Find, Replacement string
	Replace bool
}

// Object is needed for JSON umnmarshalling
type ExpObject struct {
	Expressions []Expression
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
func FindAndReplace(expressions []Expression, lib library.Library) []Result {
	log.Printf("replace.Find enter\n")
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
				for _, expr := range expressions {
					log.Printf("replace.Find looking for %s in %s\n",
					           expr.Find, doc.Title)
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
							text = WriteReplacement(corpus, col, doc, text, expr.Find,
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
	}
	return results
}

// Builds an expression object from TSV file contents
// Param
//   fName: The full path of the file
func ReadExp(fName string) ([]Expression, error) {
	expr := []Expression{}
	expFile, err := os.Open(fName)
	if err != nil {
		log.Print("replace.ReadExp error opening file: ", err)
		return expr, err
	}
	defer expFile.Close()
	reader := csv.NewReader(expFile)
	reader.FieldsPerRecord = 3
	reader.Comma = rune('\t')
	reader.Comment = rune('#')
	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		log.Print("replace.ReadExp error reading file: ", err)
		return expr, err
	}
	for _, row := range rawCSVdata {
		b, err := strconv.ParseBool(row[2])
		if err != nil {
			log.Print("replace.ReadExp error parsing boolean value: ", err)
			return expr, err
		}
		expr = append(expr, Expression{row[0], row[1], b})
	}
	return expr, nil
}

// Replace the expression in the given text with the replacement, writing to disk
func WriteReplacement(corpus library.CorpusData, col corpus.CollectionEntry,
		doc corpus.CorpusEntry, text string, find string,
		replacement string) string {
	log.Printf("replace.WriteReplacement repacing: %s: with: %s\n", find,
		replacement)
	r := strings.NewReplacer(find, replacement)
	rText := r.Replace(text)
	fName := config.CorpusDir() + "/replace/" + doc.RawFile
	f, err := os.Create(fName)
	if err != nil {
		log.Print("replace.WriteReplacement error opening file\n", err)
	}
	defer f.Close()
	_, err = f.WriteString(rText)
	if err != nil {
		log.Print("replace.WriteReplacement error writing to file\n", err)
	}
	return rText
}

// Replace the expression in the given text with the replacement, writing to disk
func WriteReport(results []Result) {
	fName := config.CorpusDir() + "/replace/report.txt"
	log.Printf("replace.WriteReport writing results to %s\n", fName)
	f, err := os.Create(fName)
	if err != nil {
		log.Print("replace.WriteReport error opening report file\n", err)
	}
	defer f.Close()
	str := fmt.Sprintf("%d matches found\n", len(results)) 
	_, err = f.WriteString(str)
	if err != nil {
		log.Print("replace.WriteReport error writing to report file\n", err)
	}
	for _, result := range results {
		str2 := fmt.Sprintf("%v\n", result) 
		_, err = f.WriteString(str2)
		if err != nil {
			log.Print("replace.WriteReport error writing to report file\n", err)
		}
	}
}

// Builds an expression object from a JSON string
func UnmarshalExp(expStr string) (ExpObject, error) {
	var exp ExpObject
	b := []byte(expStr)
	err := json.Unmarshal(b, &exp)
	return exp, err
}