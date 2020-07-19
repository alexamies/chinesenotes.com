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

package dictionary

import (
	"strings"
	"testing"
)

// Test for dictionar validation
func TestValidate(t *testing.T) {
	const posList = "noun\nverb\n"
	posReader := strings.NewReader(posList)
	const domainList = "艺术	Art	\\N	\\N\n佛教	Buddhism	\\N	\\N\n"
	domainReader := strings.NewReader(domainList)
	validator, err := NewValidator(posReader, domainReader)
	if err != nil {
		t.Fatalf("TestNewValidator: Unexpected error: %v", err)
	}
	type test struct {
		name string
		pos string
		domain string
		valid bool
  }
  tests := []test{
		{
			name: "Valid term",
			pos: "noun",
			domain: "Art",
			valid: true,
		},
		{
			name: "Invalid domain",
			pos: "noun",
			domain: "Artistic",
			valid: false,
		},
  }
  for _, tc := range tests {
		err = validator.Validate(tc.pos, tc.domain)		
		if tc.valid && err != nil {
			t.Fatalf("TestNewValidator: unexpected error for %s, %v", tc.name, err)
		}
		if !tc.valid && err == nil {
			t.Errorf("TestNewValidator: expected error for test %s", tc.name)
		}
	}
}
