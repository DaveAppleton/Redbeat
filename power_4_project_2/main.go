package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
)

func main() {
	var fName string
	flag.StringVar(&fName, "file", "", "file to understand")
	flag.Parse()
	var data map[string]interface{}
	bData, err := ioutil.ReadFile(fName)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(bData, &data)
	if err != nil {
		log.Fatal(err)
	}
	for key, val := range data {
		vo := reflect.ValueOf(val)
		if vo.Kind() == "map" {
			for k2, v2 := range val.(map[string]interface{}) {
				fmt.Println(key, ".", k2, "=", v2)
			}
		} else {
			fmt.Println(key, val, vo.Kind())
		}

	}
}
