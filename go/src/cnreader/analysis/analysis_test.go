/*
 * Unit tests for the analysis package
 */
package analysis

import (
	"cnreader/config"
	"cnreader/corpus"
	"cnreader/dictionary"
	"cnreader/index"
	"cnreader/library"
	"log"
	"strings"
	"testing"
	"time"
)

func TestDecodeUsageExample1(t *testing.T) {
	s1 := "海"
	s2 := "\\N"
	hw := dictionary.HeadwordDef{
		Id:          1,
		Simplified:  &s1,
		Traditional: &s2,
		Pinyin:      []string{"hǎi"},
		WordSenses:  &[]dictionary.WordSenseEntry{},
	}
	highlighted := decodeUsageExample("海", hw)
	expected := "<span class='usage-highlight'>海</span>"
	if highlighted != expected {
		t.Error("TestDecodeUsageExample1: Expected ", expected, ", got",
			highlighted)
	}
}

func TestDecodeUsageExample2(t *testing.T) {
	s1 := "海"
	s2 := "\\N"
	hw := dictionary.HeadwordDef{
		Id:          1,
		Simplified:  &s1,
		Traditional: &s2,
		Pinyin:      []string{"hǎi"},
		WordSenses:  &[]dictionary.WordSenseEntry{},
	}
	highlighted := decodeUsageExample("banana", hw)
	expected := "banana"
	if highlighted != expected {
		t.Error("TestDecodeUsageExample2: Expected ", expected, ", got",
			highlighted)
	}
}

func TestDecodeUsageExample3(t *testing.T) {
	s1 := "国"
	s2 := "國"
	hw := dictionary.HeadwordDef{
		Id:          1,
		Simplified:  &s1,
		Traditional: &s2,
		Pinyin:      []string{"guó"},
		WordSenses:  &[]dictionary.WordSenseEntry{},
	}
	highlighted := decodeUsageExample("國", hw)
	expected := "<span class='usage-highlight'>國</span>"
	if highlighted != expected {
		t.Error("TestDecodeUsageExample3: Expected ", expected, ", got",
			highlighted)
	}
}
func TestGetChunks1(t *testing.T) {
	files := []string{"../testdata/testwords.txt"}
	dictionary.ReadDict(files)
	chunks := GetChunks("中文")
	if chunks.Len() != 1 {
		t.Error("TestGetChunks1: Expected length of chunks 1, got ",
			chunks.Len())
	}
	chunk := chunks.Front().Value.(string)
	if chunk != "中文" {
		t.Error("TestGetChunks1: Expected first element of chunk 中文, got ",
			chunk)
	}
}

func TestGetChunks2(t *testing.T) {
	files := []string{"../testdata/testwords.txt"}
	dictionary.ReadDict(files)
	chunks := GetChunks("a中文")
	if chunks.Len() != 2 {
		t.Error("Expected length of chunks 2, got ", chunks.Len())
	}
	chunk := chunks.Front().Value.(string)
	if chunk != "a" {
		t.Error("Expected first element of chunk a, got ", chunk)
	}
}

func TestGetChunks3(t *testing.T) {
	files := []string{"../testdata/testwords.txt"}
	dictionary.ReadDict(files)
	chunks := GetChunks("a中文b")
	if chunks.Len() != 3 {
		t.Error("Expected length of chunks 3, got ", chunks.Len())
	}
	chunk := chunks.Front().Value.(string)
	if chunk != "a" {
		t.Error("Expected first element of chunk a, got ", chunk)
	}
}

// Simplified Chinese
func TestGetChunks4(t *testing.T) {
	files := []string{"../testdata/testwords.txt"}
	dictionary.ReadDict(files)
	chunks := GetChunks("简体中文")
	if chunks.Len() != 1 {
		t.Error("Simplified Chinese 简体中文: expected length of chunks 1, got ",
			chunks.Len())
	}
	chunk := chunks.Front().Value.(string)
	if chunk != "简体中文" {
		for e := chunks.Front(); e != nil; e = e.Next() {
			log.Printf("TestGetChunks4: chunk: %s\n", e.Value.(string))
		}
		t.Error("Expected first element of chunk 简体中文 to be 简体中文, got ",
			chunk)
	}
}

func TestReadText1(t *testing.T) {
	//log.Printf("TestReadText1: Begin ******** \n")
	text := ReadText("../testdata/sampletest.txt")
	expected := "繁體中文<br/>\n"
	//log.Printf("TestReadText1: Expected  '%s', got '%s'\n", expected, text)
	if text != expected {
		t.Error("Expected ", expected, ", got ", text)
	}
	//log.Printf("TestReadText1: End ******** \n")
}

func TestReadText2(t *testing.T) {
	//log.Printf("TestReadText2: Begin ******** \n")
	text := ReadText("../testdata/test.html")
	if !strings.Contains(text, "繁體中文") {
		t.Error("Expected to contain '繁體中文', got ", text)
	}
	//log.Printf("TestReadText2: End ******** \n")
}

func TestParseText1(t *testing.T) {
	//log.Printf("TestParseText1: Begin ******** \n")
	files := []string{"../testdata/testwords.txt"}
	dictionary.ReadDict(files)
	tokens, results := ParseText("繁體中文", "", corpus.NewCorpusEntry())
	if tokens.Len() != 2 {
		t.Error("Expected to get length 2, got ", tokens.Len())
		first := tokens.Front().Value.(string)
		if first != "繁體" {
			t.Error("Expected to get tokens.Front() 繁體, got ", first)
		}
	}
	if len(results.Vocab) != 2 {
		t.Error("Expected to get len(vocab) = 2, got ", len(results.Vocab))
	}
	if results.WC != 2 {
		t.Error("Expected to get wc = 2, got ", results.WC)
	}
	if results.CCount != 4 {
		t.Error("Expected to get wc = 2, got ", results.WC)
	}
	//log.Printf("TestParseText1: End ******** \n")
}

func TestParseText2(t *testing.T) {
	//log.Printf("TestParseText2: Begin ******** \n")
	files := []string{"../testdata/testwords.txt"}
	dictionary.ReadDict(files)
	tokens, results := ParseText("a繁體中文", "", corpus.NewCorpusEntry())
	if tokens.Len() != 3 {
		t.Error("Expected to get length 3, got ", tokens.Len())
	}
	if results.WC != 2 {
		t.Error("Expected to get wc 2, got ", results.WC)
	}
	first := tokens.Front().Value.(string)
	if first != "a" {
		t.Error("Expected to get tokens.Front() a, got ", first)
	}
	if len(results.Vocab) != 2 {
		t.Error("Expected to get len(vocab) = 2, got ", len(results.Vocab))
	}
	//log.Printf("TestParseText2: End ******** \n")
}

func TestParseText3(t *testing.T) {
	//log.Printf("TestParseText3: Begin +++++++++++\n")
	files := []string{"../testdata/testwords.txt"}
	dictionary.ReadDict(files)
	tokens, results := ParseText("前不见古人", "", corpus.NewCorpusEntry())
	if tokens.Len() != 3 {
		t.Error("Expected to get length 3, got ", tokens.Len())
		return
	}
	if results.WC != 3 {
		t.Error("Expected to get wc 3, got ", results.WC)
		return
	}
	expected := []string{"前", "不见", "古人"}
	i := 0
	for e := tokens.Front(); e != nil; e = e.Next() {
		word := e.Value.(string)
		if expected[i] != e.Value.(string) {
			t.Error("Failed to get expected word", expected[i], word, i)
		}
		i++
	}
	//log.Printf("TestParseText3: End +++++++++++\n")
}

func TestParseText4(t *testing.T) {
	dictionary.ReadDict(config.LUFileNames())
	text := ReadText("../testdata/test-trad.html")
	tokens, results := ParseText(text, "", corpus.NewCorpusEntry())
	if tokens.Len() != 48 {
		t.Error("Expected to get length 48, got ", tokens.Len())
	}
	if results.CCount != 49 {
		t.Error("Expected to get cc 49, got ", results.CCount)
		return
	}
	if len(results.Vocab) != 37 {
		t.Error("Expected to get Vocab 37, got ", len(results.Vocab), results.Vocab)
		return
	}
}

// Basic test with no data
func TestSampleUsage1(t *testing.T) {
	//log.Printf("TestSampleUsage1: Begin +++++++++++\n")
	usageMap := map[string]*[]WordUsage{}
	usageMap = sampleUsage(usageMap)
	l := len(usageMap)
	expected := 0
	if l != expected {
		t.Error("Expected to get length ", expected, ", got ", l)
	}
}

// Basic test with minimal data
func TestSampleUsage2(t *testing.T) {
	wu := WordUsage{
		Freq:       1,
		RelFreq:    0.01,
		Word:       "大",
		Example:    "大蛇",
		File:       "afile.txt",
		EntryTitle: "Scroll 1",
		ColTitle:   "A Big Snake",
	}
	wuArray := []WordUsage{wu}
	usageMap := map[string]*[]WordUsage{"大": &wuArray}
	usageMap = sampleUsage(usageMap)
	l := len(usageMap)
	expected := 1
	if l != expected {
		t.Error("Expected to get length ", expected, ", got ", l)
	}
}

// Basic test with more data
func TestSampleUsage3(t *testing.T) {
	log.Printf("TestSampleUsage3: Begin +++++++++++\n")
	wu1 := WordUsage{
		Freq:       1,
		RelFreq:    0.01,
		Word:       "蛇",
		Example:    "大蛇",
		File:       "afile.txt",
		EntryTitle: "Scroll 1",
		ColTitle:   "Some Snakes",
	}
	wu2 := WordUsage{
		Freq:       1,
		RelFreq:    0.01,
		Word:       "蛇",
		Example:    "小蛇",
		File:       "afile.txt",
		EntryTitle: "Scroll 2",
		ColTitle:   "Some Snakes",
	}
	wuArray := []WordUsage{wu1, wu2}
	usageMap := map[string]*[]WordUsage{"蛇": &wuArray}
	usageMap = sampleUsage(usageMap)
	l := len(*usageMap["蛇"])
	expected := 2
	if l != expected {
		t.Error("Expected to get length ", expected, ", got ", l)
	}
}

// Basic test with more data
func TestSampleUsage4(t *testing.T) {
	log.Printf("analysis.TestSampleUsage4: Begin +++++++++++\n")
	wu1 := WordUsage{
		Freq:       1,
		RelFreq:    0.01,
		Word:       "大",
		Example:    "大蛇",
		File:       "afile.txt",
		EntryTitle: "Scroll 1",
		ColTitle:   "Some Big Animals",
	}
	wu2 := WordUsage{
		Freq:       1,
		RelFreq:    0.01,
		Word:       "大",
		Example:    "大老虎",
		File:       "afile.txt",
		EntryTitle: "Scroll 2",
		ColTitle:   "Some Big Animals",
	}
	wu3 := WordUsage{
		Freq:       1,
		RelFreq:    0.01,
		Word:       "大",
		Example:    "大树",
		File:       "afile.txt",
		EntryTitle: "Scroll 1",
		ColTitle:   "Some Big Trees",
	}
	wuArray := []WordUsage{wu1, wu2, wu3}
	usageMap := map[string]*[]WordUsage{"大": &wuArray}
	usageMap = sampleUsage(usageMap)
	l := len(*usageMap["大"])
	expected := 3
	if l != expected {
		t.Error("Expected to get length ", expected, ", got ", l)
	}
	log.Printf("analysis.TestSampleUsage4: End +++++++++++\n")
}

func TestSortedFreq(t *testing.T) {
	dictionary.ReadDict(config.LUFileNames())
	text := ReadText("../testdata/test-trad.html")
	_, results := ParseText(text, "", corpus.NewCorpusEntry())
	sortedWords := index.SortedFreq(results.Vocab)
	expected := len(results.Vocab)
	got := len(sortedWords)
	if expected != got {
		t.Error("TestSortedFreq: Expected %d, got %d", expected, got)
		return
	}
}

func TestWriteAnalysis(t *testing.T) {
	log.Printf("analysis.TestWriteAnalysis: Begin +++++++++++\n")
	term := "繁"
	_, results := ParseText(term, "", corpus.NewCorpusEntry())
	srcFile := "test.txt"
	vocab := map[string]int{
		term: 1,
	}
	df := index.NewDocumentFrequency()
	df.AddVocabulary(vocab)
	df.WriteToFile()
	index.ReadDocumentFrequency()
	writeAnalysis(results, srcFile, "Test Collection", "Test Doc")
	log.Printf("analysis.TestWriteAnalysis: End +++++++++++\n")
}

/*
func TestWriteCorpusAll(t *testing.T) {
	log.Printf("analysis.TestWriteCorpusAll: Begin +++++++++++\n")
	WriteCorpusAll()
	log.Printf("analysis.TestWriteCorpusAll: End +++++++++++\n")
}
*/

/*
func TestWriteCorpusCol(t *testing.T) {
	log.Printf("analysis.TestWriteCorpusCol: Begin +++++++++++\n")
	WriteCorpusCol("lunyu.csv")
	log.Printf("analysis.TestWriteCorpusCol: End +++++++++++\n")
}
*/

func TestWriteCorpusDoc1(t *testing.T) {
	log.Printf("analysis.TestWriteCorpusDoc1: Begin +++++++++++\n")
	tokens, results := ParseText("繁", "", corpus.NewCorpusEntry())
	outfile := "../testoutput/output.html"
	writeCorpusDoc(tokens, results.Vocab, outfile, "", "", "", "")
	log.Printf("analysis.TestWriteCorpusDoc1: End +++++++++++\n")
}

func TestWriteDoc1(t *testing.T) {
	log.Printf("analysis.TestWriteDoc1: Begin +++++++++++\n")
	tokens, results := ParseText("繁", "", corpus.NewCorpusEntry())
	outfile := "../testoutput/output.html"
	WriteDoc(tokens, results.Vocab, outfile, `\N`, `\N`, true, "")
	log.Printf("analysis.TestWriteDoc1: End +++++++++++\n")
}

func TestWriteDoc2(t *testing.T) {
	files := []string{"../testdata/testwords.txt"}
	dictionary.ReadDict(files)
	text := ReadText("../testdata/test.html")
	tokens, results := ParseText(text, "", corpus.NewCorpusEntry())
	if tokens.Len() != 4 {
		t.Error("Expected to get length 4, got ", tokens.Len())
	}
	outfile := "../testoutput/test-gloss.html"
	WriteDoc(tokens, results.Vocab, outfile, `\N`, `\N`, true, "")
}

func TestWriteDoc3(t *testing.T) {
	files := []string{"../testdata/testwords.txt"}
	dictionary.ReadDict(files)
	text := ReadText("../testdata/test-simplified.html")
	tokens, results := ParseText(text, "", corpus.NewCorpusEntry())
	if tokens.Len() != 6 {
		t.Error("Expected to get length 6, got ", tokens.Len())
	}
	outfile := "../testoutput/test-simplified-gloss.html"
	WriteDoc(tokens, results.Vocab, outfile, `\N`, `\N`, true, "")
}

func TestWriteDoc4(t *testing.T) {
	files := []string{"../testdata/testwords.txt"}
	dictionary.ReadDict(files)
	text := ReadText("../testdata/test-simplified2.html")
	tokens, results := ParseText(text, "", corpus.NewCorpusEntry())
	l := 3
	if tokens.Len() != l {
		for e := tokens.Front(); e != nil; e = e.Next() {
			log.Println("analysis.TestWriteDoc4", e.Value.(string))
		}
		t.Error("TestWriteDoc4: Expected to get length ", l, ", got ",
			tokens.Len())
	}
	outfile := "../testoutput/test-simplified-gloss2.html"
	WriteDoc(tokens, results.Vocab, outfile, `\N`, `\N`, true, "")
}

func TestWriteLibraryFiles0(t *testing.T) {
	emptyLibLoader := library.EmptyLibraryLoader{"Empty"}
	dateUpdated := time.Now().Format("2006-01-02")
	lib := library.Library{
		Title: "Library",
		Summary: "Top level collection in the Library",
		DateUpdated: dateUpdated,
		TargetStatus: "public",
		Loader: emptyLibLoader,
	}
	WriteLibraryFiles(lib)
}

func TestWriteLibraryFiles1(t *testing.T) {
	mockLoader := library.MockLibraryLoader{"Mock"}
	dateUpdated := time.Now().Format("2006-01-02")
	lib := library.Library{
		Title: "Library",
		Summary: "Top level collection in the Library",
		DateUpdated: dateUpdated,
		TargetStatus: "public",
		Loader: mockLoader,
	}
	WriteLibraryFiles(lib)
}