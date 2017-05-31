package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Fruit struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Color  string `json:"color,omitempty"`
	Rating int    `json:"rating,omitempty"`
}

var fruits []Fruit

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "items.html")
}

func GetFruitsEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range fruits {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Fruit{})
}

func GetFruitEndpoint(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(fruits)
}

func main() {
	router := mux.NewRouter()
	fruits = append(fruits, Fruit{ID: "1", Name: "Orange", Color: "Orange", Rating: 1})
	fruits = append(fruits, Fruit{ID: "2", Name: "Banana", Color: "Yellow", Rating: 5})
	router.HandleFunc("/", index)
	router.HandleFunc("/fruits", GetFruitEndpoint).Methods("GET")
	router.HandleFunc("/fruits/{id}", GetFruitsEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
