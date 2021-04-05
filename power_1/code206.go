package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"givenname"`
	Age  int
}

func main() {
	var kiddy person
	str := `{"givenname":"Layla","Age":4}`
	if err := json.Unmarshal([]byte(str), &kiddy); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(kiddy.Name, "is", kiddy.Age, "years old")
	}
}
