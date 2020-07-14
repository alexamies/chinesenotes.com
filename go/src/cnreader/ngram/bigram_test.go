// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Test bigram functions
package ngram

import (
	"fmt"
	"github.com/alexamies/cnreader/dictionary"
	"testing"
)

// Test basic Bigram functions
func TestBigram(t *testing.T) {
	fmt.Printf("TestBigram: Begin unit test\n")
	s1 := "诸"
	s2 := "諸"
	hw1 := dictionary.HeadwordDef{
		Id: 1,
		Simplified: &s1, 
		Traditional: &s2,
		Pinyin: []string{},
		WordSenses: &[]dictionary.WordSenseEntry{},
	}
	s3 := "倿"
	s4 := "\\N"
	hw2 := dictionary.HeadwordDef{
		Id: 2,
		Simplified: &s3, 
		Traditional: &s4,
		Pinyin: []string{},
		WordSenses: &[]dictionary.WordSenseEntry{},
	}
	example := ""
	exFile := ""
	exDocTitle := ""
	exColTitle := ""
	b := NewBigram(hw1, hw2, example, exFile, exDocTitle, exColTitle)
	r := b.Traditional()
	e := "諸倿"
	if r != e {
		t.Error("TestBigram, expected ", e, " got, ", r)
	}
}
