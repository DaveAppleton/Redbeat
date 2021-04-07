package main

import (
	"fmt"
	"sort"
)

//start OMIT
type Bid int

type BidInfoSlice []Bid

func (a BidInfoSlice) Len() int      { return len(a) }
func (a BidInfoSlice) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a BidInfoSlice) Less(i, j int) bool {
	return a[i] < a[j]
}

func main() {
	data := BidInfoSlice{10, 1, 9, 2, 8, 3, 7, 4, 6, 5}

	sort.Sort(data)

	for pos, val := range data {
		fmt.Println(pos, val)
	}
}

// end OMIT
