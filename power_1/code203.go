package main

import "net/http"

type mux struct {
	Name string
}

func (m *mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hello " + m.Name))
}

func main() {
	m := mux{Name: "Charlie"}
	http.ListenAndServe(":8080", &m)
}
