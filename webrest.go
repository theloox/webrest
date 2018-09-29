package main

import (
    "encoding/json"
    "fmt"
    "net/http"

    //"google.golang.org/appengine"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello my baby Hello my honey, %s", r.URL.Path[1:])
}

type Message struct {
    Msg string
}

func about (w http.ResponseWriter, r *http.Request) {

    m := Message{"webapp, build v0.0"}
    b, err := json.Marshal(m)

    if err != nil {
        panic(err)
    }

     w.Write(b)
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/about/", about)

    http.ListenAndServe(":8080", nil)
}

