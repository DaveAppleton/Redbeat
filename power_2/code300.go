package main

import (
	"fmt"
	"time"
)

// source OMIT
func source(ch chan int) {
	j := 0
	for {
		fmt.Println("sending ", j)
		ch <- j
		j++
		time.Sleep(10 * time.Millisecond)
	}
}

// sink OMIT
func sink(ch chan int) {
	for {
		n := <-ch
		fmt.Println(n)
		time.Sleep(time.Second)
	}
}

// main OMIT
func main() {
	ch := make(chan int, 5)
	go source(ch)
	go sink(ch)
	time.Sleep(45 * time.Second)
}

// end OMIT
