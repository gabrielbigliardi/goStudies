package middleware

import (
	"log"
	"net/http"
)

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("got request with path %s, request id is '%s'", r.URL.Path, r.Context().Value("requestID"))

		next.ServeHTTP(w, r)
	})
}
