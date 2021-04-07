package main

import "fmt"

func main() {
	var myArray [10]int
	myArray[3] = 2
	myArray[1] = 17
	myArray[0] = 2
	fmt.Println(myArray[0])
	fmt.Println(myArray[0:3])
}
