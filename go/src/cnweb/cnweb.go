// Chinese notes web app
package main

import (
    "fmt"
    "log"
    "net/http"
)

const (
    top = `<!DOCTYPE HTML><html lang="en">
<head><title>Chinese Notes Test Page</title></head>
<body><h3>Chinese-English Dictionary Lookup</h3>
<p>Lookup a word, eg 中文</p>`
    form = `<p><form action="/" method="POST">
<input type="text" name="lookup" size="30">
<input type="submit" value="Lookup">
</form></p>`
    bottom = `</body></html>`
)

func main() {
    http.HandleFunc("/", lookup)
    if err := http.ListenAndServe(":8000", nil); err != nil {
        log.Fatal("failed to start go http server", err)
    }
}

func lookup(writer http.ResponseWriter, request *http.Request) {
    err := request.ParseForm()
    fmt.Fprint(writer, top)
    if err != nil {
        fmt.Fprintf(writer, "Error: ", err)
    } else {
        if word, post := processRequest(request); post {
            fmt.Printf("home: Got word: ", word, "\n")
            fmt.Fprint(writer, "<p>Got word: ", word, "</p>")
        } 
    }
    fmt.Fprint(writer, form, bottom)
}

func processRequest(request *http.Request) (string, bool) {
    fmt.Printf("processRequest: Got a request\n")
    if slice, found := request.Form["lookup"]; found && len(slice) > 0 {
        return slice[0], true
    }
    return "", false
}
