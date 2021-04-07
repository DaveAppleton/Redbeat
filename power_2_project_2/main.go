package main

import (
	"fmt"
	"reflect"

	"af.sorry.ds/datasupply"
)

func main() {
	myRequest := datasupply.NewBid(23, 54)
	fmt.Println(reflect.TypeOf(*myRequest), *myRequest)

	myNiceBid := datasupply.Bid{Amount: 32, Age: 23}
	fmt.Println(reflect.TypeOf(myNiceBid), myNiceBid)

	// j := datasupply.hBid{} // cannot refer to unexported name datasupply.hBid

	myHRequest := datasupply.NewHBid(100, 99)
	fmt.Println(reflect.TypeOf(*myHRequest), *myHRequest)
	myHRequest.DoSomething()

	// myHRequest.age = 19  // cannot refer to unexported field or method age
}
