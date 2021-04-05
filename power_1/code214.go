package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// start OMIT
var counter int32

func increment() {
	for {
		atomic.AddInt32(&counter, 1)
		time.Sleep(100 * time.Millisecond)
	}
}

func display() {
	for {
		fmt.Println(atomic.LoadInt32(&counter))
		time.Sleep(90 * time.Millisecond)
	}
}

func main() {
	go increment()
	go display()
	time.Sleep(time.Second * 20)
}

// end OMIT
/*
race start OMIT
$ go run -race .
1
1
2
3
4
5
6
7
8
9
10
race end OMIT
*/
