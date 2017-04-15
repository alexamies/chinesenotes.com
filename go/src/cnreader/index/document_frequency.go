/*
For every term, store the number of documents that contain the term
*/
package index

// Map from term to number of documents referencing the term
type DocumentFrequency map[string]int

// Adds the given vocabulary to the map
// Param:
//   vocab - word frequencies are ignored, only the presence of the term is 
//           important
func (df DocumentFrequency) AddVocabulary(vocab map[string]int) {
	for k, _ := range vocab {
		_, ok := df[k]
		if ok {
			df[k]++
		} else {
			df[k] = 1
		}
	}
}