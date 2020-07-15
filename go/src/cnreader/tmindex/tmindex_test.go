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

// Test tmindex functions
package tmindex

import (
	"bytes"
	"fmt"
	"github.com/alexamies/chinesenotes-go/dicttypes"
	"testing"
)

func mockDictionary() map[string]dicttypes.Word {
	s1 := "结实"
	t1 := "結實"
	pinyin := "jiēshi "
	hw1 := dicttypes.Word{
		Simplified: s1, 
		Traditional: t1,
		Pinyin: pinyin,
		Senses: []dicttypes.WordSense{},
	}
	s2 := "倿"
	t2 := "\\N"
	hw2 := dicttypes.Word{
		Simplified: s2, 
		Traditional: t2,
		Pinyin: pinyin,
		Senses: []dicttypes.WordSense{},
	}
  headwords := []dicttypes.Word{hw1, hw2}
  wdict := make(map[string]dicttypes.Word)
  for _, hw := range headwords {
  	wdict[hw.Simplified] = dicttypes.Word{}
  	trad := hw.Traditional
  	if trad != "\\N" {
  		wdict[trad] = dicttypes.Word{}
  	}
  }
  return wdict
}

// Test basic BuildIndex functions
func TestBuildIndex(t *testing.T) {
	fmt.Printf("TestBuildIndex: Begin unit test\n")
	wdict := mockDictionary()
	var buf bytes.Buffer
	BuildIndex(&buf, wdict)
	expected := 
`结	结实
实	结实
結	結實
實	結實
倿	倿
`
	result := buf.String()
	if len(result) != len(expected) {
		t.Errorf("TestBuildIndex, expected: %d, got: %d", len(expected), len(result))
	}
}
