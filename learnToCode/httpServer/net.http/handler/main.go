package main

import (
	"fmt"
	"net/http"
)

type hotdog string

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "is very good")
}

func main() {
	var doggy hotdog
	http.ListenAndServe(":8080", doggy)
}
