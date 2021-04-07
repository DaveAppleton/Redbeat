package main

import (
	"fmt"
	"reflect"
)

// start OMIT
func pr(name string, data interface{}) {
	fmt.Println(name, reflect.TypeOf(data), data)
}

func main() {
	myArray := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	mySlice := myArray[:]
	pr("mySlice", mySlice)
	pr("myArray", myArray)
	//myArray = append(myArray, 10) // does not compile
	mySlice = append(mySlice, 10)
	yourSlice := append(myArray[:], 10)
	pr("mySlice+1", mySlice)
	pr("yourSlice", yourSlice)
	var yourArray [11]int
	copy(yourArray[:], yourSlice[:])
	pr("yourArray", yourArray)
}

// end OMIT
