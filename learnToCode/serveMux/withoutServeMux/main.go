package main

import (
	"io"
	"net/http"
)

type jumba string

func (j jumba) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/dog":
		io.WriteString(w, "Jumba Lina loves Dogs!")
	case "/cat/":
		io.WriteString(w, "Jumba Lina loves Cats too!!")
	}
}

func main() {
	var luiza jumba

	// mux := http.NewServeMux()
	// mux.Handle()

	http.ListenAndServe(":8080", luiza)
}
