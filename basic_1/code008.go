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
	function, ok := allData[name]
	if !ok {
		function = "not found"
	}
	fmt.Println(name, "is", function)
}
