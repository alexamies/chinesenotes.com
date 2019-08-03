/*
Excluded words, such as corpus footer copyright boilerplate
*/
package corpus

import (
	"encoding/csv"
	"github.com/alexamies/cnreader/config"
	"log"
	"os"
)

func init() {
	loadExcluded()
}

// Map of parts of speech that distinguish words as function words
var excluded map[string]bool

// Tests whether the string should be excluded from corpus analysis
// Parameter
// chunk: the string to be tested
func IsExcluded(text string) bool  {
	_, ok := excluded[text]
	return ok
}

func loadExcluded() {
	log.Print("corpus.loadExcluded enter")
	excluded = make(map[string]bool)
	excludedFile := config.CorpusDataDir() + "/exclude.txt"
	file, err := os.Open(excludedFile)
	if err != nil {
		log.Printf("corpus.loadExcluded: Error opening excluded words file, " +
			"skipping excluded words\n")
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	reader.Comma = rune('\t')
	reader.Comment = rune('#')
	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for _, row := range rawCSVdata {
		if len(row) < 1 {
			log.Fatal("corpus.loadExcluded: no columns in row")
	  	}
		excluded[row[0]] = true
	}
}