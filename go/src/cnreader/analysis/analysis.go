/*
Library for Chinese vocabulary analysis
*/
package analysis

import (
	"bufio"
	"bytes"
	"cnreader/config"
	"cnreader/corpus"
	"cnreader/dictionary"
	"container/list"
	"fmt"
	"text/template"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
	"unicode"
    "unicode/utf8"
)

// Maximum number of word frequency entries to output to the generated
// HTML file
const MAX_WF_OUTPUT = 100

// Maximum number of unknwon characters to output to the generated
// HTML file
const MAX_UNKOWN_OUTPUT = 50

// Word frequency output file
const UNIGRAM_FILE = "unigram.txt"

// The content for a corpus entry
type CorpusEntryContent struct {
	CorpusText, DateUpdated, VocabularyJSON, CollectionURL,
	CollectionTitle, AnalysisFile string
}

// Dictionary entry content struct used for writing a dictionary entry to HTML
type DictEntry struct {
	Headword dictionary.HeadwordDef
	DateUpdated string
}

// Vocabulary analysis entry for a single word
type WFResult struct {
	Freq int
	Chinese, Pinyin, English, Usage string
}

// Holds vocabulary analysis for a corpus text
type AnalysisResults struct {
	Title string
	WC, UniqueWords int
	WordFrequencies []WFResult
	UnkownnChars []string
	DateUpdated string
	MaxWFOutput int
}

// Breaks text into a list of CJK and non CJK strings
func GetChunks(text string) (list.List) {
	var chunks list.List
	cjk := ""
	noncjk := ""
	for _, character := range text {
		if IsCJKChar(string(character)) {
			if noncjk != "" {
				chunks.PushBack(noncjk)
				noncjk = ""
			} 
			cjk += string(character)
		} else if cjk != "" {
			chunks.PushBack(cjk)
			cjk = ""
			noncjk += string(character)
		} else {
			noncjk += string(character)
		}
	}
	if cjk != "" {
		chunks.PushBack(cjk)
	}
	if noncjk != "" {
		chunks.PushBack(noncjk)
	}
	return chunks
}

// Tests whether the symbol is a CJK character, excluding punctuation
// Only looks at the first charater in the string
func IsCJKChar(character string) bool {
	r := []rune(character)
	unicode.Is(unicode.Han, r[0])
	return unicode.Is(unicode.Han, r[0]) && !unicode.IsPunct(r[0])
}

// Parses a Chinese text into words
// Parameters:
// text: the string to parse
// Returns:
// tokens: the tokens for the parsed text
// vocab: a table of the unique words found in the parsed text
// wc: total word count
// usage: the first usage fo the word in the text
func ParseText(text string) (tokens list.List, vocab map[string]int, wc int,
	unknownChars []string, usage map[string]string) {
	vocab = make(map[string]int)
	unknownChars = make([]string, 0)
	usage = make(map[string]string)
	chunks := GetChunks(text)
	wdict := dictionary.GetWDict()
	//fmt.Printf("ParseText: For text %s got %d chunks\n", text, chunks.Len())
	for e := chunks.Front(); e != nil; e = e.Next() {
		chunk := e.Value.(string)
		//fmt.Printf("ParseText: chunk %s\n", chunk)
		characters := strings.Split(chunk, "")
		if !IsCJKChar(characters[0]) {
			tokens.PushBack(chunk)
			continue
		}
		for i := 0; i < len(characters); i++ {
			for j := len(characters); j > 0; j-- {
				w := strings.Join(characters[i:j], "")
				//fmt.Printf("ParseText: i = %d, j = %d, w = %s\n", i, j, w)
				if _, ok := wdict[w]; ok {
					//fmt.Printf("ParseText: found word %s, i = %d\n", w, i)
					tokens.PushBack(w)
					wc++
					vocab[w]++
					if _, ok := usage[w]; !ok {
						usage[w] = chunk
					}
					i = j - 1
					j = 0
				} else if (utf8.RuneCountInString(w) == 1) {
					//log.Printf("ParseText: found unknown character %s\n", w)
					unknownChars = append(unknownChars, w)
					tokens.PushBack(w)
					break
				}
			}
		}
	}
	return tokens, vocab, wc, unknownChars, usage
}

// Reads a Chinese text file
func ReadText(filename string) (string) {
	var text string
	if strings.HasSuffix(filename, ".html") {
		bs, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}
		text = string(bs)
	} else { // plain text file, add line breaks
		infile, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		reader := bufio.NewReader(infile)
		var buffer bytes.Buffer
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
        	if _, err = buffer.WriteString(line + "<br/>\n"); err != nil {
            	break
        	}
    	}
    	text = buffer.String()
	}
	//fmt.Printf("ReadText: read text %s\n", text)
	return text
}

// Compute word frequencies for entire corpus.
func WordFrequencies() {
	corpusDataDir := config.ProjectHome() + "/data/corpus/"
	corpusDir := config.ProjectHome() + "/corpus/"
	outfile := config.ProjectHome() + "/data/" + UNIGRAM_FILE
	f, err := os.Create(outfile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)

	// Overall word frequencies per corpus
	wcTotal := map[string]int{}
	wfTotal := map[*CorpusWord]CorpusWordFreq{}

	collectionEntries := corpus.Collections()
	for _, col := range collectionEntries {
		colFile := corpusDataDir + col.CollectionFile
		log.Printf("WordFrequencies: input file: %s\n", colFile)
		corpusEntries := corpus.CorpusEntries(colFile)
		for _, entry := range corpusEntries {
			src := corpusDir + entry.RawFile
			text := ReadText(src)
			_, vocab, wc, _, usage := ParseText(text)
			wcTotal[col.Corpus] += wc
			for word, count := range vocab {
				cw := &CorpusWord{col.Corpus, word}
				cwf := &CorpusWordFreq{col.Corpus, word, count}					
				if cwfPrev, found := wfTotal[cw]; found {
					cwf.Freq += cwfPrev.Freq			
				}
				wfTotal[cw] = *cwf
				rel_freq := 1000.0 * float64(count) / float64(wc)
				fmt.Fprintf(w, "%s\t%d\t%f\t%s\t%s\t%s\t%s\n", word, count, rel_freq,
					entry.GlossFile, col.Title, entry.Title, usage[word])
			}
		}
	}

	// Output totals for each corpus
	for corpus, count := range wcTotal {
		log.Printf("WordFrequencies: Total word count for corpus %s: %d\n",
			corpus, count)

	}
	for _, wcf := range wfTotal {
		rel_freq := 1000.0 * float64(wcf.Freq) / float64(wcTotal[wcf.Corpus])
		fmt.Fprintf(w, "%s\t%d\t%f\t%s\t%s\t%s\n", wcf.Word, wcf.Freq,
			rel_freq, "#", wcf.Corpus, "")
	}
	w.Flush()
}

// Writes a document with vocabulary analysis of the text. The name of the
// output file will be source file with '-analysis' appended, placed in the
// web/analysis directory
// vocab: The vocabulary with word frequency counts
// wc: the total size of the vocabulary
// srcFile: The source file used
// collectionTitle: The title of the whole colleciton
// docTitle: The title of this specific document
// Returns the name of the file written to
func WriteAnalysis(vocab map[string]int, usage map[string]string, wc int, 
		unknownChars []string, srcFile string, collectionTitle string,
		docTitle string) string {

	// Parse template and organize template parameters
	sortedWords := SortedFreq(vocab)
	wfResults := make([]WFResult, 0)
	maxWFOutput:= len(sortedWords)
	if maxWFOutput > MAX_WF_OUTPUT {
		maxWFOutput = MAX_WF_OUTPUT
	}
	for _, value := range sortedWords[:maxWFOutput] {
		ws, _ := dictionary.GetWordSense(value.Word)
		wfResults = append(wfResults, WFResult{value.Freq, value.Word,
			ws.Pinyin, ws.English, usage[value.Word]})
	}

	dateUpdated := time.Now().Format("2006-01-02")
	maxUnkownOutput := len(unknownChars)
	if maxUnkownOutput > MAX_UNKOWN_OUTPUT {
		maxUnkownOutput = MAX_UNKOWN_OUTPUT
	}
	title := "Vocabulary Analysis for " + collectionTitle + ", " + docTitle
	results := AnalysisResults{title, wc, len(vocab),
		wfResults, unknownChars[:maxUnkownOutput], dateUpdated,
		maxWFOutput}
	tmpl, err := template.New("corpus-analysis-template.html").ParseFiles("/Users/alex/Documents/code/chinesenotes.com/corpus/corpus-analysis-template.html")
	if err != nil { panic(err) }
	if tmpl == nil {
		log.Fatal("WriteAnalysis: Template is nil ")
		panic(err)
	}

	// Write output
	i := strings.Index(srcFile, ".txt")
	if i <= 0 {
		log.Fatal("WriteAnalysis: Bad name for source file", srcFile)
	}
	basename := srcFile[:i] + "_analysis.html"
	filename := config.ProjectHome() + "/web/analysis/" + basename
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	err = tmpl.Execute(w, results)
	if err != nil { panic(err) }
	w.Flush()
	return basename
}

// Writes a corpus document with markup for the array of tokens
// tokens: A list of tokens forming the document
// vocab: A list of word id's in the document
// filename: The file name to write to
// HTML template to use
// collectionURL: the URL of the collection that the corpus text belongs to
// collectionTitle: The collection title that the corpus entry belongs to
// aFile: The vocabulary analysis file written to or empty string for none
func WriteCorpusDoc(tokens list.List, vocab map[string]int, filename string,
	collectionURL string, collectionTitle string, aFile string) {

	var buffer bytes.Buffer

	// Iterate over text chunks
	for e := tokens.Front(); e != nil; e=e.Next() {
		chunk := e.Value.(string)
		//fmt.Printf("WriteDoc: Word %s\n", word)
		if entries, ok := dictionary.GetWord(chunk); ok {
			wordIds := ""
			for _, ws := range entries {
				if wordIds == "" {
					wordIds = fmt.Sprintf("%d", ws.Id)
				} else {
					wordIds = fmt.Sprintf("%s,%d", wordIds, ws.Id)
				}
			}
			fmt.Fprintf(&buffer, "<span title='%s' data-wordid='%s'" +
					" class='dict-entry' data-toggle='popover'>%s</span>",
					chunk, wordIds, chunk)
		} else {
			buffer.WriteString(chunk)
		}
	}

	textContent := buffer.String()
	dateUpdated := time.Now().Format("2006-01-02")
	vocabJSON := WriteVocab("", 0, vocab)
	content := CorpusEntryContent{textContent, dateUpdated, vocabJSON,
		collectionURL, collectionTitle, aFile}

	// Write to file
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	
	templFile := config.ProjectHome() + "/corpus/corpus-template.html"
	//fmt.Println("Home: ", config.ProjectHome())
	tmpl:= template.Must(template.New("corpus-template.html").ParseFiles(templFile))
	err = tmpl.Execute(w, content)
	if err != nil { panic(err) }
	w.Flush()
}

// Writes a document with markup for the array of tokens
// tokens: A list of tokens forming the document
// vocab: A list of word id's in the document
// filename: The file name to write to
func WriteDoc(tokens list.List, vocab map[string]int, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	// Iterate over text chunks
	for e := tokens.Front(); e != nil; e=e.Next() {
		chunk := e.Value.(string)
		//fmt.Printf("WriteDoc: Word %s\n", word)
		if entries, ok := dictionary.GetWord(chunk); ok {
			wordIds := ""
			for _, ws := range entries {
				if wordIds == "" {
					wordIds = fmt.Sprintf("%d", ws.Id)
				} else {
					wordIds = fmt.Sprintf("%s,%d", wordIds, ws.Id)
				}
			}
			fmt.Fprintf(w, "<span title='%s' data-wordid='%s'" +
					" class='dict-entry' data-toggle='popover'>%s</span>",
					chunk, wordIds, chunk)
		} else {
			index := strings.Index(chunk, "<!-- words here -->")
			if index == -1 {
				w.WriteString(chunk)
			} else {
				vocabJSON := WriteVocab(chunk, index, vocab)
				w.WriteString(vocabJSON)
			}
		}
	}
	w.Flush()
}

func WriteHwFiles() {
	fmt.Printf("WriteHwFiles: Begin +++++++++++\n")
	hwArray := dictionary.GetHeadwords()
	dateUpdated := time.Now().Format("2006-01-02")

	// Prepare template
	templFile := config.ProjectHome() + "/html/templates/headword-template.html"
	//fmt.Println("Home: ", config.ProjectHome())
	tmpl := template.Must(template.New("headword-template.html").ParseFiles(templFile))

	for _, hw := range hwArray {
		dictEntry := DictEntry{hw, dateUpdated}
		filename := fmt.Sprintf("%s%s%d%s", config.ProjectHome(), "/web/words/",
			hw.Id, ".html")
		f, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		w := bufio.NewWriter(f)
		err = tmpl.Execute(w, dictEntry)
		if err != nil { panic(err) }
		w.Flush()
		f.Close()
	}
}

// Writes the vocabulary out to a string in JSON format
// chunk: A chunk of text
// vocab: A list of word id's in the document
// return a JSON formatted string with the vocabulary
func WriteVocab(chunk string, index int, vocab map[string]int) string {
	buffer := bytes.NewBufferString("");
	buffer.WriteString(chunk[:index])
	buffer.WriteString("\n")
	buffer.WriteString("<script>\n")
	buffer.WriteString("words = {\n")
	for key, _ := range vocab {
		if entries, ok := dictionary.GetWord(key); ok {
			for _, ws := range entries {
				fmt.Fprintf(buffer, "\"%d\":{\"element_text\":\"%s\"," +
					"\"simplified\":\"%s\"," +
					"\"traditional\":\"%s\"," +
					"\"pinyin\":\"%s\",\"english\":\"%s\"," +
					"\"notes\":\"%s\"},\n", ws.Id, key,
					ws.Simplified, ws.Traditional, ws.Pinyin,
					ws.English, ws.Notes)
			}
		} 
	}
	buffer.WriteString("}\n")
	buffer.WriteString("</script>\n")
	buffer.WriteString(chunk[index:])
	buffer.WriteString("\n")
	return buffer.String()
}
