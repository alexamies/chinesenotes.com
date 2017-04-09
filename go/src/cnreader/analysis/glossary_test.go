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
	firstWord := glossary.Words[0].Pinyin[0]
	pinyinExpected := hw1.Pinyin[0]
	if pinyinExpected != firstWord {
		t.Error("analysis.TestMakeGlossary3: Expected pinyin ", pinyinExpected,
			", got", firstWord)
	}
}