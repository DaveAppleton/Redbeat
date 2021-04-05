package main

import (
	"fmt"
	"sync"

	"time"
)

// start OMIT
var (
	counter int32
	mu      sync.Mutex
)

func increment() {
	for {
		mu.Lock()
		counter++
		mu.Unlock()
		time.Sleep(100 * time.Millisecond)
	}
}

func display() {
	for {
		mu.Lock()
		fmt.Println(counter)
		mu.Unlock()
		time.Sleep(90 * time.Millisecond)
	}
}

// end OMIT
func main() {
	go increment()
	go display()
	time.Sleep(time.Second * 20)
}
