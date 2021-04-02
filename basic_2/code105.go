package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func readCSV(input string) (data [][]string, err error) {
	sf, err := os.Open(input)
	if err != nil {
		fmt.Println(err, "[", input, "]")
		log.Fatal(err, input)
	}
	data, err = csv.NewReader(sf).ReadAll()
	return
}

func main() {
	data, _ := readCSV("data/test.csv")
	fmt.Println(data)
}
