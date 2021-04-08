package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var visits int

func data(w http.ResponseWriter, r *http.Request) {
	var dx struct {
		Visits int
	}
	dx.Visits = visits
	visits++
	json.NewEncoder(w).Encode(dx)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if len(name) == 0 {
		name = "you"
	}
	fmt.Fprintln(w, "hello", name)
}

func logo(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "files/1024px-AirAsia_New_Logo.svg.png")
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "files/index.html")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/data", data)
	http.HandleFunc("/sayhello", sayHello)
	http.HandleFunc("/logo", logo)
	http.ListenAndServe(":8003", nil)
}
