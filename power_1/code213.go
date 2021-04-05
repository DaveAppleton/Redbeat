package main

import (
	"fmt"
	"time"
)

// start OMIT
var counter int

func increment(tim int) {
	for {
		a := counter
		time.Sleep(time.Duration(tim) * time.Millisecond)

		fmt.Println(">", a)

		counter = a + 1
	}
}

func main() {
	go increment(100)

	go increment(200)
	time.Sleep(time.Second * 20)
}

// end OMIT
/*
race start OMIT
$ go run -race .
==================
WARNING: DATA RACE
Read at 0x00000121d848 by goroutine 7:
  main.display()
      /Users/daveappleton/Documents/training/TC/race/start.go:20 +0x3e

Previous write at 0x00000121d848 by goroutine 6:
  main.increment()
      /Users/daveappleton/Documents/training/TC/race/start.go:13 +0x56

Goroutine 7 (running) created at:
  main.main()
      /Users/daveappleton/Documents/training/TC/race/start.go:27 +0x5a

Goroutine 6 (running) created at:
  main.main()
      /Users/daveappleton/Documents/training/TC/race/start.go:26 +0x42
==================
1
race end OMIT
*/
