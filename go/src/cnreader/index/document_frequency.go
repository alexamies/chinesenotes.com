/*
For every term, store the number of documents that contain the term
*/
package index

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/alexamies/cnreader/config"
	"log"
	"math"
	"os"
	"strconv"
)

// File name for document index
const DOC_FREQ_FILE = "doc_freq.txt"
const BIGRAM_DOC_FREQ_FILE = "bigram_doc_freq.txt"

// Map from term to number of documents referencing the term
type DocumentFrequency struct {
	DocFreq map[string]int
	N       int // total number of documents
}

// Loaded from disk in contrast to partially ready and still accumulating data
var completeDF DocumentFrequency

func init () {
	df, err := ReadDocumentFrequency()
	if err != nil {
		log.Println("index.init Error reading document frequency continuing")
	}
	completeDF = df
}

// Initializes a DocumentFrequency struct
func NewDocumentFrequency() DocumentFrequency {
	return DocumentFrequency{
		DocFreq: map[string]int{},
		N: 0,
	}
}

// Adds the given vocabulary to the map and increments the document count
// Param:
//   vocab - word frequencies are ignored, only the presence of the term is 
//           important
func (df *DocumentFrequency) AddVocabulary(vocab map[string]int) {
	for k, _ := range vocab {
		_, ok := df.DocFreq[k]
		if ok {
			df.DocFreq[k]++
		} else {
			df.DocFreq[k] = 1
		}
	}
	df.N += 1
}

// Merges the given document frequency to the map and increments the counts
// Param:
//   vocab - word frequencies are ignored, only the presence of the term is 
//           important
func (df *DocumentFrequency) AddDocFreq(otherDF DocumentFrequency) {
	for k, v := range otherDF.DocFreq {
		count, ok := df.DocFreq[k]
		if ok {
			df.DocFreq[k] = count + v
		} else {
			df.DocFreq[k] = 1
		}
	}
	df.N += otherDF.N
}

// Computes the inverse document frequency for the given term
// Param:
//   term: the term to find the idf for
func (df *DocumentFrequency) IDF(term string) (val float64, ok bool) {
	ndocs, ok := df.DocFreq[term]
	if ok && ndocs > 0 {
		val = math.Log10(float64(df.N + 1) / float64(ndocs))
	//log.Println("index.IDF: term, val, df.n, ", term, val, df.N)
	} 
	return val, ok
}

// Writes a document frequency object from a json file
func ReadDocumentFrequency() (df DocumentFrequency, e error) {
	dir := config.IndexDir()

	fname := dir + "/" + DOC_FREQ_FILE
	dfFile, err := os.Open(fname)
	if err != nil {
		log.Println("index.ReadDocumentFrequency, error opening word freq file: ",
			err)
		return df, err
	}
	defer dfFile.Close()
	reader := csv.NewReader(dfFile)
	reader.FieldsPerRecord = -1
	reader.Comma = rune('\t')
	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		log.Fatal("index.ReadDocumentFrequency: Could not wf file ", err)
	}
	dfMap := map[string]int{}
	for i, row := range rawCSVdata {
		w := row[0] // Chinese text for word
		count, err := strconv.ParseInt(row[1], 10, 0)
		if err != nil {
			log.Println("Could not parse word count ", i, err)
			return df, err
		}
		dfMap[w] = int(count)
	}
	df = DocumentFrequency{dfMap, -1}
	return df, err
}

// term frequency - inverse document frequency for the string
// Params
//   term: The term (word) to compute the tf-idf from
//   count: The count of the word in a specific document
func tfIdf(term string, count int) (val float64, ok bool) {
	idf, ok := completeDF.IDF(term)
	//log.Println("index.tfIdf: idf, term, ", idf, term)
	if ok {
		val = float64(count) * idf
	} else {
		//log.Println("index.tfIdf: could not compute tf-idf for, ", term)
	}
	return val, ok
}

// Writes the document frequency to json file
func (df *DocumentFrequency) WriteToFile(filename string) {
	dir := config.IndexDir()
	fname := dir + "/" + filename
	log.Println("index.DocumentFrequency.WriteToFile: N, ", df.N)
	f, err := os.Create(fname)
	if err != nil {
		log.Println("index.DocumentFrequency.WriteToFile: error, ", err)
		return
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	for k, v := range df.DocFreq {
		fmt.Fprintf(writer, "%s\t%d\n", k, v)
	}
	writer.Flush()
}