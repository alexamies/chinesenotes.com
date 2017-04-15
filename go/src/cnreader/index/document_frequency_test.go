// Test document frequency
package index

import (
	"fmt"
	"testing"
)

// Trivial test for document frequency
func TestAddVocabulary0(t *testing.T) {
	fmt.Println("index.DocumentFrequency.TestAddVocabulary enter")
	df := DocumentFrequency{}
	vocab := map[string]int{}
	df.AddVocabulary(vocab)
}
