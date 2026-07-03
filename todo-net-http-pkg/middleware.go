package main

import (
	"fmt"
	"net/http"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		//serve
		next.ServeHTTP(w, r)

		fmt.Printf(
			"Request processed : %s %s (%s)\n",
			r.Method, r.URL.Path,
			time.Since(start),
		)
	})
}
