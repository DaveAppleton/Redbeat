package main

import (
	"fmt"
	"log"

	"af.sorry.logrates/oex"
)

func main() {
	rates, err := oex.GetLatest()
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	sgd := rates.Rates["SGD"]
	myr := rates.Rates["MYR"]
	fmt.Println(sgd)
	fmt.Println(myr)
}
