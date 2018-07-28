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
}

func (item WFDocEntry) String() string {
	return fmt.Sprintf("%s %d", item.Filename, item.Count)
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