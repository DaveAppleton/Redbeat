package main

import (
	"fmt"
)

func main() {
	allData := [3]string{"Tom", "Dick", "Harry"}

	for j := 0; j < len(allData); j++ {
		fmt.Println(allData[j], "wants to be a Go programmer")
	}
}
