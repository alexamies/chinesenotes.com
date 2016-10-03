package dictionary

import (
	"cnreader/config"
	"fmt"
	"testing"
)

// Trival test to make sure that method does not explode
func TestContainsWord0(t *testing.T) {
	//fmt.Printf("TestContainsWord0: Begin unit tests\n")
	c := "hello"
	ReadDict(config.LUFileNames())
	headwords := GetHeadwords()
	contains := ContainsWord(c, headwords)
	result := len(contains)
	expected := 0
	if result != expected {
		t.Error("Expected ", expected, ", got ", result)
	}
}

// Basic test to make sure that method does something correct
func TestContainsWord1(t *testing.T) {
	fmt.Printf("TestContainsWord1: Begin unit tests\n")
	ReadDict(config.LUFileNames())
	c := "中"
	headwords := GetHeadwords()
	contains := ContainsWord(c, headwords)
	result := len(contains)
	notexpected := 0
	if result <= notexpected {
		t.Error("not expected ", notexpected, ", got ", result)
	}
}

func TestGetHwMap(t *testing.T) {
	GetHeadwords()
	GetHwMap()
}

// Both traditional and simplified
func TestIsCJKChar1(t *testing.T) {
	//fmt.Printf("TestIsCJKChar1: Begin unit tests\n")
	c := "中"
	result := IsCJKChar(c)
	if !result {
		fmt.Printf("TestIsCJKChar1: Testing Chinese %s\n", c)
		t.Error("Expected true, got ", result)
	}
}

// Non-Chinese
func TestIsCJKChar2(t *testing.T) {
	c := "a"
	result := IsCJKChar(c)
	if result {
		fmt.Printf("TestIsCJKChar2: Testing Chinese %s\n", c)
		t.Error("Expected false, got ", result)
	}
}

// Simplified Chinese
func TestIsCJKChar3(t *testing.T) {
	c := "简"
	result := IsCJKChar(c)
	if !result {
		fmt.Printf("TestIsCJKChar3: Testing Chinese %s\n", c)
		t.Error("Expected true, got ", result)
	}
}

// Both traditional and simplified
func TestIsCJKChar4(t *testing.T) {
	c := "古"
	result := IsCJKChar(c)
	if !result {
		fmt.Printf("TestIsCJKChar4: Testing Chinese %s\n", c)
		t.Error("Expected true, got ", result)
	}
}

// Test for punctuation
func TestIsCJKChar5(t *testing.T) {
	c := "。"
	result := IsCJKChar(c)
	if result {
		fmt.Printf("TestIsCJKChar5: Testing Chinese %s\n", c)
		t.Error("Expected false, got ", result)
	}
}

// Test for punctuation
func TestIsCJKChar6(t *testing.T) {
	c := "，"
	result := IsCJKChar(c)
	if result {
		fmt.Printf("TestIsCJKChar6: Testing Chinese %s\n", c)
		t.Error("Expected false, got ", result)
	}
}

// Test for punctuation
func TestIsCJKChar7(t *testing.T) {
	c := "-"
	result := IsCJKChar(c)
	if result {
		fmt.Printf("TestIsCJKChar7: Testing Chinese %s\n", c)
		t.Error("Expected false, got ", result)
	}
}

// Tests whether the word is a function word
func TestIsFunctionWord(t *testing.T) {
	ws1 := WordSenseEntry{
		Id: 1,
		Simplified: "蓝", 
		Traditional: "藍",
		Pinyin: "lán",
		Grammar: "adjective",
	}
	result := ws1.IsFunctionWord()
	expected := false
	if result != expected {
		t.Error("TestIsFunctionWord: Expected ", expected, ", got ", result)
	}
}

// Tests whether the word is a function word
func TestIsNumericExpression1(t *testing.T) {
	ws1 := WordSenseEntry{
		Id: 1,
		Simplified: "三", 
		Traditional: "\\N",
		Pinyin: "sān",
		Grammar: "number",
	}
	result := ws1.IsNumericExpression()
	expected := true
	if result != expected {
		t.Error("TestIsNumericExpression1: Expected ", expected, ", got ", result)
	}
}

// Tests whether the word is a function word
func TestIsNumericExpression2(t *testing.T) {
	ws1 := WordSenseEntry{
		Id: 1,
		Simplified: "三个人", 
		Traditional: "\\N",
		Pinyin: "sān gè rén",
		Grammar: "phrase",
	}
	result := ws1.IsNumericExpression()
	expected := true
	if result != expected {
		t.Error("TestIsNumericExpression2: Expected ", expected, ", got ", result)
	}
}

// Tests whether the word is a proper noun
func TestIsProperNoun1(t *testing.T) {
	ws1 := WordSenseEntry{
		Id: 1,
		Simplified: "蓝", 
		Traditional: "藍",
		Pinyin: "lán",
		Grammar: "adjective",
	}
	result := ws1.IsProperNoun()
	expected := false
	if result != expected {
		t.Error("TestIsProperNoun1: Expected ", expected, ", got ", result)
	}
}

// Tests whether a word with multiple senses is a proper noun
func TestIsProperNoun2(t *testing.T) {
	ws1 := WordSenseEntry{
		Id: 1,
		Simplified: "如来", 
		Traditional: "如來",
		Pinyin: "Rúlái",
		Grammar: "proper noun",
	}
	result := ws1.IsProperNoun()
	expected := true
	if result != expected {
		fmt.Printf("TestIsProperNoun2: ws1.Simplified: %v\n", ws1.Simplified)
		t.Error("TestIsProperNoun2: Expected ", expected, ", got ", result)
	}
}

func TestWriteHeadwords(t *testing.T) {
	fmt.Printf("TestWriteHeadwords: Begin +++++++++++\n")
	ReadDict(config.LUFileNames())
	WriteHeadwords()
}

func TestReadDict1(t *testing.T) {
	fileNames := []string{"../testdata/testwords.txt"}
	ReadDict(fileNames)
	ws, ok := GetWordSense("中文")
	if !ok {
		t.Error("Expected true, got ", ok)
	}
	if ws.Id != 1 {
		t.Error("Expected 1, got ", ws.Id)
	}
	if ws.Simplified != "中文" {
		t.Error("Expected 中文, got ", ws.Simplified)
	}
	if ws.Traditional != "\\N" {
		t.Error("Expected \\N, got ", ws.Traditional)
	}
	if ws.Pinyin != "Zhōngwén" {
		t.Error("Expected Zhōngwén, got ", ws.Pinyin)
	}
	if ws.English != "Chinese language" {
		t.Error("Expected Chinese language, got ", ws.English)
	}
	if ws.Grammar != "noun" {
		t.Error("Expected noun, got ", ws.Grammar)
	}
	if ws.Concept_cn != "\\N" {
		t.Error("Expected \\N, got ", ws.Concept_cn)
	}
	if ws.Concept_en != "\\N" {
		t.Error("Expected \\N, got ", ws.Concept_en)
	}
	if ws.Topic_cn != "语言" {
		t.Error("Expected 语言, got ", ws.Topic_cn)
	}
	if ws.Topic_en != "Language" {
		t.Error("Expected Language, got ", ws.Topic_en)
	}
	if ws.Parent_cn != "\\N" {
		t.Error("Expected \\N, got ", ws.Parent_cn)
	}
	if ws.Parent_en != "\\N" {
		t.Error("Expected \\N, got ", ws.Parent_en)
	}
	if ws.Image != "\\N" {
		t.Error("Expected \\N, got ", ws.Image)
	}
	if ws.Mp3 != "zhong1wen2.mp3" {
		t.Error("Expected zhong1wen2.mp3, got ", ws.Mp3)
	}
	if ws.Notes != "May refer to" {
		t.Error("Expected 'May refer to', got ", ws.Notes)
	}
}

func TestGetWord(t *testing.T) {
	fileNames := []string{"../testdata/testwords.txt"}
	ReadDict(fileNames)
	word, ok := GetWord("中")
	if !ok {
		t.Error("Expected true, got ", ok)
	}
	if len(word) != 2 {
		t.Error("Expected len(word) = 2, got ", len(word))
	}
}
