package main

import "fmt"

// try each of the first 4 lines to see what happens

func main() {
	var mySlice []int
	//mySlice := []int{}
	//mySlice := []int{8,8,8,8}
	//mySlice := make([]int,5)
	mySlice[3] = 2
	mySlice[1] = 17
	mySlice[0] = 2
	fmt.Println(mySlice[0])
	fmt.Println(mySlice[0:3])

}
