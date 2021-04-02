package main

import (
	"fmt"
)

// test some maths
func main() {
	a := 3
	b := 4
	fmt.Println("addition  ", a, "+", b, "=", a+b)
	fmt.Println("multiply  ", a, "*", b, "=", a*b)
	fmt.Println("subtract  ", a, "-", b, "=", a-b)
	fmt.Println("divide    ", a, "/", b, "=", a/b)
	fmt.Println("divide    ", b, "/", a, "=", b/a)
}
