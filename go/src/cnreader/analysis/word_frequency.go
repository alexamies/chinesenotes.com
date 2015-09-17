/*
Library for Chinese vocabulary analysis
Word frequency functions.
*/
package analysis

import "sort"

// Sorted list of word frequencies
type SortedWF struct {
	wf map[string]int
	w []string
}

func (sortedWF *SortedWF) Len() int {
	return len(sortedWF.wf)
}

func (sortedWF *SortedWF) Less(i, j int) bool {
	return sortedWF.wf[sortedWF.w[i]] > sortedWF.wf[sortedWF.w[j]]
}

func (sortedWF *SortedWF) Swap(i, j int) {
	sortedWF.w[i], sortedWF.w[j] = sortedWF.w[j], sortedWF.w[i]
}

/*
 * Sorts based on word frequency
 */
func SortedFreq(wf map[string]int) []string {
	sortedWF := new(SortedWF)
	sortedWF.wf = wf
	sortedWF.w = make([]string, len(wf))
	i := 0
	for key, _ := range wf {
		sortedWF.w[i] = key
		i++
	}
	sort.Sort(sortedWF)
	return sortedWF.w
}