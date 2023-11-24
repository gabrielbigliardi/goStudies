package main

import (
	"io"
	"net/http"
)

func d(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "dog dog doggy")
}

func c(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "cat cat catty")
}

func main() {

	http.HandleFunc("/", d)
	http.HandleFunc("/cat/", c)

	http.ListenAndServe(":8080", nil)
}
