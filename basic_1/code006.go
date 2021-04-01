package main

import (
	"fmt"
)

func main() {
	var allData [3]string
	allData[0] = "Tom"
	allData[1] = "Dick"
	allData[2] = "Harry"

	fmt.Println(allData)
	fmt.Println("Why should you program in Go?")
	fmt.Println("Every", allData[0], ",", allData[1], "and", allData[2], "wants to be a Node programmer")
}
