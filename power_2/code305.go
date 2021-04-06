package main

import (
	"fmt"
	"time"
)

// start OMIT
func loop(id string, max int, ch chan int) {

	for j := 0; j < max; j++ {
		fmt.Println(id, j)
		ch <- j
		time.Sleep(time.Duration(j) * time.Second)
	}
	close(ch)
}

// main OMIT
func main() {
	var done bool
	ch := make(chan int, 5)

	go loop("solo", 14, ch)

	for {
		select {
		case num, ok := <-ch:
			if ok {
				fmt.Println("received", num)
			} else {
				fmt.Println("channel closed")
				done = true
			}
		case <-time.After(10 * time.Second):
			fmt.Println("got tired of waiting")

		}
		if done {
			break
		}
	}
}

// end OMIT
