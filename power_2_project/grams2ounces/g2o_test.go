package grams2ounces

import "testing"

func TestG2O(t *testing.T) {
	wio := GramsToOunces(42)
	if wio != 1.5 {
		t.Log("Weight expected : 1.5, received", wio)
		t.Fail()
	}
}

func TestG2Of(t *testing.T) {
	wio := GramsToOunces(29)
	if wio != 1 {
		t.Log("Weight expected : 1, received", wio)
		t.Fail()
	}
}
