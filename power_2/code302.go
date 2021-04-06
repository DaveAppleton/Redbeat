package main

import (
	"fmt"
	"time"
)

// start OMIT
func loop(id string, delay int, ch chan int) {

	for j := 0; j < 20; j++ {
		fmt.Println(id, j)
		ch <- j
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
}

// main OMIT
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go loop("first", 150, ch1)
	go loop("second", 120, ch2)
	var done1, done2 bool
	var weiting int
	for {
		select {
		case num1 := <-ch1:
			done1 = done1 || num1 == 19

		case num2 := <-ch2:
			done2 = done2 || num2 == 19
		default:
			weiting++
			time.Sleep(10 * time.Millisecond)
		}
		if done1 && done2 {
			fmt.Println("waited", weiting, "times")
			break
		}
	}
}

// end OMIT
