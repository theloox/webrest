package main

import (
	"strconv"
    "encoding/json"
    "fmt"
    "net/http"

	"github.com/gorilla/mux"
    //"google.golang.org/appengine"
)

type Element struct {
	ID int		`json:"id,omitempty`
	Times int	`json:"times,omitempty`
	Name string	`json:"name,omitempty`
}

type Elements []Element

type Message struct {
	Msg string
}

var list = Elements {
	Element{ID: 1, Times: 100, Name: "this is 1"},
	Element{ID: 2, Times: 2000, Name: "maybe 2"},
	Element{ID: 3, Times: 30000, Name: "and 3"},
}


func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<html><body>")

    fmt.Fprintf(w, "Hello my baby Hello my honey, %s", r.URL.Path[1:])
    fmt.Fprintf(w, "<br>", )
    fmt.Fprintf(w, "now go to <a href=\"/app\">/app</a>")

	fmt.Fprintf(w, "</body></html>")

}

func app(w http.ResponseWriter, r *http.Request) {

    b, err := json.Marshal(list)

    if err != nil {
        panic(err)
    }

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(b)
}

func appid(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
    id_ := vars["appid"]

	id, _ := strconv.Atoi(id_)

	var ret Element

	for _, e := range list {
		if e.ID == id {
			ret = e
		}
	}

    b, err := json.Marshal(ret)

    if err != nil {
        panic(err)
    }

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(b)
}

func about(w http.ResponseWriter, r *http.Request) {

    m := Message{"webrest, build v0.0"}
    b, err := json.Marshal(m)

    if err != nil {
        panic(err)
    }

     w.Write(b)
}

func main() {
	router := mux.NewRouter()

    router.HandleFunc("/", index)
    router.HandleFunc("/app", app)
    router.HandleFunc("/app/", app)
    router.HandleFunc("/app/{appid}", appid)

    http.ListenAndServe(":8080", router)
}
