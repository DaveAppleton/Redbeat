package main

import (
	"fmt"
)

func main() {
	myName := "Dave"
	myAge := 16
	fmt.Println("My name is", myName, "and I am", myAge, "years old")
	myName := 16
	myAge := "Dave"
	fmt.Println("My name is", myName, "and I am", myAge, "years old")
}
