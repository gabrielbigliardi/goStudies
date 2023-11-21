package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

func Uuid(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.New().String()

		ctx := context.WithValue(r.Context(), "requestID", requestID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
