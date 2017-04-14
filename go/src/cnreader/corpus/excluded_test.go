package corpus

import (
	"fmt"
	"testing"
)

// Test excluded strings
func TestIsExcluded0(t *testing.T) {
	fmt.Printf("corpus.TesIsExcluded: Begin unit tests\n")
	IsExcluded("")
	if IsExcluded("") {
		t.Error("TestIsExcluded0: Do not expect empty string to be excluded")
	}
}
