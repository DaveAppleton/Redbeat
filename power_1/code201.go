package main

import (
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if len(name) == 0 {
		name = "you"
	}
	fmt.Fprintln(w, "hello", name)
}

func main() {
	http.HandleFunc("/", sayHello)
	http.ListenAndServe(":8003", nil)
}
