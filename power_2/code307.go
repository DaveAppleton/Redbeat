package main

import (
	"fmt"
	"time"
)

// main OMIT
func main() {
	timer3 := time.NewTimer(time.Second)
	count := 0
	var done bool
	for {
		select {
		case <-timer3.C:
			fmt.Println("timer fired at ", count)
			done = true
		default:
			count++
		}
		if done {
			break
		}
	}
}

// end OMIT
