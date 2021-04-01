package main

import (
	"fmt"
)

func main() {
	allData := make(map[string]string)
	allData["Tom"] = "The fool"
	allData["Dick"] = "The clown"
	allData["Harry"] = "The brave"

	name := "Tom"
	fmt.Println(name, "is", allData[name])
}
