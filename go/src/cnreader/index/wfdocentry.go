/*
Library for sorting word frequency item by relative frequency
*/
package index

import (
	"fmt"
)

// A document-specific word frequency entry record
type WFDocEntry struct {
	Filename string
	Count    int
	Freq    float64 // frequency per 10,000 words
}

func (item WFDocEntry) String() string {
	return fmt.Sprintf("%s %f", item.Filename, item.Freq)
}

type ByFrequencyDoc []WFDocEntry

func (items ByFrequencyDoc) Len() int {
	return len(items)
}

func (items ByFrequencyDoc) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

func (items ByFrequencyDoc) Less(i, j int) bool {
	return items[i].Count > items[j].Count
}