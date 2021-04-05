package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	Name string `json:"givenname"`
	Age  int
}

func index(w http.ResponseWriter, r *http.Request) {
	layla := person{Name: "Layla", Age: 4}
	if err := json.NewEncoder(w).Encode(layla); err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
