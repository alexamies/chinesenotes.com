/*
 * Unit tests for the fulltext package
 */
package fulltext

import (
 	"fmt"
	"testing"
)

// Trival test
func TestGetMatches0(t *testing.T) {
	fmt.Printf("fulltext.TestGetMatches0: Begin unit test\n")
	queryTerms := []string{}
	fileNames := []string{}
	dm := GetMatches(fileNames, queryTerms)
	fmt.Printf("fulltext.TestGetMatches0: match: %v\n", dm)
}

// Trival test
func TestGetMatches1(t *testing.T) {
	fmt.Printf("fulltext.TestGetMatches1: Begin unit test\n")
	queryTerms := []string{"曰風"}
	fn := "shijing/shijing001.txt"
	fileNames := []string{fn}
	docMatches := GetMatches(fileNames, queryTerms)
	if len(docMatches) == 0 {
		t.Errorf("TestGetMatches1: docMatches empty\n")
		return
	}
	snippet := docMatches[fn].MT.Snippet
	if len(snippet) == 0 {
		t.Errorf("TestGetMatches1: snippet empty\n")
		return
	}
	fmt.Printf("fulltext.TestGetMatches1: match: %v\n", docMatches)
}

// Trival test
func TestGetMatches2(t *testing.T) {
	fmt.Printf("fulltext.TestGetMatches1: Begin unit test\n")
	queryTerms := []string{"曰風"}
	fn0 := "shijing/shijing001.txt"
	fn1 := "shijing/shijing002.txt"
	fileNames := []string{fn0, fn1}
	docMatches := GetMatches(fileNames, queryTerms)
	if len(docMatches) == 0 {
		t.Errorf("TestGetMatches2: docMatches empty\n")
		return
	}
	snippet := docMatches[fn0].MT.Snippet
	if len(snippet) == 0 {
		t.Errorf("TestGetMatches2: snippet empty\n")
		return
	}
	fmt.Printf("fulltext.TestGetMatches2: match: %v\n", docMatches)
}
