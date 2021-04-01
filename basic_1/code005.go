package main

import "fmt"

func main() {
	var myData struct {
		Name string
		Age  int
	}
	myData.Name = "Dave"
	myData.Age = 16
	fmt.Println(myData)
	fmt.Println("I am", myData.Name, "and I am", myData.Age, "years old")
}
