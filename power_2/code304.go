package main

import (
	"fmt"
)

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	x, ok := <-queue
	if ok {
		fmt.Println("i found", x)
	} else {
		fmt.Println("not OK")
	}
	// wow we can walk through the data!
	for elem := range queue {
		fmt.Println(elem)
	}
}
