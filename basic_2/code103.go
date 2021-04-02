package main

import "fmt"

func main() {
	j := 2
	switch j {
	case 0:
		fmt.Println("kosong")
	case 1:
		fmt.Println("satu")
	case 2:
		fmt.Println("dua")
	case 3:
		fmt.Println("tiga")
	default:
		fmt.Println("blur lah.")
	}
}
