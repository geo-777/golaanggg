package main

import "net/http"

func rootHandler(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Hello, welcome to todo app"))
}
