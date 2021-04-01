package main

import (
	"fmt"
)

func buyTicket(show string) {
	fmt.Println("buying ticket for", show)
}
func takeTaxi(destination string) {
	fmt.Println("taking taxi to", destination)
}
func seeShow(show string) {
	fmt.Println("watching", show)
}

func main() {
	movie := "Ms Marvel"
	buyTicket(movie)
	takeTaxi("movie theatre")
	seeShow(movie)
}
