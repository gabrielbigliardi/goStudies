package handlers

import (
	"fmt"
	"net/http"
)

type First struct{}

func (rcv First) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "first Response")
}
