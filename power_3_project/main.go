package main

import (
	"fmt"
	"log"

	"af.sorry.logrates/exdb"
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
	fmt.Println("SGD", sgd)
	fmt.Println("MYR", myr)
	fmt.Println(myr.Div(sgd))
	n := exdb.NewExchange("SGD", "MYR")
	n.UpdateExchange(myr.Div(sgd))

}
