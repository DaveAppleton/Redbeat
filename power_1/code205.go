package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	Name string `json:"givenname"`
	Age  int
}

func main() {
	layla := person{Name: "Layla Asifa binti Saya", Age: 6}
	coded, err := json.Marshal(&layla)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(coded))
}
