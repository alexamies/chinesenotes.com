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
	"github.com/alexamies/cnweb/find"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

const STATIC_DIR string = "./static"

// Finds documents matching the given query
func findHandler(response http.ResponseWriter, request *http.Request) {
	log.Print("main.findHandler, enter")
	results := find.QueryResults{}
	resultsJson, _ := json.Marshal(results)
	response.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(response, string(resultsJson))
}

func main() {
	log.Print("End-to-end test server started")
	r := mux.NewRouter()
	r.HandleFunc("/find/", findHandler)
	r.Handle("/", handlers.ContentTypeHandler(http.FileServer(http.Dir(STATIC_DIR))))
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := fmt.Sprintf(":%s", port)
	srv := &http.Server{
        Handler:      r,
        Addr:         addr,
        WriteTimeout: 60 * time.Second,
        ReadTimeout:  60 * time.Second,
    }
	log.Fatal(srv.ListenAndServe())
}