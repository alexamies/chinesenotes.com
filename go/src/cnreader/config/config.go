/*
Package for command line tool configuration
*/
package config

import (
	"encoding/csv"
	"log"
	"os"
)

const conversionsFile = "data/corpus/html-conversion.csv"

var projectHome string

// A type that holds the source and destination files for HTML conversion
type HTMLConversion struct {
	SrcFile, DestFile string
}

func init() {
	projectHome = "../../../.."
}

// Gets the public home directory, relative to the cnreader command line tool
func ProjectHome() string {
	return projectHome
}

// Gets the list of source and destination files for HTML conversion
func GetHTMLConversions() []HTMLConversion {
	conversionsFile := projectHome + "/" + conversionsFile
	convFile, err := os.Open(conversionsFile)
	if err != nil {
		log.Fatal(err)
	}
	defer convFile.Close()
	reader := csv.NewReader(convFile)
	reader.FieldsPerRecord = -1
	reader.Comma = rune('\t')
	reader.Comment = rune('#')
	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	conversions := make([]HTMLConversion, 0)
	for _, row := range rawCSVdata {
		conversions = append(conversions, HTMLConversion{row[0], row[1]})
	}
	return conversions
}

// Sets the public home directory, relative to the cnreader command line tool
func SetProjectHome(home string) {
	projectHome = home
}
