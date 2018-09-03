/*
 * Unit tests for the fulltext package
 */
package fulltext

import (
 	"fmt"
	"testing"
)

func TestgetMatch0(t *testing.T) {
	fmt.Printf("fulltext.TestgetMatch0: Begin unit test\n")
	txt := "厚人倫，美教化，移風俗。故詩有六義焉：一曰風，二曰賦，三曰比，四曰興，五曰雅，六曰頌。"
	queryTerms := []string{"曰", "風"}
	mt := getMatch(txt, queryTerms)
	if mt.Snippet == "" {
		t.Errorf("TestgetMatch0: snippet empty\n")
	}
	expectLM := "曰風"
	if mt.LongestMatch != expectLM {
		t.Errorf("TestgetMatch0: expect %s. got %s\n", expectLM, mt.LongestMatch)
	}
	expectEM := true
	if mt.ExactMatch != expectEM {
		t.Errorf("TestgetMatch0: expect %v. got %v\n", expectEM, mt.ExactMatch)
	}
}

func TestgetMatch1(t *testing.T) {
	fmt.Printf("fulltext.TestgetMatch0: Begin unit test\n")
	txt := "厚人倫，美教化，移風俗。故詩有六義焉：一曰風，二曰賦，三曰比，四曰興，五曰雅，六曰頌。"
	queryTerms := []string{"一", "曰風"}
	mt := getMatch(txt, queryTerms)
	if mt.Snippet == "" {
		t.Errorf("TestgetMatch1: snippet empty\n")
	}
	expectLM := "一曰風"
	if mt.LongestMatch != expectLM {
		t.Errorf("TestgetMatch1: expect %s. got %s\n", expectLM, mt.LongestMatch)
	}
	expectEM := true
	if mt.ExactMatch != expectEM {
		t.Errorf("TestgetMatch1: expect %v. got %v\n", expectEM, mt.ExactMatch)
	}
}

func TestgetMatch2(t *testing.T) {
	fmt.Printf("fulltext.TestgetMatch0: Begin unit test\n")
	txt := "厚人倫，美教化，移風俗。故詩有六義焉：一曰風，二曰賦，三曰比，四曰興，五曰雅，六曰頌。"
	queryTerms := []string{"故", "詩", "一"}
	mt := getMatch(txt, queryTerms)
	if mt.Snippet == "" {
		t.Errorf("TestgetMatch2: snippet empty\n")
	}
	expectLM := "故詩"
	if mt.LongestMatch != expectLM {
		t.Errorf("TestgetMatch2: expect %s. got %s\n", expectLM, mt.LongestMatch)
	}
	expectEM := false
	if mt.ExactMatch != expectEM {
		t.Errorf("TestgetMatch2: expect %v. got %v\n", expectEM, mt.ExactMatch)
	}
}

func TestgetMatch3(t *testing.T) {
	fmt.Printf("fulltext.TestgetMatch0: Begin unit test\n")
	txt := "厚人倫，美教化，移風俗。故詩有六義焉：一曰風，二曰賦，三曰比，四曰興，五曰雅，六曰頌。"
	queryTerms := []string{"一", "詩", "有"}
	mt := getMatch(txt, queryTerms)
	if mt.Snippet == "" {
		t.Errorf("TestgetMatch3: snippet empty\n")
	}
	expectLM := "詩有"
	if mt.LongestMatch != expectLM {
		t.Errorf("TestgetMatch3: expect %s. got %s\n", expectLM, mt.LongestMatch)
	}
	expectEM := false
	if mt.ExactMatch != expectEM {
		t.Errorf("TestgetMatch3: expect %v. got %v\n", expectEM, mt.ExactMatch)
	}
}

func TestgetMatch4(t *testing.T) {
	fmt.Printf("fulltext.TestgetMatch0: Begin unit test\n")
	txt := "厚人倫，美教化，移風俗。故詩有六義焉：一曰風，二曰賦，三曰比，四曰興，五曰雅，六曰頌。"
	queryTerms := []string{"美", "移", "故"}
	mt := getMatch(txt, queryTerms)
	if mt.Snippet == "" {
		t.Errorf("TestgetMatch4: snippet empty\n")
	}
	if mt.LongestMatch == "" {
		t.Errorf("TestgetMatch4: LongestMatch empty\n")
	}
	expectEM := false
	if mt.ExactMatch != expectEM {
		t.Errorf("TestgetMatch4: expect %v. got %v\n", expectEM, mt.ExactMatch)
	}
}

// Test to load a local file
func TestGetMatching1(t *testing.T) {
	loader := LocalTextLoader{"../../../../corpus"}
	queryTerms := []string{"曰風"}
	mt, err := loader.GetMatching("shijing/shijing001.txt", queryTerms)
	if err != nil {
		t.Errorf("TestGetMatching1: got an error %v\n", err)
	}
	if mt.Snippet == "" {
		t.Errorf("TestGetMatching1: snippet empty\n")
	}
	fmt.Printf("fulltext.TestGetMatching1: match: %v\n", mt)
}

// Test to load a local file
func TestGetMatching2(t *testing.T) {
	fmt.Printf("fulltext.TestGetMatching: Begin unit test\n")
	loader := LocalTextLoader{"../../../../corpus"}
	queryTerms := []string{"曰風", "曰"}
	mt, err := loader.GetMatching("shijing/shijing001.txt", queryTerms)
	if err != nil {
		t.Errorf("TestGetMatching2: got an error %v\n", err)
	}
	if mt.Snippet == "" {
		t.Errorf("TestGetMatching2: snippet empty\n")
	}
	fmt.Printf("fulltext.TestGetMatching2: match: %v\n", mt)
}

// Test to load a local file
func TestGetMatching3(t *testing.T) {
	fmt.Printf("fulltext.TestGetMatching: Begin unit test\n")
	loader := LocalTextLoader{"../../../../corpus"}
	queryTerms := []string{"曰", "曰風"}
	mt, err := loader.GetMatching("shijing/shijing001.txt", queryTerms)
	if err != nil {
		t.Errorf("TestGetMatching3: got an error %v\n", err)
	}
	if mt.Snippet == "" {
		t.Errorf("TestGetMatching3: snippet empty\n")
	}
	fmt.Printf("fulltext.TestGetMatching3: match: %v\n", mt)
}
