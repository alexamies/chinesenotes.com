// Unit tests for the find package
package find

import (
	"log"
	"testing"
)

// Test package initialization, which requires a database connection
func TestInit(t *testing.T) {
	log.Printf("TestInit: Begin unit tests\n")
}

func TestFindDocuments1(t *testing.T) {
	qr, err := FindDocuments("Assembly")
	if err != nil {
		t.Error("TestFindDocuments1: got error, ", err)
	}
	if len(qr.Words) != 0 {
		t.Error("TestFindDocuments1: len(qr.Words) != 0, ", len(qr.Words) != 0)
	}
}

func TestFindDocuments2(t *testing.T) {
	_, err := FindDocuments("")
	if err == nil {
		t.Error("TestFindDocuments2: expected error for empty string")
	}
}

func TestFindWords1(t *testing.T) {
	words, err := findWords("Assembly")
	if err != nil {
		t.Error("TestFindWords1: got error, ", err)
	}
	if len(words) != 0 {
		t.Error("TestFindWords1: len(words) != 0, ", len(words))
	}
}

func TestFindWords2(t *testing.T) {
	words, err := findWords("金剛")
	if err != nil {
		t.Error("TestFindWords2: got error, ", err)
	}
	if len(words) != 1 {
		t.Error("TestFindWords2: len(words) != 1, ", len(words))
	}
}

