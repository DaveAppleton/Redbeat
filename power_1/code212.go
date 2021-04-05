package main

import (
	"fmt"
	"time"
)

func thread(prefix string, delay int) {
	i := 0
	for {
		fmt.Println(prefix, i)
		i++
		time.Sleep(time.Millisecond * time.Duration(delay))
	}
}

func main() {
	go thread("500", 500)
	go thread("600", 600)
	time.Sleep(time.Second * 15)
}
