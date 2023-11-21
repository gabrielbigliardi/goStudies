package main

import (
	"fmt"
	"io"
	"net/http"
)

type foo int

func (f foo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handle")
	io.WriteString(w, "Handle no /dog")
}

func g(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandleFunc")
	io.WriteString(w, "HandleFunc no /cat")
}

func main() {
	var f foo

	http.Handle("/dog", f)
	http.HandleFunc("/cat", g)

	http.ListenAndServe(":8080", nil)
}
