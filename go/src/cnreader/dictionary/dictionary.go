/*
Package for command line tool configuration
*/
package dictionary

import (
	"bufio"
	"cnreader/config"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// Headword output file
const HEADWORD_FILE = "headwords.txt"

// Defines a headword with mapping between simplified and traditional
type HeadwordDef struct {
	Id int // key for the headword
	Simplified, Traditional string
	Pinyin []string
	WordSenses []WordSenseEntry //keys for the lexical units
}

// Defines a single sense of a Chinese word
type WordSenseEntry struct {
	Id, HeadwordId int
	Simplified, Traditional, Pinyin, English, Grammar, Concept_cn,
		Concept_en, Topic_cn, Topic_en, Parent_cn, Parent_en, Image,
		Mp3, Notes string
}

// The dictionary is a map of pointers to word senses, indexed by simplified
// and traditional text
var wdict map[string][]*WordSenseEntry

// Map of parts of speech that distinguish words as function words
var functionPOS map[string]bool

// Map of words that are function words regardless of PoS
var functionalWords map[string]bool

func init() {
	functionPOS = map[string]bool{
		"adverb": true,
		"conjunction": true,
		"interjection": true,
		"interrogative pronoun": true,
		"measure word": true,
		"particle": true,
		"prefix": true,
		"preposition": true,
		"pronoun": true,
		"suffix": true,
	}	
	functionalWords = map[string]bool{
		"有": true,
		"是": true,
		"诸": true,
		"故": true,
		"出": true,
		"当": true,
		"若": true,
		"如": true,
	}
}

// Get a list of words that containst the given word
func ContainsWord(word string, headwords []HeadwordDef) []HeadwordDef {
	//log.Printf("dictionary.ContainsWord: Enter\n")
	contains := []HeadwordDef{}
	for _, hw := range headwords {
		if len(contains) <= 20 && hw.Simplified != word && strings.Contains(hw.Simplified, word) {
			contains = append(contains, hw)
		}
	}
	return contains
}

// Compute headword numbers for all lexical units listed in data/words.txt
// Return a sorted array of headwords
func GetHeadwords() []HeadwordDef {
	//log.Printf("dictionary.GetHeadwords: Enter\n")
	wsMap := readWSMap(config.DictionaryDir() + "/words.txt")

	// Read lexical units
	hwmap := make(map[string][]*WordSenseEntry)
	hwcount := 0
	for _, ws := range wdict {
		key := fmt.Sprintf("%s:%s", ws[0].Simplified, ws[0].Traditional)
		if _, ok := hwmap[key]; !ok {
			hwmap[key] = ws
			hwcount++
		}
	}
	//log.Printf("dictionary.GetHeadwords: hwcount = %d\n", hwcount)

	// Organize the headwords
	hwIdArray := make([]int, 0)
	hwIdMap := make(map[int]HeadwordDef)
	for _, senses := range hwmap {
		wsIds := []int{}
		pinyinMap := make(map[string]bool)
		for _, ws:= range senses {
			wsIds = append(wsIds, ws.Id)
			pinyinMap[ws.Pinyin] = true
		}
		sort.Ints(wsIds)
		hwId := wsIds[0]
		hwIdArray = append(hwIdArray, hwId)
		wsArray := make([]WordSenseEntry, 0)
		for _, wsId := range wsIds {
			wsArray = append(wsArray, wsMap[wsId])
		}
		pinyinArr := []string{}
		for pinyin, _ := range pinyinMap {
			pinyinArr = append(pinyinArr, pinyin)
		}
		hw := HeadwordDef{hwId, senses[0].Simplified, senses[0].Traditional,
			pinyinArr, wsArray}
		hwIdMap[hwId] = hw
	}

	// Write the headwords to the output file
	sort.Ints(hwIdArray)
	hwArray := []HeadwordDef{}
	for _, hwId := range hwIdArray {
		hw := hwIdMap[hwId]
		hwArray = append(hwArray, hw)
	}
	return hwArray
}

// Gets the dictionary definition of a word
// Parameters
//   chinese: The Chinese (simplified or traditional) text of the word
// Return
//   word: an array of word senses
//   ok: true if the word is in the dictionary
func GetWord(chinese string) (word []*WordSenseEntry, ok bool) {
	word, ok = wdict[chinese]
	return word, ok
}

// Gets the dictionary definition of the first word sense matching the word
func GetWordSense(chinese string) (WordSenseEntry, bool) {
	wSenses, ok := wdict[chinese]
	if ok {
		return *(wSenses[0]), ok
	}
	ws := new(WordSenseEntry)
	return *ws, ok
}

// Gets the dictionary, loads it if it is not loaded already
func GetWDict() map[string][]*WordSenseEntry {
	return wdict
}

// Tests whether the symbol is a CJK character, excluding punctuation
// Only looks at the first charater in the string
func IsCJKChar(character string) bool {
	r := []rune(character)
	return unicode.Is(unicode.Han, r[0]) && !unicode.IsPunct(r[0])
}

// Tests whether the word is a function word
func (ws *WordSenseEntry) IsFunctionWord() bool {
	return functionalWords[ws.Simplified] || functionPOS[ws.Grammar]
}

// Tests whether the word string contains a number
func (ws *WordSenseEntry) IsNumericExpression() bool {
	if ws.Grammar == "number" {
		return true
	}
	if ws.Grammar == "phrase" ||  ws.Grammar == "set phrase" {
		for _, r := range []rune(ws.Simplified) {
			wsChar := wdict[string(r)]
			if wsChar[0].Grammar == "number" {
				return true
			}
		}
	}
	return false
}

// IsProperNoun tests whether the word is a function word.
// If the majority of word senses are proper nouns, then the word is marked
// as a proper noun.
func (ws *WordSenseEntry) IsProperNoun() bool {
	if wsArray, ok := wdict[ws.Simplified]; ok {
		count := 0
		for _, s := range wsArray {
			if s.Grammar == "proper noun" {
				count++
			}
		}
		return float64(count) / float64(len(wsArray)) > 0.5
	}
	return ws.Grammar == "proper noun"
}

// Reads the Chinese-English lexical units into memory from the words.txt file
// Parameters:
//   wsfilename The name of the word sense file
func ReadDict(wsfilename string) {
	log.Printf("dictionary.ReadDict: wsfilename: %s\n", wsfilename)
	wsfile, err := os.Open(wsfilename)
	if err != nil {
		log.Fatal(err)
	}
	defer wsfile.Close()
	reader := csv.NewReader(wsfile)
	reader.FieldsPerRecord = -1
	reader.Comma = rune('\t')
	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Could not parse lexical units file", err)
	}
	wdict = make(map[string][]*WordSenseEntry)
	for i, row := range rawCSVdata {
		id, err := strconv.ParseInt(row[0], 10, 0)
		if err != nil {
			log.Fatal("Could not parse word id for word ", i, err)
		}
		simp := row[1]
		trad := row[2]
		pinyin := row[3]
		english := row[4]
		grammar := row[5]
		conceptCn := row[6]
		hwId := 0
		if len(row) == 16 {
			hwIdInt, err := strconv.ParseInt(row[15], 10, 0)
			if err != nil {
				log.Printf("ReadDict, id: %d, simp: %s, trad: %s, " + 
					"pinyin: %s, english: %s, grammar: %s, conceptCn: %s\n",
					id, simp, trad, pinyin, english, grammar, conceptCn)
				log.Fatal("ReadDict: Could not parse headword id for word ",
					id, err)
			}
			hwId = int(hwIdInt)
		} else {
			log.Printf("ReadDict, No. cols: %d\n",len(row))
			log.Printf("ReadDict, id: %d, simp: %s, trad: %s, pinyin: %s, " +
				"english: %s, grammar: %s, conceptCn: %s\n",
				id, simp, trad, pinyin, english, grammar, conceptCn)
			log.Fatal("ReadDict wrong number of columns ", id, err)
		}
		newWs := &WordSenseEntry{Id: int(id),
				HeadwordId: int(hwId),
				Simplified: simp,
				Traditional: trad,
				Pinyin: pinyin,
				English: english,
				Grammar: grammar,
				Concept_cn: conceptCn,
				Concept_en: row[7], 
				Topic_cn: row[8],
				Topic_en: row[9],
				Parent_cn: row[10],
				Parent_en: row[11],
				Image: row[12],
				Mp3: row[13],
				Notes: row[14]}
		if trad != "\\N" {
			wSenses, ok := wdict[trad]
			if !ok {
				wsSlice := make([]*WordSenseEntry, 1)
				wsSlice[0] = newWs
				wdict[trad] = wsSlice
			} else {
				wdict[trad] = append(wSenses, newWs)
			}
		}
		wSenses, ok := wdict[simp]
		if !ok {
			wsSlice := make([]*WordSenseEntry, 1)
			wsSlice[0] = newWs
			wdict[simp] = wsSlice
		} else {
			//fmt.Printf("ReadDict: found simplified %s already in dict\n", simp)
			wdict[simp] = append(wSenses, newWs)
		}
	}
}


// Reads the Chinese-English lexical units into memory from the words.txt file
// and returns a map from word sense id to word sense definition
// Parameters:
//   wsfilename The name of the word sense file
func readWSMap(wsfilename string) map[int]WordSenseEntry {
	wsfile, err := os.Open(wsfilename)
	if err != nil {
		log.Fatal(err)
	}
	defer wsfile.Close()
	reader := csv.NewReader(wsfile)
	reader.FieldsPerRecord = -1
	reader.Comma = rune('\t')
	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	wsMap := make(map[int]WordSenseEntry)
	for _, row := range rawCSVdata {
		id, error := strconv.ParseInt(row[0], 10, 0)
		if error != nil {
			log.Fatal("readWSMap: Could not parse word id for row %s\n", row[0])
		}
		simp := row[1]
		trad := row[2]
		ws := WordSenseEntry{Id: int(id), Simplified: simp,
				Traditional: trad, Pinyin: row[3], English: row[4],
				Grammar: row[5], Concept_cn: row[6], Concept_en: row[7], 
				Topic_cn: row[8], Topic_en: row[9], Parent_cn: row[10],
				Parent_en: row[11], Image: row[12], Mp3: row[13],
				Notes: row[14]}
		wsMap[int(id)] = ws
	}
	return wsMap
}

// Compute headword numbers for all lexical units listed in data/words.txt,
//writing to the headword.txt file.
func WriteHeadwords() {

	hwArray := GetHeadwords()

	// Prepare head words file for writing
	outfile := config.DictionaryDir() + "/" + HEADWORD_FILE
	f, err := os.Create(outfile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)

	// Write the headwords to the output file
	for _, hw := range hwArray {
		wsNumStr := ""
		for i, ws := range hw.WordSenses {
			if i == 0 {
				wsNumStr = fmt.Sprintf("%d", ws.Id)
			} else {
				wsNumStr = fmt.Sprintf("%s, %d", wsNumStr, ws.Id)
			}
		}
		pinyinStr := ""
		for i, pinyin := range hw.Pinyin {
			if i == 0 {
				pinyinStr = fmt.Sprintf("%s", pinyin)
			} else {
				pinyinStr = fmt.Sprintf("%s, %s", pinyinStr, pinyin)
			}
		}
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\n", hw.Id, hw.Simplified,
			hw.Traditional, pinyinStr, wsNumStr)
	}
	w.Flush()
}