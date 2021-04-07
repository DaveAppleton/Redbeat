package main

import (
	"encoding/json"
	"fmt"
)

type Bid struct {
	Amount int
	Age    int
}

func main() {
	// define and initialise slice of strings
	mySlice := []string{"one", "two", "three"}
	fmt.Println(mySlice)

	yourBid := Bid{Amount: 23, Age: 84}
	fmt.Println(yourBid)

	myBid := struct {
		Amount int
		Age    int
	}{Amount: 17, Age: 42}
	//	myBid.Age = 17
	//	myBid.Amount = 42

	fmt.Println(myBid)

	data, _ := json.Marshal(myBid)
	fmt.Println(string(data))

}
