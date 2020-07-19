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

// Package for dictionary definitions and loading

package dictionary

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Performs validation of dictionary entries.
// Use NewValidator to create a Validator.
type Validator interface {
	Validate(pos, domain string) error
}

// Performs validation of dictionary entries.
// Use NewValidator to create a Validator.
type validator struct {
	validPos map[string]bool
	validDomains map[string]bool
}

// Crates a Validator with the given readers
// Params:
//   posReader Reader to load the valid parts of speech from
//   domainReader Reader to load the valid subject domains
// Returns:
//   An initialized Validator
func NewValidator(posReader io.Reader, domainReader io.Reader) (Validator, error) {
	validPos := make(map[string]bool)
	posFScanner := bufio.NewScanner(posReader)
	for posFScanner.Scan() {
		pos := posFScanner.Text()
		validPos[pos] = true
	}
	if err := posFScanner.Err(); err != nil {
		return nil, fmt.Errorf("Could not read list of valid parts of speech: %v", err)
	}
	validDomains := make(map[string]bool)
	dFScanner := bufio.NewScanner(domainReader)
	for dFScanner.Scan() {
		line := dFScanner.Text()
		parts := strings.Split(line, "\t")
		if len(parts) != 4 {
			return nil, fmt.Errorf("Could not parse list of valid domains: %s", line)
		}
		validDomains[parts[1]] = true
	}
	if err := dFScanner.Err(); err != nil {
		return nil, fmt.Errorf("Could not read list of valid domains: %v", err)
	}
	return validator{
		validPos: validPos,
		validDomains: validDomains,
	}, nil
}

func (val validator) Validate(pos, domain string) error {
	if pos != "\\N" {
		if _, ok := val.validPos[pos]; !ok {
			return fmt.Errorf("%s is not a recognized part of speech", pos)
		}
	}
	if _, ok := val.validDomains[domain]; !ok {
		return fmt.Errorf("%s is not a recognized domain", domain)
	}
	return nil
}