package corpus

import (
	"fmt"
	"testing"
)

// Trivial test for excluded string
func TestIsExcluded0(t *testing.T) {
	fmt.Printf("corpus.TestIsExcluded0: Begin unit tests\n")
	if IsExcluded("") {
		t.Error("corpus.TestIsExcluded0: Do not expect '' to be excluded")
	}
}

// Easy test for excluded string
func TestIsExcluded1(t *testing.T) {
	fmt.Printf("corpus.TesIsExcluded: Begin unit tests\n")
	if !IsExcluded("如需引用文章") {
		t.Error("corpus.TestIsExcluded0: Expect '如需引用文章' to be excluded")
	}
}
