package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("got / request")
	//io.WriteString(w, "This is my website!\n")
	w.Write([]byte("This is my website"))
}

func getHello(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("got /hello request")
	io.WriteString(w, "Hello! HTTP\n")
}
func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(":3333", mux)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server was closed.")
	} else if err != nil {
		fmt.Println("Error starting server", err)
		os.Exit(1)
	}
}
