package main

import (
	"fmt"
	"sort"
)

//start OMIT
type Bid struct {
	Amount int
	Age    int
}

type BidInfoSlice []Bid

func (a BidInfoSlice) Len() int      { return len(a) }
func (a BidInfoSlice) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a BidInfoSlice) Less(i, j int) bool {
	if a[i].Amount < a[j].Amount {
		return true
	}
	if a[i].Amount > a[j].Amount {
		return false
	}
	return a[i].Age < a[j].Age
}

//mid OMIT
func main() {
	var data BidInfoSlice
	data = append(data, Bid{Amount: 100, Age: 20})
	data = append(data, Bid{Amount: 80, Age: 20})
	data = append(data, Bid{Amount: 60, Age: 55})
	data = append(data, Bid{Amount: 80, Age: 75})
	data = append(data, Bid{Amount: 60, Age: 20})
	data = append(data, Bid{Amount: 40, Age: 20})

	sort.Sort(data)

	for pos, val := range data {
		fmt.Println(pos, val)
	}
}

// end OMIT
