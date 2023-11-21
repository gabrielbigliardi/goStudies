package main

import (
	"net/http"

	"github.com/gabrielbigliardi/serverStudies/handlers"
	"github.com/gabrielbigliardi/serverStudies/middleware"
)

func main() {
	http.Handle("/first", middleware.Uuid(middleware.Log(handlers.First{})))

	http.ListenAndServe(":8080", nil)
}
