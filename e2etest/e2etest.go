/**
 * Licensed  under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package main

/**
 * End-to-end test program
 */

import (
	"encoding/json"
	"fmt"
	"github.com/alexamies/cnweb/dictionary"
	"github.com/alexamies/cnweb/find"
	"log"
	"net/http"
	"os"
)

const STATIC_DIR string = "./static"

// Finds documents matching the given query
func findHandler(response http.ResponseWriter, request *http.Request) {
	log.Print("main.findHandler, enter")
	url := request.URL
	queryString := url.Query()
	query := queryString["query"]
	q := "No Query"
	if len(query) > 0 {
		q = query[0]
	} else {
		query := queryString["text"]
		if len(query) > 0 {
			q = query[0]
		}
	}
  col := []find.Collection{}
  doc := []find.Document{}
	senses0 := []dictionary.WordSense{}
  terms0 := []find.TextSegment{}
	results := find.QueryResults{q, "", 0, 0, col, doc, terms0}
	sense1 := dictionary.WordSense{
		Id: 74517,
		HeadwordId: 74517,
		Simplified: "大哥大",
		Traditional: "\\N",
		Pinyin: "dàgēdà",
		English: "cell phone",
		Notes: "",
	}
	senses1 := []dictionary.WordSense{sense1}
	entry1 := dictionary.Word{
		Simplified: "大哥大",
		Traditional: "\\N",
		Pinyin: "dàgēdà",
		HeadwordId: 74517,
		Senses: senses1,
	}
	ts1 := find.TextSegment{
		QueryText: "大哥大",
		DictEntry: entry1,
		Senses: senses0,
	}
	terms1 := []find.TextSegment{ts1}
	sense2 := dictionary.WordSense{
		Id: 3940,
		HeadwordId: 3940,
		Simplified: "旧",
		Traditional: "舊",
		Pinyin: "jiù",
		English: "old",
		Notes: "",
	}
	senses2 := []dictionary.WordSense{sense2}
	entry2 := dictionary.Word{
		Simplified: "旧",
		Traditional: "舊",
		Pinyin: "jiù",
		HeadwordId: 3940,
		Senses: senses2,
	}
	ts2 := find.TextSegment{
		QueryText: "舊",
		DictEntry: entry2,
		Senses: senses0,
	}
	terms2 := []find.TextSegment{ts1, ts2}
	if (q == "大哥大") {
	  results = find.QueryResults{q, "", 0, 0, col, doc, terms1}
	} else if (q == "舊大哥大") {
	  results = find.QueryResults{q, "", 0, 0, col, doc, terms2}
	}

	resultsJson, _ := json.Marshal(results)
	response.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(response, string(resultsJson))
}

func main() {
	log.Print("End-to-end test server started")
	http.HandleFunc("/find/", findHandler)
	http.Handle("/", http.FileServer(http.Dir(STATIC_DIR)))
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := fmt.Sprintf(":%s", port)
	log.Fatal(http.ListenAndServe(addr, nil))
}