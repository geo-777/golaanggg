package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

// trying out Handle interface
type HomeHandler struct{}

func (h HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	name := r.URL.Query().Get("name")
	fmt.Println("Got /")
	fmt.Fprintln(w, "/ page has been loaded for", name)
}

func getHello(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("got /hello request")
	io.WriteString(w, "Hello! HTTP\n")
}

func echoFunc(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, string(body))
}

func main() {

	mux := http.NewServeMux()
	mux.Handle("/", HomeHandler{})
	mux.HandleFunc("/hello", getHello)
	mux.HandleFunc("/echo", echoFunc)

	err := http.ListenAndServe(":3333", mux)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server was closed.")
	} else if err != nil {
		fmt.Println("Error starting server", err)
		os.Exit(1)
	}
}
