package main

import (
	"io"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "dog dog doggy")
}

type hotcat int

func (c hotcat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "cat cat catty")
}

func main() {
	var dog hotdog
	var cat hotcat

	//Handle registers the handler for the given pattern in the DefaultServeMux.
	//The documentation for ServeMux explains how patterns are matched.
	http.Handle("/", dog)
	http.Handle("/cat/", cat)

	//The handler is typically nil, in which case the DefaultServeMux is used.
	http.ListenAndServe(":8080", nil)
}
