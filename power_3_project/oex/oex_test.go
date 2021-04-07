package oex

import (
	"fmt"
	"testing"
)

func TestOex(t *testing.T) {
	responses, err := GetLatest()
	if err != nil {
		if responses != nil {
			fmt.Println(responses.Error.Description)
		}
		t.Fatal(err)
	}
	sgd := responses.Rates["SGD"]
	myr := responses.Rates["MYR"]
	fmt.Println("SGD", sgd)
	fmt.Println("MYR", myr)
	thb := responses.Rates["THB"]
	fmt.Println("THB", thb)
	gbp := responses.Rates["GBP"]
	fmt.Println("GBP", gbp)
	t.Fail()
}
