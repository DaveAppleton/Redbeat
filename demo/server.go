package main

import (
	"fmt"
	"net/http"
)

func normal(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "files/index.html")

}

func special(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "this is our test")

}

func icon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "files/DaveTheHatIcon.png")
}
func logo(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "files/1024px-AirAsia_New_Logo.svg.png")
}

func main() {
	http.HandleFunc("/", normal)
	http.HandleFunc("/special", special)
	http.HandleFunc("/favicon.ico", icon)
	http.HandleFunc("/logo", logo)
	err := http.ListenAndServe(":8003", nil)
	if err != nil {
		fmt.Println(err)
	}
}
