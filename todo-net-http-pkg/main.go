package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)
	handler := loggingMiddleware(mux)

	//handle errors
	err := http.ListenAndServe(":8000", handler)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server was terminated")
	} else if err != nil {
		fmt.Println("Error starting server", err)
		os.Exit(1)
	}
}
