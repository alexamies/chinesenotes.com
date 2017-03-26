/* 
Web utility for cnreader
 */
package main

import (
	//"cnreader/index"
	"fmt"
	"log"
	"net/http"
)

func handler(response http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(response, "Hello Go proxy!")
}

//Entry point for the chinesenotes go web application
func main() {
	log.Printf("cnweb: Web app started\n")
	//index.LoadKeywordIndex()
	//documents := index.FindForKeyword("你")
	//log.Printf("cnweb: found %v\n", documents)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}