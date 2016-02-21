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
    "unicode/utf8"
)

// Maximum number of word frequency entries to output to the generated
// HTML file
const MAX_WF_OUTPUT = 100000

// Maximum number of unknwon characters to output to the generated
// HTML file
const MAX_UNKOWN_OUTPUT = 50

// Word frequency output file
const UNIGRAM_FILE = "unigram.txt"

// Holds vocabulary analysis for a corpus text
type AnalysisResults struct {
	Title string
	WC, UniqueWords int
	WordFrequencies []WFResult
	UnkownnChars []SortedWordItem
	DateUpdated string
	MaxWFOutput int
}

// The content for a corpus entry
type CorpusEntryContent struct {
	CorpusText, DateUpdated, CollectionURL, CollectionTitle, AnalysisFile string
}

// Dictionary entry content struct used for writing a dictionary entry to HTML
type DictEntry struct {
	Headword dictionary.HeadwordDef
	UsageArr []WordUsage
	DateUpdated string
}

// Word usage
type WordUsage struct {
	Freq int
	RelFreq float64
	Word, Example, File, EntryTitle, ColTitle string
}

// Vocabulary analysis entry for a single word
type WFResult struct {
	Freq, HeadwordId int
	Chinese, Pinyin, English, Usage string
}

// HTML content for template
type HTMLContent struct {
	Content, VocabularyJSON, DateUpdated string
}

/* Break usage example text into links with highlight on headword
   Parameters:
      usageText - usage example or notes on lexical unit
      headword - Includes simplified and traditional Chinese text
   Return
      marked up text with links and highlight
*/
func decodeUsageExample(usageText string, headword dictionary.HeadwordDef) string {
	tokens, _ := ParseText(usageText)
	replacementText := ""
	for e := tokens.Front(); e != nil; e = e.Next() {
		word := e.Value.(string)
		if word == headword.Simplified || word == headword.Traditional {
			replacementText = replacementText +
				"<span class='usage-highlight'>" + word + "</span>"
		} else {
			ws, ok := dictionary.GetWordSense(word)
			if ok {
				replacementText = replacementText + hyperlink(ws, word)
			} else {
				replacementText = replacementText + word
			}
		}
	}
	return replacementText
}

// Breaks text into a list of CJK and non CJK strings
func GetChunks(text string) (list.List) {
	var chunks list.List
	cjk := ""
	noncjk := ""
	for _, character := range text {
		if dictionary.IsCJKChar(string(character)) {
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

// Compute word frequencies for entire corpus
func GetWordFrequencies() (map[string]*[]WordUsage,
		map[*CorpusWord]CorpusWordFreq, map[string]int) {

	// Overall word frequencies per corpus
	usageMap := map[string]*[]WordUsage{}
	wcTotal := map[string]int{}
	wfTotal := map[*CorpusWord]CorpusWordFreq{}
	corpusDir := config.ProjectHome() + "/corpus/"
	corpusDataDir := config.ProjectHome() + "/data/corpus/"

	collectionEntries := corpus.Collections()
	for _, col := range collectionEntries {
		colFile := corpusDataDir + col.CollectionFile
		log.Printf("GetWordFrequencies: input file: %s\n", colFile)
		corpusEntries := corpus.CorpusEntries(colFile)
		for _, entry := range corpusEntries {
			src := corpusDir + entry.RawFile
			text := ReadText(src)
			_, results := ParseText(text)
			wcTotal[col.Corpus] += results.WC
			for word, count := range results.Vocab {
				cw := &CorpusWord{col.Corpus, word}
				cwf := &CorpusWordFreq{col.Corpus, word, count}					
				if cwfPrev, found := wfTotal[cw]; found {
					cwf.Freq += cwfPrev.Freq			
				}
				wfTotal[cw] = *cwf
				rel_freq := 1000.0 * float64(count) / float64(results.WC)
				usage := WordUsage{cwf.Freq, rel_freq, word, results.Usage[word],
					entry.GlossFile, entry.Title, col.Title}
				usageArr, ok := usageMap[word]
				if !ok {
					usageArr = new([]WordUsage)
					usageMap[word] = usageArr
				}
				*usageArr = append(*usageArr, usage)
				//fmt.Fprintf(w, "%s\t%d\t%f\t%s\t%s\t%s\t%s\n", word, count, rel_freq,
				//	entry.GlossFile, col.Title, entry.Title, usage[word])
			}
		}
	}

	// Print out totals for each corpus
	for corpus, count := range wcTotal {
		log.Printf("WordFrequencies: Total word count for corpus %s: %d\n",
			corpus, count)
	}
	return usageMap, wfTotal, wcTotal
}

// Constructs a hyperlink for a headword, including Pinyin and English in the
// title attribute for the link mouseover
func hyperlink(entry dictionary.WordSenseEntry, text string) string {
	mouseover := fmt.Sprintf("%s | %s", entry.Pinyin, entry.English)
	link := fmt.Sprintf("/words/%d.html", entry.HeadwordId)
	return fmt.Sprintf("<a title='%s' href='%s'>%s</a>", mouseover, link, text)
}

// Parses a Chinese text into words
// Parameters:
// text: the string to parse
// Returns:
// tokens: the tokens for the parsed text
// vocab: a table of the unique words found in the parsed text
// wc: total word count
// usage: the first usage of the word in the text
func ParseText(text string) (tokens list.List, results CollectionAResults) {
	vocab := map[string]int{}
	unknownChars := map[string]int{}
	usage := map[string]string{}
	wc := 0
	chunks := GetChunks(text)
	wdict := dictionary.GetWDict()
	//fmt.Printf("ParseText: For text %s got %d chunks\n", text, chunks.Len())
	for e := chunks.Front(); e != nil; e = e.Next() {
		chunk := e.Value.(string)
		//fmt.Printf("ParseText: chunk %s\n", chunk)
		characters := strings.Split(chunk, "")
		if !dictionary.IsCJKChar(characters[0]) {
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
					unknownChars[w]++
					tokens.PushBack(w)
					break
				}
			}
		}
	}
	results = CollectionAResults {
		Vocab: vocab,
		Usage: usage,
		WC: wc,
		UnknownChars: unknownChars,
	}
	return tokens, results
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
		defer infile.Close()
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

// Write out word frequencies and example use for the entire corpus
func WordFrequencies() {
	usageMap, wfTotal, wcTotal := GetWordFrequencies()
	outfile := config.ProjectHome() + "/data/" + UNIGRAM_FILE
	f, err := os.Create(outfile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)

	for word, usageArr := range usageMap {
		for _, usage := range *usageArr {
			fmt.Fprintf(w, "%s\t%d\t%f\t%s\t%s\t%s\t%s\n", word,
				usage.Freq, usage.RelFreq, usage.File, usage.EntryTitle,
				usage.ColTitle, usage.Example)
		}
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
// results: The results of vocabulary analysis
// collectionTitle: The title of the whole colleciton
// docTitle: The title of this specific document
// Returns the name of the file written to
func WriteAnalysis(results CollectionAResults, srcFile, collectionTitle,
		docTitle string) string {

	// Parse template and organize template parameters
	sortedWords := SortedFreq(results.Vocab)
	wfResults := make([]WFResult, 0)
	sortedUnknownWords := SortedFreq(results.UnknownChars)
	maxWFOutput:= len(sortedWords)
	if maxWFOutput > MAX_WF_OUTPUT {
		maxWFOutput = MAX_WF_OUTPUT
	}
	for _, value := range sortedWords[:maxWFOutput] {
		ws, _ := dictionary.GetWordSense(value.Word)
		wfResults = append(wfResults, WFResult{
			Freq: value.Freq,
			HeadwordId: ws.HeadwordId,
			Chinese: value.Word,
			Pinyin: ws.Pinyin, 
			English: ws.English, 
			Usage: results.Usage[value.Word]})
	}

	dateUpdated := time.Now().Format("2006-01-02")
	maxUnkownOutput := len(results.UnknownChars)
	if maxUnkownOutput > MAX_UNKOWN_OUTPUT {
		maxUnkownOutput = MAX_UNKOWN_OUTPUT
	}
	title := "Vocabulary Analysis for " + collectionTitle + ", " + docTitle
	aResults := AnalysisResults{title, results.WC, len(results.Vocab),
		wfResults, sortedUnknownWords, dateUpdated, maxWFOutput}
	tmplFile := config.TemplateDir() + "/corpus-analysis-template.html"
	tmpl, err := template.New("corpus-analysis-template.html").ParseFiles(tmplFile)
	if err != nil { panic(err) }
	if tmpl == nil {
		log.Fatal("WriteAnalysis: Template is nil", err)
	}

	// Write output
	i := strings.Index(srcFile, ".txt")
	if i <= 0 {
		i = strings.Index(srcFile, ".html")
		if i <= 0 {
			i = strings.Index(srcFile, ".csv")
			if i <= 0 {
				log.Fatal("WriteAnalysis: Bad name for source file: ", srcFile)
			}
		}
	}
	basename := srcFile[:i] + "_analysis.html"
	filename := config.ProjectHome() + "/web/analysis/" + basename
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	err = tmpl.Execute(w, aResults)
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

	var b bytes.Buffer

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
			/*
			// Popover
			fmt.Fprintf(&buffer, "<span title='%s' data-wordid='%s'" +
					" class='dict-entry' data-toggle='popover'>%s</span>",
					chunk, wordIds, chunk)
			*/
			// Regular HTML link
			fmt.Fprintf(&b, hyperlink(*entries[0], chunk))
		} else {
			b.WriteString(chunk)
		}
	}

	textContent := b.String()
	dateUpdated := time.Now().Format("2006-01-02")
	content := CorpusEntryContent{textContent, dateUpdated, collectionURL,
		collectionTitle, aFile}

	// Write to file
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	
	templFile := config.TemplateDir() + "/corpus-template.html"
	tmpl:= template.Must(template.New("corpus-template.html").ParseFiles(templFile))
	err = tmpl.Execute(w, content)
	if err != nil { panic(err) }
	w.Flush()
}

// Writes a document with markup for the array of tokens
// tokens: A list of tokens forming the document
// vocab: A list of word id's in the document
// filename: The file name to write to
func WriteDoc(tokens list.List, vocab map[string]int, filename,
		templateName, templateFile string) {
	if templateFile != `\N` {
		writeHTMLDoc(tokens, vocab, filename, templateName, templateFile)
		return
	}
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

// Writes a document with markup for the array of tokens
// tokens: A list of tokens forming the document
// vocab: A list of word id's in the document
// filename: The file name to write to
func writeHTMLDoc(tokens list.List, vocab map[string]int, filename,
		templateName, templateFile string) {
	var b bytes.Buffer

	// Iterate over text chunks
	for e := tokens.Front(); e != nil; e=e.Next() {
		chunk := e.Value.(string)
		//fmt.Printf("WriteDoc: Word %s\n", word)
		if entries, ok := dictionary.GetWord(chunk); ok {
			// Popover
			/*
			wordIds := ""
			for _, ws := range entries {
				if wordIds == "" {
					wordIds = fmt.Sprintf("%d", ws.Id)
				} else {
					wordIds = fmt.Sprintf("%s,%d", wordIds, ws.Id)
				}
			}
			fmt.Fprintf(&b, "<span title='%s' data-wordid='%s'" +
					" class='dict-entry' data-toggle='popover'>%s</span>",
					chunk, wordIds, chunk)
					*/
			// Regular HTML link
			mouseover := fmt.Sprintf("%s | %s", entries[0].Pinyin,
				entries[0].English)
			link := fmt.Sprintf("/words/%d.html", entries[0].HeadwordId)
			fmt.Fprintf(&b, "<a title='%s' href='%s'>%s</a>", mouseover, link,
				chunk)
		} else {
			fmt.Fprintf(&b, chunk)
		}
	}
	vocabJSON := WriteVocab("", 0, vocab)
	dateUpdated := time.Now().Format("2006-01-02")
	content := HTMLContent{b.String(), vocabJSON, dateUpdated}

	// Prepare template
	tmpl := template.Must(template.New(templateName).ParseFiles(templateFile))
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	w := bufio.NewWriter(f)
	err = tmpl.Execute(w, content)
	if err != nil { panic(err) }
	w.Flush()
	f.Close()

}

// Writes word entries for headwords
func WriteHwFiles() {
	fmt.Printf("WriteHwFiles: Begin +++++++++++\n")
	hwArray := dictionary.GetHeadwords()
	usageMap, _, _ := GetWordFrequencies()
	dateUpdated := time.Now().Format("2006-01-02")

	// Prepare template
	templFile := config.ProjectHome() + "/html/templates/headword-template.html"
	tmpl := template.Must(template.New("headword-template.html").ParseFiles(templFile))

	for _, hw := range hwArray {

		// Combine usage arrays for both simplified and traditional characters
		usageArrPtr, ok := usageMap[hw.Simplified]
		if !ok {
			usageArrPtr, ok = usageMap[hw.Traditional]
			if !ok {
				//log.Printf("WriteHwFiles: no usage for %s", hw.Simplified)
				usageArrPtr = &[]WordUsage{}
			}
		} else {
			usageArrTradPtr, ok := usageMap[hw.Traditional]
			if ok {
				usageArr := *usageArrPtr
				usageArrTrad := *usageArrTradPtr
				for j, _ := range usageArrTrad {
					usageArr = append(usageArr, usageArrTrad[j])
				}
				usageArrPtr = &usageArr
			}
		}

		// Decorate useage text
		hlUsageArr := []WordUsage{}
		for _, wu := range *usageArrPtr {
			hlText := decodeUsageExample(wu.Example, hw)
			hlWU := WordUsage{
				Freq: wu.Freq,
				RelFreq: wu.RelFreq,
				Word: wu.Word,
				Example: hlText,
				File: wu.File,
				EntryTitle: wu.EntryTitle,
				ColTitle: wu.ColTitle,
			}
			hlUsageArr = append(hlUsageArr, hlWU)
		}

		dictEntry := DictEntry{hw, hlUsageArr, dateUpdated}
		filename := fmt.Sprintf("%s%s%d%s", config.ProjectHome(), "/web/words/",
			hw.Id, ".html")
		f, err := os.Create(filename)
		if err != nil {
			log.Printf("WriteHwFiles: Error creating file for hw.Id %d, " +
				"Simplified %s", hw.Id, hw.Simplified)
			log.Fatal(err)
		}
		w := bufio.NewWriter(f)
		err = tmpl.Execute(w, dictEntry)
		if err != nil {
			log.Printf("WriteHwFiles: error executing template for hw.Id: %d," +
				" filename: %s, Simplified: %s", hw.Id, filename, hw.Simplified)
			log.Fatal(err)
		}
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

