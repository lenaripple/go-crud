package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type cities struct {
	name    string
	state   string
	country string
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", hello)
	router.HandleFunc("/items", items)
	router.HandleFunc("/new", newItem)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func hello(w http.ResponseWriter, r *http.Request) {
	// http.ServeFile(w, r, "index.html")
	fmt.Fprintln(w, "Hello, World", html.EscapeString(r.URL.Path))
}

func items(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Items")
}

func newItem(w http.ResponseWriter, r *http.Request) {

}
