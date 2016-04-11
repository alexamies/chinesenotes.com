/*
Library for Chinese vocabulary analysis
*/
package analysis

import (
	"bufio"
	"bytes"
	"cnreader/alignment"
	"cnreader/config"
	"cnreader/corpus"
	"cnreader/dictionary"
	"cnreader/ngram"
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
const MAX_WF_OUTPUT = 500

// Maximum number of unknwon characters to output to the generated
// HTML file
const MAX_UNKOWN_OUTPUT = 50

// Word frequency output file
const UNIGRAM_FILE = "unigram.txt"

// Max usage elements for a word
const MAX_USAGE = 100

// Max number of occurrences of same title in a list of word usages
const MAX_TITLE = 10

// Holds vocabulary analysis for a corpus text
type AnalysisResults struct {
	Title string
	WC, UniqueWords int
	Cognates []alignment.CorpEntryCognates
	WordFrequencies []WFResult
	LexicalWordFreq []WFResult
	BigramFreqSorted []ngram.BigramFreq
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
	Contains []dictionary.HeadwordDef
	Collocations []ngram.BigramFreq
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
	tokens, _ := ParseText(usageText, "", corpus.NewCorpusEntry())
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

// Compute word frequencies, collocations, and usage for the entire corpus
func GetWordFrequencies() (map[string]*[]WordUsage,
		map[*CorpusWord]CorpusWordFreq, map[string]int, ngram.CollocationMap) {

	// Overall word frequencies per corpus
	collocations := ngram.CollocationMap{}
	usageMap := map[string]*[]WordUsage{}
	wcTotal := map[string]int{}
	wfTotal := map[*CorpusWord]CorpusWordFreq{}
	corpusDir := config.ProjectHome() + "/corpus/"
	corpusDataDir := config.ProjectHome() + "/data/corpus/"

	collectionEntries := corpus.Collections()
	for _, col := range collectionEntries {
		colFile := corpusDataDir + col.CollectionFile
		//log.Printf("GetWordFrequencies: input file: %s\n", colFile)
		corpusEntries := corpus.CorpusEntries(colFile)
		for _, entry := range corpusEntries {
			src := corpusDir + entry.RawFile
			text := ReadText(src)
			_, results := ParseText(text, col.Title, &entry)
			wcTotal[col.Corpus] += results.WC

			// Process collocations
			collocations.MergeCollocationMap(results.Collocations)

			// Find word usage for the given word
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

	usageMap = sampleUsage(usageMap)

	// Print out totals for each corpus
	for corpus, count := range wcTotal {
		log.Printf("WordFrequencies: Total word count for corpus %s: %d\n",
			corpus, count)
	}
	return usageMap, wfTotal, wcTotal, collocations
}

// Constructs a hyperlink for a headword, including Pinyin and English in the
// title attribute for the link mouseover
func hyperlink(entry dictionary.WordSenseEntry, text string) string {
	mouseover := fmt.Sprintf("%s | %s", entry.Pinyin, entry.English)
	link := fmt.Sprintf("/words/%d.html", entry.HeadwordId)
	classTxt := ""
	if entry.IsProperNoun() {
		classTxt = " class='propernoun'"
	}
	return fmt.Sprintf("<a title='%s' %s href='%s'>%s</a>", mouseover, classTxt,
		link, text)
}

// Parses a Chinese text into words
// Parameters:
// text: the string to parse
// ColTitle: Optional parameter used for tracing collocation usage
// document: Optional parameter used for tracing collocation usage
// Returns:
// tokens: the tokens for the parsed text
// results: vocabulary analysis results
func ParseText(text string, colTitle string, document *corpus.CorpusEntry) (
		tokens list.List, results CollectionAResults) {
	vocab := map[string]int{}
	bigramMap := ngram.BigramFreqMap{}
	collocations := ngram.CollocationMap{}
	corpEntryCogs := alignment.NewCorpEntryCognates(*document)
	unknownChars := map[string]int{}
	usage := map[string]string{}
	wc := 0
	chunks := GetChunks(text)
	wdict := dictionary.GetWDict()
	lastHWPtr := new(dictionary.HeadwordDef)
	lastHW := *lastHWPtr
	//fmt.Printf("ParseText: For text %s got %d chunks\n", text, chunks.Len())
	for e := chunks.Front(); e != nil; e = e.Next() {
		chunk := e.Value.(string)
		//fmt.Printf("ParseText: chunk %s\n", chunk)
		characters := strings.Split(chunk, "")
		if !dictionary.IsCJKChar(characters[0]) {
			tokens.PushBack(chunk)
			lastHWPtr = new(dictionary.HeadwordDef)
			lastHW = *lastHWPtr
			continue
		}
		for i := 0; i < len(characters); i++ {
			for j := len(characters); j > 0; j-- {
				w := strings.Join(characters[i:j], "")
				//fmt.Printf("ParseText: i = %d, j = %d, w = %s\n", i, j, w)
				if wsArray, ok := wdict[w]; ok {
					//fmt.Printf("ParseText: found word %s, i = %d\n", w, i)
					tokens.PushBack(w)
					wc++
					vocab[w]++
					if _, ok := usage[w]; !ok {
						usage[w] = chunk
					}
					i = j - 1
					j = 0
					hw := dictionary.HeadwordDef{
						Id: wsArray[0].HeadwordId,
						Simplified: wsArray[0].Simplified,
						Traditional: wsArray[0].Traditional,
						Pinyin: []string{},
						WordSenses: []dictionary.WordSenseEntry{*wsArray[0]},
					}
					if lastHW.Id != 0 {
						bigram := ngram.Bigram{
							HeadwordDef1: lastHW,
							HeadwordDef2: hw,
							Example: chunk,
							ExFile: document.GlossFile,
							ExDocTitle: document.Title,
							ExColTitle: colTitle,
						}
						bigramMap.PutBigram(bigram)
						collocations.PutBigram(bigram.HeadwordDef1.Id, bigram)
						collocations.PutBigram(bigram.HeadwordDef2.Id, bigram)
					}
					lastHW = hw
					corpEntryCogs.AddCognate(wsArray[0])
				} else if (utf8.RuneCountInString(w) == 1) {
					//log.Printf("ParseText: found unknown character %s\n", w)
					unknownChars[w]++
					tokens.PushBack(w)
					break
				}
			}
		}
	}
	collectionCogs := []alignment.CorpEntryCognates{}
	collectionCogs = append(collectionCogs, corpEntryCogs)
	results = CollectionAResults {
		Vocab: vocab,
		Usage: usage,
		BigramFrequencies: bigramMap,
		Collocations: collocations,
		CollectionCogs: collectionCogs,
		WC: wc,
		UnknownChars: unknownChars,
	}
	return tokens, results
}

// Sample word usage for usability making sure that the list of word usage
// samples is not dominated by any one title and truncating at MAX_USAGE
// examples.
func sampleUsage(usageMap map[string]*[]WordUsage) map[string]*[]WordUsage {
	for word, usagePtr := range usageMap {
		sampleMap := map[string]int{}
		usage := *usagePtr
		usageCapped := new([]WordUsage)
		j := 0
		for _, wu := range usage {
			count, _ := sampleMap[wu.ColTitle]
			if count < MAX_TITLE && j < MAX_USAGE {
				*usageCapped = append(*usageCapped, wu)
				sampleMap[wu.ColTitle]++
				j++
			}
		}
		usageMap[word] = usageCapped
	}
	return usageMap
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
	usageMap, wfTotal, wcTotal, _ := GetWordFrequencies()
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

// Writes out an analysis of the entire corpus, including word frequencies
// and other data. The output file is called 'corpus-analysis.html' in the
// web/analysis directory.
// results: The results of corpus analysis
// Returns the name of the file written to
func writeAnalysisCorpus(results CollectionAResults) string {

	// Parse template and organize template parameters
	sortedWords := SortedFreq(results.Vocab)
	wfResults := results.GetWordFreq(sortedWords)
	lexicalWordFreq := results.GetLexicalWordFreq(sortedWords)
	sortedUnknownWords := SortedFreq(results.UnknownChars)

	// Bigrams, also truncated
	bFreq := ngram.SortedFreq(results.BigramFrequencies)
	maxBFOutput:= len(bFreq)
	if maxBFOutput > MAX_WF_OUTPUT {
		maxBFOutput = MAX_WF_OUTPUT
	}

	dateUpdated := time.Now().Format("2006-01-02")
	title := "Corpus Analysis"
	aResults := AnalysisResults{
		Title: title,
		WC: results.WC,
		Cognates: []alignment.CorpEntryCognates{},
		UniqueWords: len(results.Vocab),
		WordFrequencies: wfResults,
		LexicalWordFreq: lexicalWordFreq,
		BigramFreqSorted: bFreq[:maxBFOutput],
		UnkownnChars: sortedUnknownWords, 
		DateUpdated: dateUpdated, 
		MaxWFOutput: len(wfResults),
	}
	tmplFile := config.TemplateDir() + "/corpus-analysis-template.html"
	tmpl, err := template.New("corpus-analysis-template.html").ParseFiles(tmplFile)
	if err != nil { panic(err) }
	if tmpl == nil {
		log.Fatal("writeAnalysis: Template is nil", err)
	}
	basename := "corpus_analysis.html"
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

// Writes a document with vocabulary analysis of the text. The name of the
// output file will be source file with '-analysis' appended, placed in the
// web/analysis directory
// results: The results of vocabulary analysis
// collectionTitle: The title of the whole colleciton
// docTitle: The title of this specific document
// Returns the name of the file written to
func writeAnalysis(results CollectionAResults, srcFile, collectionTitle,
		docTitle string) string {

	// Parse template and organize template parameters
	sortedWords := SortedFreq(results.Vocab)
	wfResults := results.GetWordFreq(sortedWords)
	lexicalWordFreq := results.GetLexicalWordFreq(sortedWords)
	sortedUnknownWords := SortedFreq(results.UnknownChars)

	// Bigrams, also truncated
	bFreq := ngram.SortedFreq(results.BigramFrequencies)
	maxBFOutput:= len(bFreq)
	if maxBFOutput > MAX_WF_OUTPUT {
		maxBFOutput = MAX_WF_OUTPUT
	}

	dateUpdated := time.Now().Format("2006-01-02")
	title := "Content Analysis for " + collectionTitle + ", " + docTitle
	aResults := AnalysisResults{
		Title: title,
		WC: results.WC,
		Cognates: results.CollectionCogs,
		UniqueWords: len(results.Vocab),
		WordFrequencies: wfResults,
		LexicalWordFreq: lexicalWordFreq,
		BigramFreqSorted: bFreq[:maxBFOutput],
		UnkownnChars: sortedUnknownWords, 
		DateUpdated: dateUpdated, 
		MaxWFOutput: len(wfResults),
	}
	tmplFile := config.TemplateDir() + "/corpus-analysis-template.html"
	tmpl, err := template.New("corpus-analysis-template.html").ParseFiles(tmplFile)
	if err != nil { panic(err) }
	if tmpl == nil {
		log.Fatal("writeAnalysis: Template is nil", err)
	}

	// Write output
	i := strings.Index(srcFile, ".txt")
	if i <= 0 {
		i = strings.Index(srcFile, ".html")
		if i <= 0 {
			i = strings.Index(srcFile, ".csv")
			if i <= 0 {
				log.Fatal("writeAnalysis: Bad name for source file: ", srcFile)
			}
		}
	}
	basename := srcFile[:i] + "_analysis.html"
	filename := config.ProjectHome() + "/web/analysis/" + basename
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal("analysis.writeAnalysis", err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	err = tmpl.Execute(w, aResults)
	if err != nil { panic(err) }
	w.Flush()
	return basename
}

// Writes a corpus document collection to HTML, including all the entries
// contained in the collection
// collectionEntry: the CollectionEntry struct
func writeCollection(collectionEntry corpus.CollectionEntry) CollectionAResults {
	corpusEntries := corpus.CorpusEntries(config.CorpusDataDir() + "/" +
			collectionEntry.CollectionFile)
	aResults := NewCollectionAResults()
	for _, entry := range corpusEntries {
		src := config.CorpusDir() + "/" + entry.RawFile
		dest := config.WebDir() + "/" + entry.GlossFile
		//log.Printf("analysis.writeCollection: input file: %s, output file:
			//%s\n", src, dest)
		text := ReadText(src)
		tokens, results := ParseText(text, collectionEntry.Title, &entry)
		aFile := writeAnalysis(results, entry.RawFile, collectionEntry.Title,
			entry.Title)
		writeCorpusDoc(tokens, results.Vocab, dest, collectionEntry.GlossFile,
			collectionEntry.Title, aFile)
		aResults.AddResults(results)
	}
	aFile := writeAnalysis(aResults, collectionEntry.CollectionFile,
		collectionEntry.Title, "")
	corpus.WriteCollectionFile(collectionEntry.CollectionFile, aFile)
	log.Printf("analysis.writeCollection: completed: %s\n",
		collectionEntry.CollectionFile)
	return aResults
}

func WriteCorpusAll() {
	collections := corpus.Collections()
	aResults := NewCollectionAResults()
	for _, collectionEntry := range collections {
		results := writeCollection(collectionEntry)
		aResults.AddResults(results)
	}
	writeAnalysisCorpus(aResults)
}

// Writes a corpus document collection to HTML, including all the entries
// contained in the collection
// collectionFile: the name of the collection file
func WriteCorpusCol(collectionFile string) {
	collectionEntry, err := corpus.GetCollectionEntry(collectionFile)
	if err != nil {
		log.Fatalf("analysis.WriteCorpusCol: fatal error %v", err)
	}
	writeCollection(collectionEntry)
}

// Writes a corpus document with markup for the array of tokens
// tokens: A list of tokens forming the document
// vocab: A list of word id's in the document
// filename: The file name to write to
// HTML template to use
// collectionURL: the URL of the collection that the corpus text belongs to
// collectionTitle: The collection title that the corpus entry belongs to
// aFile: The vocabulary analysis file written to or empty string for none
func writeCorpusDoc(tokens list.List, vocab map[string]int, filename string,
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

// Writes dictionary headword entries
func WriteHwFiles() {
	log.Printf("analysis.WriteHwFiles: Begin +++++++++++\n")
	hwArray := dictionary.GetHeadwords()
	usageMap, _, _, collocations := GetWordFrequencies()
	dateUpdated := time.Now().Format("2006-01-02")

	// Prepare template
	templFile := config.ProjectHome() + "/html/templates/headword-template.html"
	tmpl := template.Must(template.New("headword-template.html").ParseFiles(templFile))

	i := 0
	for _, hw := range hwArray {

		if i % 1000 == 0 {
			log.Printf("analysis.WriteHwFiles: wrote %d words\n", i)
		}

		// Words that contain this word
		contains := dictionary.ContainsWord(hw.Simplified, hwArray)

		// Sorted array of collocations
		wordCollocations := collocations.SortedCollocations(hw.Id)

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

		dictEntry := DictEntry{
			Headword: hw,
			Contains: contains,
			Collocations: wordCollocations,
			UsageArr: hlUsageArr,
			DateUpdated: dateUpdated,
		}
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
		i++
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

