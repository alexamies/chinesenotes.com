/* 
Utility for adding headword numbers to the words.txt file 
 */
package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"strconv"
)

//Entry point for the command line utility.
func main() {
	log.Printf("main: Reading words file\n")
	hwMap := ReadHeadwords()
	addHeadwordId(hwMap)
}

//Read the words file and add head word numbers
func addHeadwordId(hwMap map[int]int) {
	log.Printf("addHeadwordId: Enter\n")

	// input file
	infile, err := os.Open("../../../data/words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer infile.Close()
	reader := bufio.NewReader(infile)

	// Output file
	outfile, err := os.Create("../../../data/lexical_units.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()
	w := bufio.NewWriter(outfile)

    eof := false
    for !eof {
        var line string
        line, err = reader.ReadString('\n')
        if err == io.EOF {
            err = nil
            eof = true
        } else if err != nil {
            break
        }
        fieldArr := strings.Split(line, "\t")
        id, err := strconv.Atoi(fieldArr[0])
		if err != nil {
			log.Fatal(err)
		}
        length := len(line)
        headwordId := hwMap[id]
        modLine := fmt.Sprintf(line[:length-1] + "\t%d\n", headwordId)
        if _, err = w.WriteString(modLine); err != nil {
        	log.Fatal(err)
        }
    }
	w.Flush()
}

//Read the words file and add head word numbers
//Returns a map of lexical unit id's to headword id's
func ReadHeadwords() map[int]int {
	log.Printf("ReadHeadwords: enter\n")
	hwMap := make(map[int]int)

	file, err := os.Open("../../../data/headwords.txt")
	if err != nil {
		log.Fatal(err)
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
		id, err := strconv.Atoi(row[0])
		if err != nil {
			log.Printf("ReadHeadwords: could not parse headword %s\n", row[0])
		}

		luArr := strings.Split(row[4], ",")
		for _, val := range luArr {
			luId, err := strconv.Atoi(strings.Trim(val, " "))
			if err != nil {
				log.Printf("ReadHeadwords: could not parse lu %s\n", row[0])
			}
			hwMap[luId] = id
		}
	}
	return hwMap
}