package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// start OMIT
type person struct {
	Name string `json:"givenname"`
	Age  int
}

func index(w http.ResponseWriter, r *http.Request) {
	kiddie := new(person)
	err := json.NewDecoder(r.Body).Decode(kiddie)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(kiddie.Name, "is", kiddie.Age, "years old")
	}
}

// end OMIT
func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
