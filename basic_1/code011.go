package main

import (
	"fmt"
)

func factorial(start uint) (result uint) {
	result = 1
	for j := start; j > 1; j-- {
		result = result * j
	}
	return
}

func main() {
	n := uint(5)
	fmt.Println("the factorial of", n, "is", factorial(n))
}
