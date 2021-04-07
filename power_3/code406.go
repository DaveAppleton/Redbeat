package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func pr(name string, data interface{}) {
	fmt.Println(name, reflect.TypeOf(data), data)
}

// start OMIT
type Bid struct {
	Amount int
	Age    int
}

func main() {
	yourBid := Bid{Amount: 23, Age: 84}
	pr("yourBid", yourBid)

	myBid := struct {
		Amount int
		Age    int
	}{Amount: 17, Age: 42}
	myBid.Age = 99

	pr("myBid", myBid)

	data, _ := json.Marshal(myBid)
	fmt.Println(string(data))
}

// end OMIT
