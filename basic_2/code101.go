package main

import (
	"fmt"
)

func main() {

	for index := 0; index < 3; index++ {
		fmt.Println(index)
	}

	names := []string{"tom", "dick", "harry"}

	for pos, name := range names {
		fmt.Println(pos, name)
	}

}
