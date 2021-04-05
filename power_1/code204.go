package main

import (
	"fmt"
	"net/http"
)

type mux struct{}

func (m *mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	if len(r.URL.Path) > 6 && r.URL.Path[:7] == "/greet/" {
		fmt.Fprintln(w, "Greetings", r.URL.Path[7:])
	} else if len(r.URL.Path) > 5 && r.URL.Path[:5] == "/inc/" {
		http.ServeFile(w, r, r.URL.Path[5:])
	} else {
		w.Write([]byte("I'm confused"))
	}
}

func main() {
	m := mux{}
	http.ListenAndServe(":8080", &m)
}
