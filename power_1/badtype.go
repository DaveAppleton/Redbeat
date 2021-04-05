package main

import "fmt"

var j int

func printMessage(message string) {
	fmt.Println(message)
}

func main() {
	printMessage("Hello")
	printMessage(42)
	j = "97"
}
