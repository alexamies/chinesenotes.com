/*
Package for command line tool configuration
*/
package dictionary

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/alexamies/cnreader/config"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// Headword output file
const HEADWORD_FILE = "headwords.txt"

// Maximum number of containing words to output to the generated
// HTML file
const MAX_CONTAINS = 50

// The dictionary is a map of pointers to word senses, indexed by simplified
// and traditional text
var wdict map[string][]*WordSenseEntry

var hwIdMap map[int]HeadwordDef

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

// Clones the headword definition without the attached array of word senses
func CloneHeadword(hw HeadwordDef) HeadwordDef {
	wsArray := []WordSenseEntry{}
	return HeadwordDef{
		Id: hw.Id,
		Simplified: hw.Simplified,
		Traditional: hw.Traditional,
		Pinyin: hw.Pinyin,
		WordSenses: &wsArray,
	}
}

// Get a list of words that containst the given word
func ContainsWord(word string, headwords []HeadwordDef) []HeadwordDef {
	//log.Printf("dictionary.ContainsWord: Enter\n")
	contains := []HeadwordDef{}
	for _, hw := range headwords {
		if len(contains) <= MAX_CONTAINS && *hw.Simplified != word && strings.Contains(*hw.Simplified, word) {
			contains = append(contains, hw)
		}
	}
	return contains
}

// Filter the list of headwords by the given domain
// Parameters
//   hws: A list of headwords
//   domain_en: the domain to filter by, ignored if empty
// Return
//   hw: an array of headwords matching the domain, with senses not matching the
//       domain also removed
func FilterByDomain(hws []HeadwordDef, domain_en string) []HeadwordDef {
	if domain_en == "" {
		return hws
	}
	headwords := []HeadwordDef{}
	for _, hw := range hws {
		wsArr := []WordSenseEntry{}
		for _, ws := range *hw.WordSenses {
			if ws.Topic_en == domain_en {
				wsArr = append(wsArr, ws)
			}
		}
		if len(wsArr) > 0 {
			h := CloneHeadword(hw)
			h.WordSenses = &wsArr
			headwords = append(headwords, h)
		}
	}
	return headwords
}

// Gets the dictionary definition of a word
// Parameters
//   chinese: The Chinese (simplified or traditional) text of the word
// Return
//   hw: the headword for the string
//   ok: true if the word is in the dictionary
func GetHeadword(chinese string) (hw HeadwordDef, ok bool) {
	wsArray, ok := GetWord(chinese)
	if ok && len(wsArray) > 0 {
		hw := hwIdMap[wsArray[0].HeadwordId]
		return hw, ok
	}
	empty := ""
	wsArr0 := []WordSenseEntry{}
	return HeadwordDef{-1, &empty, &empty, []string{}, &wsArr0}, ok
}

// Compute headword numbers for all lexical units listed in data/words.txt
// Return a sorted array of headwords
func GetHeadwords() []HeadwordDef {
	//log.Printf("dictionary.GetHeadwords: Enter\n")

	// Read lexical units
	hwmap := make(map[int][]*WordSenseEntry)
	hwcount := 0
	for _, ws := range wdict {
		for _, lu := range ws {
			key := lu.HeadwordId
			//if key == 9806 {
				//log.Printf("dictionary.GetHeadwords: key == 9806\n")
			//}
			if _, ok := hwmap[key]; !ok {
				hwmap[key] = ws
				hwcount++
			}
		}
	}
	log.Printf("dictionary.GetHeadwords: hwcount = %d\n", hwcount)

	// Organize the headwords
	hwIdArray := make([]int, 0)
	hwIdMap = make(map[int]HeadwordDef)
	for hwId, senses := range hwmap {
		wsIds := []int{}
		pinyinMap := make(map[string]bool)
		for _, ws:= range senses {
			if ws.HeadwordId == hwId {
				wsIds = append(wsIds, ws.Id)
				pinyinMap[ws.Pinyin] = true
			} //else {
			//	log.Printf("dictionary.GetHeadwords: hw.Id != ws.HeadwordId: %d, %d\n",
			//		hwId, ws.HeadwordId)
			//}
		}
		simplified := &senses[0].Simplified
		traditional := &senses[0].Traditional
		sort.Ints(wsIds)
		hwIdArray = append(hwIdArray, hwId)
		wsArray := make([]WordSenseEntry, 0)
		for _, sense := range wdict[*simplified] {
			wsArray = append(wsArray, *sense)
		}
		pinyinArr := []string{}
		for pinyin, _ := range pinyinMap {
			pinyinArr = append(pinyinArr, pinyin)
		}
		// In case the simplified form maps to multiple traditional variants
		//for _, sense := range senses {
		//	if sense.HeadwordId == hwId {
		//		traditional = &sense.Traditional
		//		break
		//	}
		//}
		hw := HeadwordDef{hwId, simplified, traditional,
			pinyinArr, &wsArray}
		hwIdMap[hwId] = hw
		//if hwId == 821 {
		//	fmt.Printf("dictionary.GetHeadwords: %s, %s, %d\n", *simplified,
		//			*traditional, len(wsArray))
		//}
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

func GetHwMap() map[int]HeadwordDef {
	if hwIdMap == nil {
		GetHeadwords()
	}
	return hwIdMap
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

// Gets the dictionary, 
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
			wsChar, ok := wdict[string(r)]
			if ok && wsChar[0].Grammar == "number" {
				return true
			} else if !ok {
				log.Printf("dictionary.IsNumericExpression not found: '%s'", 
					string(r))
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

// Reads the Chinese-English lexical units into memory from the lexical unit 
// files.
// Parameters:
//   wsfilename The name of the word sense file
func ReadDict(wsFilenames []string) []HeadwordDef {
	wdict = make(map[string][]*WordSenseEntry)
	avoidSub := config.AvoidSubDomains()
	for _, wsfilename := range wsFilenames {
		log.Printf("dictionary.ReadDict: wsfilename: %s\n", wsfilename)
		wsfile, err := os.Open(wsfilename)
		if err != nil {
			log.Fatal("dictionary.ReadDict, error: ", err)
		}
		defer wsfile.Close()
		reader := csv.NewReader(wsfile)
		reader.FieldsPerRecord = -1
		reader.Comma = rune('\t')
		reader.Comment = '#'
		rawCSVdata, err := reader.ReadAll()
		if err != nil {
			log.Fatal("Could not parse lexical units file", err)
		}
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
			conceptEn := row[7]
			topicCn := row[8]
			topicEn := row[9]
			hwId := 0
			if len(row) == 16 {
				hwIdInt, err := strconv.ParseInt(row[15], 10, 0)
				if err != nil {
					log.Printf("ReadDict, id: %d, simp: %s, trad: %s, " + 
						"pinyin: %s, english: %s, grammar: %s, conceptCn: %s" +
						", conceptEn: %s, topicCn: %s, topicEn: %s\n",
						id, simp, trad, pinyin, english, grammar, conceptCn,
						conceptEn, topicCn, topicEn)
					log.Fatal("ReadDict: Could not parse headword id for word ",
						id, err)
				}
				hwId = int(hwIdInt)
			} else {
				log.Printf("ReadDict, No. cols: %d\n",len(row))
					log.Printf("ReadDict, id: %d, simp: %s, trad: %s, " + 
						"pinyin: %s, english: %s, grammar: %s, conceptCn: %s" +
						", conceptEn: %s, topicCn: %s, topicEn: %s\n",
						id, simp, trad, pinyin, english, grammar, conceptCn,
						conceptEn, topicCn, topicEn)
				log.Printf("ReadDict, Line: %s\n", strings.Join(row, ";"))
				log.Fatalf("ReadDict, line %d, wrong number of columns: %d", id, len(row))
			}
			parent_en :=  row[11]
			// If subdomain, aka parent, should be avoided, then skip
			if _, ok := avoidSub[parent_en]; ok {
				continue
			}
			newWs := &WordSenseEntry{Id: int(id),
				HeadwordId: int(hwId),
				Simplified: simp,
				Traditional: trad,
				Pinyin: pinyin,
				English: english,
				Grammar: grammar,
				Concept_cn: conceptCn,
				Concept_en: conceptEn, 
				Topic_cn: topicCn,
				Topic_en: topicEn,
				Parent_cn: row[10],
				Parent_en: parent_en,
				Image: row[12],
				Mp3: row[13],
				Notes: row[14]}
			if trad != "\\N" && trad != simp {
				wSenses, ok := wdict[trad]
				if !ok {
					// Multiple mappings for simplified to traditional
					wSenses1, ok1 := wdict[simp]
					if !ok1 {
						wsSlice := make([]*WordSenseEntry, 1)
						wsSlice[0] = newWs
					 	wdict[trad] = wsSlice
					} else {
						wdict[trad] = append(wSenses1, newWs)
					}
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
				wdict[simp] = append(wSenses, newWs)
			}
			//if int(hwId) == 821 {
			//	fmt.Printf("ReadDict: %d found simp %s in dict, len %d\n",
			//		id, simp, len(wdict[simp]))
			//}
		}
	}
	return GetHeadwords()
}

// Reads the Chinese-English lexical units into memory from the words.txt file
// and returns a map from word sense id to word sense definition
// Parameters:
//   wsfilename The name of the word sense file
func readWSMap(wsFilenames []string) map[int]WordSenseEntry {
	wsMap := make(map[int]WordSenseEntry)
	for _, wsfilename := range wsFilenames {
		wsfile, err := os.Open(wsfilename)
		if err != nil {
			log.Fatal(err)
		}
		defer wsfile.Close()
		reader := csv.NewReader(wsfile)
		reader.FieldsPerRecord = -1
		reader.Comma = rune('\t')
		reader.Comment = '#'
		rawCSVdata, err := reader.ReadAll()
		if err != nil {
			log.Fatal(err)
		}
		for _, row := range rawCSVdata {
			id, error := strconv.ParseInt(row[0], 10, 0)
			if error != nil {
				log.Fatal("readWSMap: Could not parse word id for row \n",
					row[0])
			}
			simp := row[1]
			trad := row[2]
			hwId, error := strconv.ParseInt(row[15], 10, 0)
			if error != nil {
				log.Fatal("readWSMap: Could not parse hwId for row \n", row[0])
			}
			ws := WordSenseEntry{Id: int(id), Simplified: simp,
				Traditional: trad, Pinyin: row[3], English: row[4],
				Grammar: row[5], Concept_cn: row[6], Concept_en: row[7], 
				Topic_cn: row[8], Topic_en: row[9], Parent_cn: row[10],
				Parent_en: row[11], Image: row[12], Mp3: row[13],
				Notes: row[14], HeadwordId: int(hwId)}
			wsMap[int(id)] = ws
		}
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
		for i, ws := range *hw.WordSenses {
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
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\n", hw.Id, *hw.Simplified,
			*hw.Traditional, pinyinStr, wsNumStr)
	}
	w.Flush()
}