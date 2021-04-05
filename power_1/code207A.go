package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// start OMIT
func main() {
	var p struct {
		Name struct {
			First string `json:"first_name"`
			Last  string
		}
		Age int `json:"age_of_hero"`
	}
	js := `{ "name" : { "first_name": "Steve", "last" : "Miller"}, "Age_of_hero": 70}`

	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(p)

}

// end OMIT
