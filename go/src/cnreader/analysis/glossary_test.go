// Test sorting of glossary
package analysis

import (
	"cnreader/dictionary"
	"log"
	"testing"
)

// make example data
func makeHW0() dictionary.HeadwordDef {
	simp := "国"
	trad := "國"
	pinyin := "guó"
	english := "a country / a state / a kingdom"
	topic_cn := "古文"
	topic_en := "Classical Chinese"
	ws0 := dictionary.WordSenseEntry{
		Id: 1075,
		HeadwordId: 1075,
		Simplified: simp,
		Traditional: trad,
		Pinyin: pinyin,
		English: english,
		Grammar: "noun",
		Concept_cn: "\\N",
		Concept_en: "\\N",
		Topic_cn: topic_cn,
		Topic_en: topic_en,
		Parent_cn: "",
		Parent_en: "",
		Image: "",
		Mp3: "",
		Notes: "",
	}
	wsArray := []dictionary.WordSenseEntry{ws0}
	return dictionary.HeadwordDef{
		Id: 1,
		Simplified: &simp,
		Traditional: &trad,
		Pinyin: []string{pinyin},
		WordSenses: &wsArray,
	}
}

// make example data
func makeHW1() dictionary.HeadwordDef {
	simp := "严净"
	trad := "嚴淨"
	pinyin := "yán jìng"
	english := "majestic and pure"
	topic_cn := "佛教"
	topic_en := "Buddhism"
	ws0 := dictionary.WordSenseEntry{
		Id: 62267,
		HeadwordId: 62267,
		Simplified: simp,
		Traditional: trad,
		Pinyin: pinyin,
		English: english,
		Grammar: "set phrase",
		Concept_cn: "\\N",
		Concept_en: "\\N",
		Topic_cn: topic_cn,
		Topic_en: topic_en,
		Parent_cn: "",
		Parent_en: "",
		Image: "",
		Mp3: "",
		Notes: "",
	}
	wsArray := []dictionary.WordSenseEntry{ws0}
	return dictionary.HeadwordDef{
		Id: 1,
		Simplified: &simp,
		Traditional: &trad,
		Pinyin: []string{pinyin},
		WordSenses: &wsArray,
	}
}

// make example data
func makeHW2() dictionary.HeadwordDef {
	simp := "缘缘"
	trad := "緣緣"
	pinyin := "yuányuán"
	english := "observed object condition"
	topic_cn := "佛教"
	topic_en := "Buddhism"
	ws0 := dictionary.WordSenseEntry{
		Id: 62252,
		HeadwordId: 62252,
		Simplified: simp,
		Traditional: trad,
		Pinyin: pinyin,
		English: english,
		Grammar: "set phrase",
		Concept_cn: "\\N",
		Concept_en: "\\N",
		Topic_cn: topic_cn,
		Topic_en: topic_en,
		Parent_cn: "",
		Parent_en: "",
		Image: "",
		Mp3: "",
		Notes: "",
	}
	wsArray := []dictionary.WordSenseEntry{ws0}
	return dictionary.HeadwordDef{
		Id: 1,
		Simplified: &simp,
		Traditional: &trad,
		Pinyin: []string{pinyin},
		WordSenses: &wsArray,
	}
}

// make example data
func makeHW3() dictionary.HeadwordDef {
	simp := "禅"
	trad := "禪"
	pinyin := "chán"
	english := "meditative concentration"
	topic_cn := "佛教"
	topic_en := "Buddhism"
	ws0 := dictionary.WordSenseEntry{
		Id: 3182,
		HeadwordId: 3182,
		Simplified: simp,
		Traditional: trad,
		Pinyin: pinyin,
		English: english,
		Grammar: "noun",
		Concept_cn: "\\N",
		Concept_en: "\\N",
		Topic_cn: topic_cn,
		Topic_en: topic_en,
		Parent_cn: "",
		Parent_en: "",
		Image: "",
		Mp3: "",
		Notes: "",
	}
	wsArray := []dictionary.WordSenseEntry{ws0}
	return dictionary.HeadwordDef{
		Id: 1,
		Simplified: &simp,
		Traditional: &trad,
		Pinyin: []string{pinyin},
		WordSenses: &wsArray,
	}
}

// Trivial test of MakeGlossary function
func TestMakeGlossary0(t *testing.T) {
	log.Printf("analysis.TestMakeGlossary0: Begin ******** \n")
	headwords := []dictionary.HeadwordDef{}
	MakeGlossary("test_label", headwords)
}

// Easy test of MakeGlossary function
func TestMakeGlossary1(t *testing.T) {
	hw := makeHW0()
	headwords := []dictionary.HeadwordDef{hw}
	glossary := MakeGlossary("test_label", headwords)
	len := len(glossary.Words)
	expected := 0
	if expected != len {
		t.Error("analysis.TestMakeGlossary2: Expected ", expected, ", got",
			len)
	}
}

// Happy path test of MakeGlossary function
func TestMakeGlossary2(t *testing.T) {
	hw0 := makeHW0()
	hw1 := makeHW1()
	headwords := []dictionary.HeadwordDef{hw0, hw1}
	glossary := MakeGlossary("Buddhism", headwords)
	len := len(glossary.Words)
	expected := 1
	if expected != len {
		t.Error("analysis.TestMakeGlossary2: Expected ", expected, ", got",
			len)
	}
}

// Test sorting in MakeGlossary method
func TestMakeGlossary3(t *testing.T) {
	hw0 := makeHW0()
	hw1 := makeHW1()
	hw2 := makeHW2()
	headwords := []dictionary.HeadwordDef{hw0, hw2, hw1}
	glossary := MakeGlossary("Buddhism", headwords)
	len := len(glossary.Words)
	expected := 2
	if expected != len {
		t.Error("TestMakeGlossary3: Expected ", expected, ", got",
			len)
	}
	//fmt.Println("analysis.TestMakeGlossary3, words: ", glossary.Words)
	firstWord := glossary.Words[0].Pinyin[0]
	pinyinExpected := hw1.Pinyin[0]
	if pinyinExpected != firstWord {
		t.Error("analysis.TestMakeGlossary3: Expected first ", pinyinExpected,
			", got", firstWord)
	}
}

// Test sorting in MakeGlossary method
func TestMakeGlossary4(t *testing.T) {
	hw0 := makeHW0()
	hw1 := makeHW1()
	hw2 := makeHW2()
	hw3 := makeHW3()
	headwords := []dictionary.HeadwordDef{hw0, hw2, hw1, hw3}
	glossary := MakeGlossary("Buddhism", headwords)
	len := len(glossary.Words)
	expected := 3
	if expected != len {
		t.Error("TestMakeGlossary4: Expected ", expected, ", got",
			len)
	}
	//fmt.Println("analysis.TestMakeGlossary4, words: ", glossary.Words)
	firstWord := glossary.Words[0].Pinyin[0]
	firstExpected := hw3.Pinyin[0]
	secondExpected := hw1.Pinyin[0]
	//result := firstExpected < secondExpected
	//fmt.Printf("analysis.TestMakeGlossary4, %s < %s = %v\n", firstExpected,
	//	secondExpected, result)
	if firstExpected != firstWord {
		t.Error("analysis.TestMakeGlossary3: Expected first ", firstExpected,
			", got", firstWord)
	}
	secondWord := glossary.Words[1].Pinyin[0]
	if secondExpected != secondWord {
		t.Error("analysis.TestMakeGlossary3: Expected second ", secondExpected,
			", got", secondWord)
	}
}