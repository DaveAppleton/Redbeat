package main

import "fmt"

func testv(p int) {
	for j := 0; j < 5; j++ {
		fmt.Println(p)
		p++
	}
}

func testp(p *int) {
	for j := 0; j < 5; j++ {
		fmt.Println(*p)
		*p++
	}
}

func main() {
	myVariable := 0
	testv(myVariable)
	fmt.Println("***", myVariable)
	testp(&myVariable)
	fmt.Println("***", myVariable)
}
