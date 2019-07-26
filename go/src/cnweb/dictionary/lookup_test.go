// Unit tests for lookup package
package dictionary

import (
	"log"
	"testing"
)

// Query expecting empty list
func TestAddWordSense2Map(t *testing.T) {
	wmap := map[string]Word{}
	ws := WordSense{
		Id: 1,
		HeadwordId: 1,
		Simplified: "我",
		Traditional: "",
		Pinyin: "wǒ",
		English: "me",
		Notes: "No notes",
	}
	addWordSense2Map(wmap, ws)
	if len(wmap) != 1 {
		t.Error("TestAddWordSense2Map: unexpected length, ", len(wmap))
	}
}

// Test trivial query with empty query, expect error
func TestLookupSubstrEmpty(t *testing.T) {
	log.Printf("TestLookupSubstr: Begin unit tests\n")
	_, err := LookupSubstr("", "")
	if err == nil {
		t.Error("TestLookupSubstrEmpty: expected error")
	}
}

// Query expecting empty list
func TestLookupEmptyResult(t *testing.T) {
	results, err := LookupSubstr("我還不知道", "")
	if err != nil {
		t.Error("TestLookupEmptyResult: unexpected error, ", err)
		return
	}
	if len(results.Words) != 0 {
		t.Error("TestLookupEmptyResult: unexpected result length, ",
			      len(results.Words))
	}
}

// Query expecting empty list
func TestLookupOneResult(t *testing.T) {
	results, err := LookupSubstr("男扮", "Idiom")
	if err != nil {
		log.Print("TestLookupOneResult: unexpected error, ", err)
		return
	}
	if len(results.Words) != 1 {
		log.Print("TestLookupOneResult: unexpected result length, ",
			        len(results.Words))
		return
	}
	log.Print("TestLookupOneResult: got expected result length of 1")
}

// Query expecting empty list
func TestLookupNanResult(t *testing.T) {
	results, err := LookupSubstr("男", "Idiom")
	if err != nil {
		log.Print("TestLookupNanResult: unexpected error, ", err)
		return
	}
	if len(results.Words) != 1 {
		log.Print("TestLookupNanResult: unexpected result length, ",
			        len(results.Words))
		return
	}
	log.Print("TestLookupNanResult: got expected result length of 1")
}
