// Test sorting of relative word frequencies
package index

import (
	"fmt"
	"sort"
	"testing"
)

// Trivial test for sorting of word frequencies
func TestSort0(t *testing.T) {
	items := []WFDocEntry{}
	sort.Sort(ByFrequencyDoc(items))
}

// Simple test of sorting of word frequencies
func TestSort1(t *testing.T) {
	item1 := WFDocEntry{"test.txt", 2}
	items := []WFDocEntry{item1}
	sort.Sort(ByFrequencyDoc(items))
	fmt.Println(items)
}

// Easy test of sorting of word frequencies
func TestSort2(t *testing.T) {
	item1 := WFDocEntry{"test.txt", 3}
	item2 := WFDocEntry{"test2.txt", 4}
	item3 := WFDocEntry{"test3.txt", 5}
	items := []WFDocEntry{item1, item2, item3}
	sort.Sort(ByFrequencyDoc(items))
	fmt.Println(items)
}