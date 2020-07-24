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

// Package for translation memory index
package tmindex

import (
	"bufio"
	"fmt"
	"github.com/alexamies/chinesenotes-go/dicttypes"
	"io"
	"os"
)

const (
	fnameUni = "tmindex_unigram.tsv"
	fnameDomain = "tmindex_uni_domain.tsv"
)

type indexEntry struct {
	c string
	term string
	count int
}

// Builds a unigram index with domain
func BuildIndexes(indexDir string, wdict map[string]dicttypes.Word) error {
	pathUni := fmt.Sprintf("%s/%s", indexDir, fnameUni)
	fUni, err := os.Create(pathUni)
	defer fUni.Close()
	if err != nil {
		return fmt.Errorf("Could not create index file %s, err: %v\n", fnameUni, err)
	}
	wUni := bufio.NewWriter(fUni)
	err = buildUnigramIndex(wUni, wdict)
	if err != nil {
		return fmt.Errorf("could not write to index file %s, err: %v\n", fnameUni, err)
	}

	pathDomain := fmt.Sprintf("%s/%s", indexDir, fnameDomain)
	fDomain, err := os.Create(pathDomain)
	defer fDomain.Close()
	if err != nil {
		return fmt.Errorf("could not create index file %s, err: %v\n", fnameDomain, err)
	}
	wDomain := bufio.NewWriter(fDomain)
	err = buildUniDomainIndex(wDomain, wdict)
	if err != nil {
		return fmt.Errorf("could not write to index file %s, err: %v\n", fnameDomain, err)
	}
	return nil
}

// Builds a unigram index with domain
func buildUniDomainIndex(w io.Writer, wdict map[string]dicttypes.Word) error {
	tmindexUni := make(map[string]bool)
	for term, word := range wdict {
		for _, sense := range word.Senses {
			domain := sense.Domain
			for _, c := range term {
		  	line := fmt.Sprintf("%c\t%s\t%s\n", c, term, domain)
				if _, ok := tmindexUni[line]; ok {
					continue
				}
		  	_, err := io.WriteString(w, line)
		  	if err != nil {
		  		return err
		  	}
		  	tmindexUni[line] = true
			}
		}
	}
	return nil
}

// Builds a unigram index
func buildUnigramIndex(w io.Writer, wdict map[string]dicttypes.Word) error {
	tmindexUni := make(map[string]bool)
	for term := range wdict {
		for _, c := range term {
	  	line := fmt.Sprintf("%c\t%s\n", c, term)
			if _, ok := tmindexUni[line]; ok {
				continue
			}
	  	_, err := io.WriteString(w, line)
	  	if err != nil {
	  		return err
	  	}
	  	tmindexUni[line] = true
		}
	}
	return nil
}