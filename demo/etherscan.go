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

func checkSender(data [][]string) (outData [][]string) {
	// check that "From" column contains "0x0000000000000000000000000000000000000000"
	checkString := "0x0000000000000000000000000000000000000000"

	for _, line := range data {
		if line[4] == checkString {
			outData = append(outData, line)
		}
	}
	return
}

func checkLastDigits(data [][]string) (d2, d3, d4 int) {
	for _, line := range data {
		fmt.Println("++", line[6])
		switch len(line[6]) {
		case 2:
			d2++
		case 3:
			d3++
		case 4:
			d4++
		}
	}
	return
}

func main() {
	data, err := readCSV("./etherscan.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	filteredData := checkSender(data)
	d2Counter, d3Counter, d4Counter := checkLastDigits(filteredData)
	fmt.Println(d2Counter, d3Counter, d4Counter)
}
