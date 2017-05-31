package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	fmt.Println("hello world")
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}
