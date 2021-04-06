package main

import (
	"fmt"
	"time"
)

// start OMIT
func loop(id string, delay int, max int, ch chan int) {

	for j := 0; j < max; j++ {
		fmt.Println(id, j)
		ch <- j
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
	close(ch)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go loop("first", 150, 20, ch1)
	go loop("second", 120, 12, ch2)
	var done1, done2 bool
	var weiting int
	// main OMIT
	for {
		select {
		case num, ok := <-ch1:
			if ok {
				fmt.Println("rx 1", num)
			} else {
				done1 = true
			}

		case num, ok := <-ch2:
			if ok {
				fmt.Println("rx 2", num)
			} else {
				done2 = true
			}
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
