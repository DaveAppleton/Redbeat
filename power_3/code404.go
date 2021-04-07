package main

import "fmt"

func main() {
	mySlice := []int{5, 4, 3, 2, 1}
	yourSlice := []int{1, 2, 3, 4, 5}
	mySlice = append(mySlice, 0)
	fmt.Println(mySlice)
	// notice the three dots
	yourSlice = append(yourSlice, mySlice...)
	fmt.Println(yourSlice)
}
