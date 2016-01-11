/* 
Utility for transcoding from traditional Chinese into the format required
by the dictionary for lexical units. Depends on the cnreader go library.
 */
package main

import (
	"bufio"
	"cnreader/config"
	"cnreader/dictionary"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//Entry point for the command line utility.
func main() {
	log.Printf("main: enter\n")
	var infileStr = flag.String("infile", "testinput.txt",
		"Input file")
	var outfileStr = flag.String("outfile", "testoutput.txt",
		"Output file")
	notesDef := "\tphrase\t\\N\t\\N\t佛教\tBuddhism\t佛光山\tFo Guang Shan\t\\N" +
		"\t\\N\tVenerable Master Hsing Yun's One-Stroke Calligraphy, " +
		"translation by: Ven. Miao Guang 妙光 (FoguangPedia)"
	var notesStr = flag.String("notes", notesDef, "Remainder of the line")
	var startatStr = flag.String("startat", "0",
		"Integer to start word id's at")
	flag.Parse()
	startat, err := strconv.Atoi(*startatStr)
	if err != nil {
		log.Fatal("Could not parse startat to an integer", *startatStr, err)
	}

	config.SetProjectHome("../../..")
	trad2LU(*infileStr, *outfileStr, *notesStr, startat)
}

//Read the input file and add details to the output file
func trad2LU(infileStr, outfileStr, notes string, startat int) {
	log.Printf("trad2LU: %s, %s, %d\n", infileStr, outfileStr, startat)

	// input file
	infile, err := os.Open(infileStr)
	if err != nil {
		log.Fatal(err)
	}
	defer infile.Close()
	reader := csv.NewReader(infile)
	reader.FieldsPerRecord = -1
	reader.Comma = rune('\t')
	reader.Comment = rune('#')
	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Output file
	outfile, err := os.Create(outfileStr)
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()
	w := bufio.NewWriter(outfile)

	// Open dictionary
	dictionary.ReadDict(config.LUFileName())

	i := 0
	for _, row := range rawCSVdata {
        i++
		if len(row) != 2 {
			if len(row) > 0 {
				log.Printf("Problem on line %d: %s", i, row[0])
			}
			log.Fatal("Wrong number of fields for on line ", i, len(row))
		}
		trad := row[0]
		english := row[1]
		j := startat + i - 1
		simplified, pinyin := getSimpPinyin(trad)
        lu := fmt.Sprintf("%d\t%s\t%s\t%s\t%s%s\n", j, simplified, trad, pinyin,
        	english, notes)
        if _, err = w.WriteString(lu); err != nil {
        	log.Fatal("Error on line ", i, err)
        }
    }
	w.Flush()
}

// Gets simplified Chinese and pinyin for the traditional text string
func getSimpPinyin(trad string) (string, string) {
	simplified := ""
	pinyin := ""
	for _, character := range trad {
		ch := string(character)
		if dictionary.IsCJKChar(ch) {
			ws, ok := dictionary.GetWordSense(ch)
			if !ok {
				log.Fatal("Could not find entry for ", trad)
			}
			simplified = simplified + ws.Simplified
			pinyin = pinyin + " " + ws.Pinyin
		} else {
			simplified = simplified + string(character)
			if ch == "，" {
				pinyin = pinyin + ","
			} else if ch == "；" {
				pinyin = pinyin + ";"
			} else if ch == " " {
				pinyin = pinyin + ""
			} else if ch == "?" {
				pinyin = pinyin + "?"
			} else if ch == "（" {
				pinyin = pinyin + "("
			} else if ch == "）" {
				pinyin = pinyin + ")"
			} else if ch == "！" {
				pinyin = pinyin + "!"
			} else {
				log.Fatal("Confused about puncuation for ", trad)
			}
		}
	}
	return simplified, strings.TrimSpace(pinyin)
}
